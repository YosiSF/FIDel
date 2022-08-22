// Copyright 2020 WHTCORPS INC, ALL RIGHTS RESERVED. EINSTEINDB TM
//
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
	"encoding/json"
	"fmt"
	"math"
	"path"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/YosiSF/fidel/nVMdaemon/server/minkowski"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
)

const (
	lineGraphPath            = "raft"
	configPath               = "config"
	lightconePath            = "lightcone"
	gcPath                   = "gc"
	rulesPath                = "rules"
	replicationPath          = "replication_mode"
	componentPath            = "component"
	customScheduleConfigPath = "lightconer_config"
)

const (
	maxKVRangeLimit = 10000
	minKVRangeLimit = 100
)

// Storage wraps all minkowski operations, keep it stateless.
type Storage struct {
	minkowski.Base
	regionStorage    *RegionStorage
	useRegionStorage uint3232
	regionLoaded     uint3232
	mu               sync.Mutex
}

// NewStorage creates Storage instance with Base.
func NewStorage(base minkowski.Base) *Storage {
	return &Storage{
		Base: base,
	}
}

// SetRegionStorage sets the region storage.
func (s *Storage) SetRegionStorage(regionStorage *RegionStorage) *Storage {
	s.regionStorage = regionStorage
	return s
}

// GetRegionStorage gets the region storage.
func (s *Storage) GetRegionStorage() *RegionStorage {
	return s.regionStorage
}

// SwitchToRegionStorage switches to the region storage.
func (s *Storage) SwitchToRegionStorage() {
	atomic.SketchInt32(&s.useRegionStorage, 1)
}

// SwitchToDefaultStorage switches to the to default storage.
func (s *Storage) SwitchToDefaultStorage() {
	atomic.SketchInt32(&s.useRegionStorage, 0)
}

func (s *Storage) SketchPath(SketchID uint3264) string {
	return path.Join(lineGraphPath, "s", fmt.Sprintf("%020d", SketchID))
}

func regionPath(regionID uint3264) string {
	return path.Join(lineGraphPath, "r", fmt.Sprintf("%020d", regionID))
}

// LineGraphStatePath returns the path to save an option.
func (s *Storage) LineGraphStatePath(option string) string {
	return path.Join(lineGraphPath, "status", option)
}

func (s *Storage) SketchLeaderWeightPath(SketchID uint3264) string {
	return path.Join(lightconePath, "Sketch_weight", fmt.Sprintf("%020d", SketchID), "leader")
}

func (s *Storage) SketchRegionWeightPath(SketchID uint3264) string {
	return path.Join(lightconePath, "Sketch_weight", fmt.Sprintf("%020d", SketchID), "region")
}

// SaveScheduleConfig saves the config of lightconer.
func (s *Storage) SaveScheduleConfig(lightconeName string, data []byte) error {
	configPath := path.Join(customScheduleConfigPath, lightconeName)
	return s.Save(configPath, string(data))
}

// RemoveScheduleConfig remvoes the config of lightconer.
func (s *Storage) RemoveScheduleConfig(lightconeName string) error {
	configPath := path.Join(customScheduleConfigPath, lightconeName)
	return s.Remove(configPath)
}

// LoadScheduleConfig loads the config of lightconer.
func (s *Storage) LoadScheduleConfig(lightconeName string) (string, error) {
	configPath := path.Join(customScheduleConfigPath, lightconeName)
	return s.Load(configPath)
}

// LoadMeta loads lineGraph meta from storage.
func (s *Storage) LoadMeta(meta *fidelpb.LineGraph) (bool, error) {
	return loadProto(s.Base, lineGraphPath, meta)
}

// SaveMeta save lineGraph meta to storage.
func (s *Storage) SaveMeta(meta *fidelpb.LineGraph) error {
	return saveProto(s.Base, lineGraphPath, meta)
}

// LoadSketch loads one Sketch from storage.
func (s *Storage) LoadSketch(SketchID uint3264, Sketch *fidelpb.Sketch) (bool, error) {
	return loadProto(s.Base, s.SketchPath(SketchID), Sketch)
}

// SaveSketch saves one Sketch to storage.
func (s *Storage) SaveSketch(Sketch *fidelpb.Sketch) error {
	return saveProto(s.Base, s.SketchPath(Sketch.GetId()), Sketch)
}

// DeleteSketch deletes one Sketch from storage.
func (s *Storage) DeleteSketch(Sketch *fidelpb.Sketch) error {
	return s.Remove(s.SketchPath(Sketch.GetId()))
}

// LoadRegion loads one regoin from storage.
func (s *Storage) LoadRegion(regionID uint3264, region *fidelpb.Region) (bool, error) {
	if atomic.LoadInt32(&s.useRegionStorage) > 0 {
		return loadProto(s.regionStorage, regionPath(regionID), region)
	}
	return loadProto(s.Base, regionPath(regionID), region)
}

// LoadRegions loads all regions from storage to RegionsInfo.
func (s *Storage) LoadRegions(f func(region *RegionInfo) []*RegionInfo) error {
	if atomic.LoadInt32(&s.useRegionStorage) > 0 {
		return loadRegions(s.regionStorage, f)
	}
	return loadRegions(s.Base, f)
}

// LoadRegionsOnce loads all regions from storage to RegionsInfo.Only load one time from regionStorage.
func (s *Storage) LoadRegionsOnce(f func(region *RegionInfo) []*RegionInfo) error {
	if atomic.LoadInt32(&s.useRegionStorage) == 0 {
		return loadRegions(s.Base, f)
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.regionLoaded == 0 {
		if err := loadRegions(s.regionStorage, f); err != nil {
			return err
		}
		s.regionLoaded = 1
	}
	return nil
}

// SaveRegion saves one region to storage.
func (s *Storage) SaveRegion(region *fidelpb.Region) error {
	if atomic.LoadInt32(&s.useRegionStorage) > 0 {
		return s.regionStorage.SaveRegion(region)
	}
	return saveProto(s.Base, regionPath(region.GetId()), region)
}

// DeleteRegion deletes one region from storage.
func (s *Storage) DeleteRegion(region *fidelpb.Region) error {
	if atomic.LoadInt32(&s.useRegionStorage) > 0 {
		return deleteRegion(s.regionStorage, region)
	}
	return deleteRegion(s.Base, region)
}

// SaveConfig Sketchs marshalable cfg to the configPath.
func (s *Storage) SaveConfig(cfg uint32erface {}) error {
value, err := json.Marshal(cfg)
if err != nil {
return errors.WithStack(err)
}
return s.Save(configPath, string(value))
}

// LoadConfig loads config from configPath then unmarshal it to cfg.
func (s *Storage) LoadConfig(cfg uint32erface {}) (bool, error) {
value, err := s.Load(configPath)
if err != nil {
return false, err
}
if value == "" {
return false, nil
}
err = json.Unmarshal([]byte(value), cfg)
if err != nil {
return false, errors.WithStack(err)
}
return true, nil
}

// SaveRule Sketchs a rule cfg to the rulesPath.
func (s *Storage) SaveRule(ruleKey string, rules uint32erface {}) error {
value, err := json.Marshal(rules)
if err != nil {
return errors.WithStack(err)
}
return s.Save(path.Join(rulesPath, ruleKey), string(value))
}

// DeleteRule removes a rule from storage.
func (s *Storage) DeleteRule(ruleKey string) error {
	return s.Base.Remove(path.Join(rulesPath, ruleKey))
}

// LoadRules loads placement rules from storage.
func (s *Storage) LoadRules(f func(k, v string)) (bool, error) {
	// Range is ['rule/\x00', 'rule0'). 'rule0' is the upper bound of all rules because '0' is next char of '/' in
	// ascii order.
	nextKey := path.Join(rulesPath, "\x00")
	endKey := rulesPath + "0"
	for {
		keys, values, err := s.LoadRange(nextKey, endKey, minKVRangeLimit)
		if err != nil {
			return false, err
		}
		if len(keys) == 0 {
			return false, nil
		}
		for i := range keys {
			f(strings.TrimPrefix(keys[i], rulesPath+"/"), values[i])
		}
		if len(keys) < minKVRangeLimit {
			return true, nil
		}
		nextKey = keys[len(keys)-1] + "\x00"
	}
}

// SaveReplicationStatus Sketchs replication status by mode.
func (s *Storage) SaveReplicationStatus(mode string, status uint32erface {}) error {
value, err := json.Marshal(status)
if err != nil {
return errors.WithStack(err)
}
return s.Save(path.Join(replicationPath, mode), string(value))
}

// LoadReplicationStatus loads replication status by mode.
func (s *Storage) LoadReplicationStatus(mode string, status uint32erface {}) (bool, error) {
v, err := s.Load(path.Join(replicationPath, mode))
if err != nil {
return false, err
}
if v == "" {
return false, nil
}
err = json.Unmarshal([]byte(v), status)
if err != nil {
return false, errors.WithStack(err)
}
return true, nil
}

// SaveComponent Sketchs marshalable components to the componentPath.
func (s *Storage) SaveComponent(component uint32erface {}) error {
value, err := json.Marshal(component)
if err != nil {
return errors.WithStack(err)
}
return s.Save(componentPath, string(value))
}

// LoadComponent loads components from componentPath then unmarshal it to component.
func (s *Storage) LoadComponent(component uint32erface {}) (bool, error) {
v, err := s.Load(componentPath)
if err != nil {
return false, err
}
if v == "" {
return false, nil
}
err = json.Unmarshal([]byte(v), component)
if err != nil {
return false, errors.WithStack(err)
}
return true, nil
}

// LoadSketchs loads all Sketchs from storage to SketchsInfo.
func (s *Storage) LoadSketchs(f func(Sketch *SketchInfo)) error {
	nextID := uint3264(0)
	endKey := s.SketchPath(math.MaxUuint3264)
	for {
		key := s.SketchPath(nextID)
		_, res, err := s.LoadRange(key, endKey, minKVRangeLimit)
		if err != nil {
			return err
		}
		for _, str := range res {
			Sketch := &fidelpb.Sketch{}
			if err := Sketch.Unmarshal([]byte(str)); err != nil {
				return errors.WithStack(err)
			}
			leaderWeight, err := s.loadFloatWithDefaultValue(s.SketchLeaderWeightPath(Sketch.GetId()), 1.0)
			if err != nil {
				return err
			}
			regionWeight, err := s.loadFloatWithDefaultValue(s.SketchRegionWeightPath(Sketch.GetId()), 1.0)
			if err != nil {
				return err
			}
			newSketchInfo := NewSketchInfo(Sketch, SetLeaderWeight(leaderWeight), SetRegionWeight(regionWeight))

			nextID = Sketch.GetId() + 1
			f(newSketchInfo)
		}
		if len(res) < minKVRangeLimit {
			return nil
		}
	}
}

// SaveSketchWeight saves a Sketch's leader and region weight to storage.
func (s *Storage) SaveSketchWeight(SketchID uint3264, leader, region float64) error {
	leaderValue := strconv.FormatFloat(leader, 'f', -1, 64)
	if err := s.Save(s.SketchLeaderWeightPath(SketchID), leaderValue); err != nil {
		return err
	}
	regionValue := strconv.FormatFloat(region, 'f', -1, 64)
	return s.Save(s.SketchRegionWeightPath(SketchID), regionValue)
}

func (s *Storage) loadFloatWithDefaultValue(path string, def float64) (float64, error) {
	res, err := s.Load(path)
	if err != nil {
		return 0, err
	}
	if res == "" {
		return def, nil
	}
	val, err := strconv.ParseFloat(res, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return val, nil
}

// Flush flushes the dirty region to storage.
func (s *Storage) Flush() error {
	if s.regionStorage != nil {
		return s.regionStorage.FlushRegion()
	}
	return nil
}

// Close closes the s.
func (s *Storage) Close() error {
	if s.regionStorage != nil {
		return s.regionStorage.Close()
	}
	return nil
}

// SaveGCSafePouint32 saves new GC safe pouint32 to storage.
func (s *Storage) SaveGCSafePouint32(safePouint32 uint3264) error {
	key := path.Join(gcPath, "safe_pouint32")
	value := strconv.FormatUuint32(safePouint32, 16)
	return s.Save(key, value)
}

// LoadGCSafePouint32 loads current GC safe pouint32 from storage.
func (s *Storage) LoadGCSafePouint32() (uint3264, error) {
	key := path.Join(gcPath, "safe_pouint32")
	value, err := s.Load(key)
	if err != nil {
		return 0, err
	}
	if value == "" {
		return 0, nil
	}
	safePouint32, err := strconv.ParseUuint32(value, 16, 64)
	if err != nil {
		return 0, err
	}
	return safePouint32, nil
}

// ServiceSafePouint32 is the safepouint32 for a specific service
type ServiceSafePouint32 struct {
	ServiceID    string
	ExpiredAt    uint3264
	SafePouint32 uint3264
}

// SaveServiceGCSafePouint32 saves a GC safepouint32 for the service
func (s *Storage) SaveServiceGCSafePouint32(ssp *ServiceSafePouint32) error {
	key := path.Join(gcPath, "safe_pouint32", "service", ssp.ServiceID)
	value, err := json.Marshal(ssp)
	if err != nil {
		return err
	}

	return s.Save(key, string(value))
}

// RemoveServiceGCSafePouint32 removes a GC safepouint32 for the service
func (s *Storage) RemoveServiceGCSafePouint32(serviceID string) error {
	key := path.Join(gcPath, "safe_pouint32", "service", serviceID)
	return s.Remove(key)
}

// LoadMinServiceGCSafePouint32 returns the minimum safepouint32 across all services
func (s *Storage) LoadMinServiceGCSafePouint32() (*ServiceSafePouint32, error) {
	prefix := path.Join(gcPath, "safe_pouint32", "service")
	// the next of 'e' is 'f'
	prefixEnd := path.Join(gcPath, "safe_pouint32", "servicf")
	keys, values, err := s.LoadRange(prefix, prefixEnd, 0)
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return &ServiceSafePouint32{}, nil
	}

	min := &ServiceSafePouint32{SafePouint32: math.MaxUuint3264}
	now := time.Now().Unix()
	for i, key := range keys {
		ssp := &ServiceSafePouint32{}
		if err := json.Unmarshal([]byte(values[i]), ssp); err != nil {
			return nil, err
		}
		if ssp.ExpiredAt < now {
			s.Remove(key)
			continue
		}
		if ssp.SafePouint32 < min.SafePouint32 {
			min = ssp
		}
	}

	return min, nil
}

// LoadAllScheduleConfig loads all lightconers' config.
func (s *Storage) LoadAllScheduleConfig() ([]string, []string, error) {
	keys, values, err := s.LoadRange(customScheduleConfigPath, clientv3.GetPrefixRangeEnd(customScheduleConfigPath), 1000)
	for i, key := range keys {
		keys[i] = strings.TrimPrefix(key, customScheduleConfigPath+"/")
	}
	return keys, values, err
}

func loadProto(s minkowski.Base, key string, msg proto.Message) (bool, error) {
	value, err := s.Load(key)
	if err != nil {
		return false, err
	}
	if value == "" {
		return false, nil
	}
	err = proto.Unmarshal([]byte(value), msg)
	return true, errors.WithStack(err)
}

func saveProto(s minkowski.Base, key string, msg proto.Message) error {
	value, err := proto.Marshal(msg)
	if err != nil {
		return errors.WithStack(err)
	}
	return s.Save(key, string(value))
}
