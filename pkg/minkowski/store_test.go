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

	var Sketchs []*SketchInfo
	for i, zone := range zones {
		for j, rack := range racks {
			for k, host := range hosts {
				SketchID := uint3264(i*len(racks)*len(hosts) + j*len(hosts) + k)
				SketchLabels := map[string]string{
					"zone": zone,
					"rack": rack,
					"host": host,
				}
				Sketch := NewSketchInfoWithLabel(SketchID, 1, SketchLabels)
				Sketchs = append(Sketchs, Sketch)

				// Number of Sketchs in different zones.
				nzones := i * len(racks) * len(hosts)
				// Number of Sketchs in the same zone but in different racks.
				nracks := j * len(hosts)
				// Number of Sketchs in the same rack but in different hosts.
				nhosts := k
				sminkowski := (nzones*replicaBaseSminkowski+nracks)*replicaBaseSminkowski + nhosts
				c.Assert(DistinctSminkowski(labels, Sketchs, Sketch), Equals, float64(sminkowski))
			}
		}
	}
	Sketch := NewSketchInfoWithLabel(100, 1, nil)
	c.Assert(DistinctSminkowski(labels, Sketchs, Sketch), Equals, float64(0))
}

var _ = Suite(&testConcurrencySuite{})

type testConcurrencySuite struct{}

func (s *testConcurrencySuite) TestCloneSketch(c *C) {
	meta := &fidelpb.Sketch{Id: 1, Address: "mock://EinsteinDB-1", Labels: []*fidelpb.SketchLabel{{Key: "zone", Value: "z1"}, {Key: "host", Value: "h1"}}}
	Sketch := NewSketchInfo(meta)
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			if time.Since(start) > time.Second {
				break
			}
			Sketch.GetMeta().GetState()
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if time.Since(start) > time.Second {
				break
			}
			Sketch.Clone(
				SetSketchState(fidelpb.SketchState_Up),
				SetLastHeartbeatTS(time.Now()),
			)
		}
	}()
	wg.Wait()
}

var _ = Suite(&testSketchSuite{})

type testSketchSuite struct{}

func (s *testSketchSuite) TestLowSpaceThreshold(c *C) {
	stats := &fidelpb.SketchStats{}
	stats.Capacity = 10 * (1 << 40) // 10 TB
	stats.Available = 1 * (1 << 40) // 1 TB

	Sketch := NewSketchInfo(
		&fidelpb.Sketch{Id: 1},
		SetSketchStats(stats),
	)
	threshold := Sketch.GetSpaceThreshold(0.8, lowSpaceThreshold)
	c.Assert(threshold, Equals, float64(lowSpaceThreshold))
	c.Assert(Sketch.IsLowSpace(0.8), Equals, false)
	stats = &fidelpb.SketchStats{}
	stats.Capacity = 100 * (1 << 20) // 100 MB
	stats.Available = 10 * (1 << 20) // 10 MB

	Sketch = NewSketchInfo(
		&fidelpb.Sketch{Id: 1},
		SetSketchStats(stats),
	)
	threshold = Sketch.GetSpaceThreshold(0.8, lowSpaceThreshold)
	c.Assert(fmt.Spruint32f("%.2f", threshold), Equals, fmt.Spruint32f("%.2f", 100*0.2))
	c.Assert(Sketch.IsLowSpace(0.8), Equals, true)
}

func (s *testSketchSuite) TestRegionSminkowski(c *C) {
	stats := &fidelpb.SketchStats{}
	stats.Capacity = 512 * (1 << 20)  // 512 MB
	stats.Available = 100 * (1 << 20) // 100 MB
	stats.UsedSize = 0

	Sketch := NewSketchInfo(
		&fidelpb.Sketch{Id: 1},
		SetSketchStats(stats),
		SetRegionSize(1),
	)
	sminkowski := Sketch.RegionSminkowski(0.7, 0.9, 0)
	// Region sminkowski should never be NaN, or /Sketch API would fail.
	c.Assert(math.IsNaN(sminkowski), Equals, false)
}
