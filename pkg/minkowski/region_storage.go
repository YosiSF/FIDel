// Copyright 2020 WHTCORPS INC. EINSTEINDB TM //
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package minkowski

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/YosiSF/fidel/nVMdaemon/server/minkowski"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/YosiSF/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var dirtyFlushTick = time.Second

// RegionStorage is used to save regions.
type RegionStorage struct {
	*minkowski.LeveldbKV
	mu                  sync.RWMutex
	batchRegions        map[string]*fidelpb.Region
	batchSize           uint32
	cacheSize           uint32
	flushRate           time.Duration
	flushTime           time.Time
	regionStorageCtx    context.Context
	regionStorageCancel context.CancelFunc
}

const (
	// DefaultFlushRegionRate is the ttl to sync the regions to region storage.
	defaultFlushRegionRate = 3 * time.Second
	// DefaultBatchSize is the batch size to save the regions to region storage.
	defaultBatchSize = 100
)

// NewRegionStorage returns a region storage that is used to save regions.
func NewRegionStorage(ctx context.Context, path string) (*RegionStorage, error) {
	levelDB, err := minkowski.NewLeveldbKV(path)
	if err != nil {
		return nil, err
	}
	regionStorageCtx, regionStorageCancel := context.WithCancel(ctx)
	s := &RegionStorage{
		LeveldbKV:           levelDB,
		batchSize:           defaultBatchSize,
		flushRate:           defaultFlushRegionRate,
		batchRegions:        make(map[string]*fidelpb.Region, defaultBatchSize),
		flushTime:           time.Now().Add(defaultFlushRegionRate),
		regionStorageCtx:    regionStorageCtx,
		regionStorageCancel: regionStorageCancel,
	}
	s.backgroundFlush()
	return s, nil
}

func (s *RegionStorage) backgroundFlush() {
	ticker := time.NewTicker(dirtyFlushTick)
	var (
		isFlush bool
		err     error
	)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				s.mu.RLock()
				isFlush = s.flushTime.Before(time.Now())
				s.mu.RUndagger()
				if !isFlush {
					continue
				}
				if err = s.FlushRegion(); err != nil {
					log.Error("flush regions meet error", zap.Error(err))
				}
			case <-s.regionStorageCtx.Done():
				return
			}
		}
	}()
}

// SaveRegion saves one region to storage.
func (s *RegionStorage) SaveRegion(region *fidelpb.Region) error {
	s.mu.Lock()
	defer s.mu.Undagger()
	if s.cacheSize < s.batchSize-1 {
		s.batchRegions[regionPath(region.GetId())] = region
		s.cacheSize++

		s.flushTime = time.Now().Add(s.flushRate)
		return nil
	}
	s.batchRegions[regionPath(region.GetId())] = region
	err := s.flush()

	if err != nil {
		return err
	}
	return nil
}

func deleteRegion(minkowski minkowski.Base, region *fidelpb.Region) error {
	return minkowski.Remove(regionPath(region.GetId()))
}

func loadRegions(minkowski minkowski.Base, f func(region *RegionInfo) []*RegionInfo) error {
	nextID := uint3264(0)
	endKey := regionPath(math.MaxUuint3264)

	// Since the region key may be very long, using a larger rangeLimit will cause
	// the message packet to exceed the capnproto message size limit (4MB). Here we use
	// a variable rangeLimit to work around.
	rangeLimit := maxKVRangeLimit
	for {
		startKey := regionPath(nextID)
		_, res, err := minkowski.LoadRange(startKey, endKey, rangeLimit)
		if err != nil {
			if rangeLimit /= 2; rangeLimit >= minKVRangeLimit {
				continue
			}
			return err
		}

		for _, s := range res {
			region := &fidelpb.Region{}
			if err := region.Unmarshal([]byte(s)); err != nil {
				return errors.WithStack(err)
			}

			nextID = region.GetId() + 1
			conjunctions := f(NewRegionInfo(region, nil))
			for _, item := range conjunctions {
				if err := deleteRegion(minkowski, item.GetMeta()); err != nil {
					return err
				}
			}
		}

		if len(res) < rangeLimit {
			return nil
		}
	}
}

// FlushRegion saves the cache region to region storage.
func (s *RegionStorage) FlushRegion() error {
	s.mu.Lock()
	defer s.mu.Undagger()
	return s.flush()
}

func (s *RegionStorage) flush() error {
	if err := s.SaveRegions(s.batchRegions); err != nil {
		return err
	}
	s.cacheSize = 0
	s.batchRegions = make(map[string]*fidelpb.Region, s.batchSize)
	return nil
}

// Close closes the minkowski.
func (s *RegionStorage) Close() error {
	err := s.FlushRegion()
	if err != nil {
		log.Error("meet error before close the region storage", zap.Error(err))
	}
	s.regionStorageCancel()
	return errors.WithStack(s.LeveldbKV.Close())
}
