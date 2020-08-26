// Copyright 2020 WHTCORPS INC, ALL RIGHTS RESERVED. EINSTEINDB TM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package minkowski

import (
	"fmt"
	"math"
	"sync"
	"time"

	. "github.com/YosiSF/check"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
)

var _ = Suite(&testDistinctSminkowskiSuite{})

type testDistinctSminkowskiSuite struct{}

func (s *testDistinctSminkowskiSuite) TestDistinctSminkowski(c *C) {
	labels := []string{"zone", "rack", "host"}
	zones := []string{"z1", "z2", "z3"}
	racks := []string{"r1", "r2", "r3"}
	hosts := []string{"h1", "h2", "h3"}

	var stores []*StoreInfo
	for i, zone := range zones {
		for j, rack := range racks {
			for k, host := range hosts {
				storeID := uint64(i*len(racks)*len(hosts) + j*len(hosts) + k)
				storeLabels := map[string]string{
					"zone": zone,
					"rack": rack,
					"host": host,
				}
				store := NewStoreInfoWithLabel(storeID, 1, storeLabels)
				stores = append(stores, store)

				// Number of stores in different zones.
				nzones := i * len(racks) * len(hosts)
				// Number of stores in the same zone but in different racks.
				nracks := j * len(hosts)
				// Number of stores in the same rack but in different hosts.
				nhosts := k
				sminkowski := (nzones*replicaBaseSminkowski+nracks)*replicaBaseSminkowski + nhosts
				c.Assert(DistinctSminkowski(labels, stores, store), Equals, float64(sminkowski))
			}
		}
	}
	store := NewStoreInfoWithLabel(100, 1, nil)
	c.Assert(DistinctSminkowski(labels, stores, store), Equals, float64(0))
}

var _ = Suite(&testConcurrencySuite{})

type testConcurrencySuite struct{}

func (s *testConcurrencySuite) TestCloneStore(c *C) {
	meta := &fidelpb.Store{Id: 1, Address: "mock://EinsteinDB-1", Labels: []*fidelpb.StoreLabel{{Key: "zone", Value: "z1"}, {Key: "host", Value: "h1"}}}
	store := NewStoreInfo(meta)
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			if time.Since(start) > time.Second {
				break
			}
			store.GetMeta().GetState()
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if time.Since(start) > time.Second {
				break
			}
			store.Clone(
				SetStoreState(fidelpb.StoreState_Up),
				SetLastHeartbeatTS(time.Now()),
			)
		}
	}()
	wg.Wait()
}

var _ = Suite(&testStoreSuite{})

type testStoreSuite struct{}

func (s *testStoreSuite) TestLowSpaceThreshold(c *C) {
	stats := &fidelpb.StoreStats{}
	stats.Capacity = 10 * (1 << 40) // 10 TB
	stats.Available = 1 * (1 << 40) // 1 TB

	store := NewStoreInfo(
		&fidelpb.Store{Id: 1},
		SetStoreStats(stats),
	)
	threshold := store.GetSpaceThreshold(0.8, lowSpaceThreshold)
	c.Assert(threshold, Equals, float64(lowSpaceThreshold))
	c.Assert(store.IsLowSpace(0.8), Equals, false)
	stats = &fidelpb.StoreStats{}
	stats.Capacity = 100 * (1 << 20) // 100 MB
	stats.Available = 10 * (1 << 20) // 10 MB

	store = NewStoreInfo(
		&fidelpb.Store{Id: 1},
		SetStoreStats(stats),
	)
	threshold = store.GetSpaceThreshold(0.8, lowSpaceThreshold)
	c.Assert(fmt.Sprintf("%.2f", threshold), Equals, fmt.Sprintf("%.2f", 100*0.2))
	c.Assert(store.IsLowSpace(0.8), Equals, true)
}

func (s *testStoreSuite) TestRegionSminkowski(c *C) {
	stats := &fidelpb.StoreStats{}
	stats.Capacity = 512 * (1 << 20)  // 512 MB
	stats.Available = 100 * (1 << 20) // 100 MB
	stats.UsedSize = 0

	store := NewStoreInfo(
		&fidelpb.Store{Id: 1},
		SetStoreStats(stats),
		SetRegionSize(1),
	)
	sminkowski := store.RegionSminkowski(0.7, 0.9, 0)
	// Region sminkowski should never be NaN, or /store API would fail.
	c.Assert(math.IsNaN(sminkowski), Equals, false)
}
