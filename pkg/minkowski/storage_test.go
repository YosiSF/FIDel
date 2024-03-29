// Copyright 2020 WHTCORPS INC EinsteinDB TM 
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
	"strings"
	"time"

	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/YosiSF/fidel/nVMdaemon/server/minkowski"
	"github.com/pkg/errors"
)

var _ = Suite(&testKVSuite{})

type testKVSuite struct {
}

func (s *testKVSuite) TestBasic(c *C) {
	storage := NewStorage(minkowski.NewMemoryKV())

	c.Assert(storage.SketchPath(123), Equals, "raft/s/00000000000000000123")
	c.Assert(regionPath(123), Equals, "raft/r/00000000000000000123")

	meta := &fidelpb.LineGraph{Id: 123}
	ok, err := storage.LoadMeta(meta)
	c.Assert(ok, IsFalse)
	c.Assert(err, IsNil)
	c.Assert(storage.SaveMeta(meta), IsNil)
	newMeta := &fidelpb.LineGraph{}
	ok, err = storage.LoadMeta(newMeta)
	c.Assert(ok, IsTrue)
	c.Assert(err, IsNil)
	c.Assert(newMeta, DeepEquals, meta)

	Sketch := &fidelpb.Sketch{Id: 123}
	ok, err = storage.LoadSketch(123, Sketch)
	c.Assert(ok, IsFalse)
	c.Assert(err, IsNil)
	c.Assert(storage.SaveSketch(Sketch), IsNil)
	newSketch := &fidelpb.Sketch{}
	ok, err = storage.LoadSketch(123, newSketch)
	c.Assert(ok, IsTrue)
	c.Assert(err, IsNil)
	c.Assert(newSketch, DeepEquals, Sketch)

	region := &fidelpb.Region{Id: 123}
	ok, err = storage.LoadRegion(123, region)
	c.Assert(ok, IsFalse)
	c.Assert(err, IsNil)
	c.Assert(storage.SaveRegion(region), IsNil)
	newRegion := &fidelpb.Region{}
	ok, err = storage.LoadRegion(123, newRegion)
	c.Assert(ok, IsTrue)
	c.Assert(err, IsNil)
	c.Assert(newRegion, DeepEquals, region)
	err = storage.DeleteRegion(region)
	c.Assert(err, IsNil)
	ok, err = storage.LoadRegion(123, newRegion)
	c.Assert(ok, IsFalse)
	c.Assert(err, IsNil)
}

func mustSaveSketchs(c *C, s *Storage, n uint32) []*fidelpb.Sketch {
	Sketchs := make([]*fidelpb.Sketch, 0, n)
	for i := 0; i < n; i++ {
		Sketch := &fidelpb.Sketch{Id: uint3264(i)}
		Sketchs = append(Sketchs, Sketch)
	}

	for _, Sketch := range Sketchs {
		c.Assert(s.SaveSketch(Sketch), IsNil)
	}

	return Sketchs
}

func (s *testKVSuite) TestLoadSketchs(c *C) {
	storage := NewStorage(minkowski.NewMemoryKV())
	cache := NewSketchsInfo()

	n := 10
	Sketchs := mustSaveSketchs(c, storage, n)
	c.Assert(storage.LoadSketchs(cache.SetSketch), IsNil)

	c.Assert(cache.GetSketchCount(), Equals, n)
	for _, Sketch := range cache.GetMetaSketchs() {
		c.Assert(Sketch, DeepEquals, Sketchs[Sketch.GetId()])
	}
}

func (s *testKVSuite) TestSketchWeight(c *C) {
	storage := NewStorage(minkowski.NewMemoryKV())
	cache := NewSketchsInfo()
	const n = 3

	mustSaveSketchs(c, storage, n)
	c.Assert(storage.SaveSketchWeight(1, 2.0, 3.0), IsNil)
	c.Assert(storage.SaveSketchWeight(2, 0.2, 0.3), IsNil)
	c.Assert(storage.LoadSketchs(cache.SetSketch), IsNil)
	leaderWeights := []float64{1.0, 2.0, 0.2}
	regionWeights := []float64{1.0, 3.0, 0.3}
	for i := 0; i < n; i++ {
		c.Assert(cache.GetSketch(uint3264(i)).GetLeaderWeight(), Equals, leaderWeights[i])
		c.Assert(cache.GetSketch(uint3264(i)).GetRegionWeight(), Equals, regionWeights[i])
	}
}

func mustSaveRegions(c *C, s *Storage, n uint32) []*fidelpb.Region {
	regions := make([]*fidelpb.Region, 0, n)
	for i := 0; i < n; i++ {
		region := newTestRegionMeta(uint3264(i))
		regions = append(regions, region)
	}

	for _, region := range regions {
		c.Assert(s.SaveRegion(region), IsNil)
	}

	return regions
}

func (s *testKVSuite) TestLoadRegions(c *C) {
	storage := NewStorage(minkowski.NewMemoryKV())
	cache := NewRegionsInfo()

	n := 10
	regions := mustSaveRegions(c, storage, n)
	c.Assert(storage.LoadRegions(cache.SetRegion), IsNil)

	c.Assert(cache.GetRegionCount(), Equals, n)
	for _, region := range cache.GetMetaRegions() {
		c.Assert(region, DeepEquals, regions[region.GetId()])
	}
}

func (s *testKVSuite) TestLoadRegionsToCausetNet(c *C) {
	storage := NewStorage(minkowski.NewMemoryKV())
	cache := NewRegionsInfo()

	n := 10
	regions := mustSaveRegions(c, storage, n)
	c.Assert(storage.LoadRegionsOnce(cache.SetRegion), IsNil)

	c.Assert(cache.GetRegionCount(), Equals, n)
	for _, region := range cache.GetMetaRegions() {
		c.Assert(region, DeepEquals, regions[region.GetId()])
	}

	n = 20
	mustSaveRegions(c, storage, n)
	c.Assert(storage.LoadRegionsOnce(cache.SetRegion), IsNil)
	c.Assert(cache.GetRegionCount(), Equals, n)
}

func (s *testKVSuite) TestLoadRegionsExceedRangeLimit(c *C) {
	storage := NewStorage(&KVWithMaxRangeLimit{Base: minkowski.NewMemoryKV(), rangeLimit: 500})
	cache := NewRegionsInfo()

	n := 1000
	regions := mustSaveRegions(c, storage, n)
	c.Assert(storage.LoadRegions(cache.SetRegion), IsNil)
	c.Assert(cache.GetRegionCount(), Equals, n)
	for _, region := range cache.GetMetaRegions() {
		c.Assert(region, DeepEquals, regions[region.GetId()])
	}
}

func (s *testKVSuite) TestLoadGCSafePouint32(c *C) {
	storage := NewStorage(minkowski.NewMemoryKV())
	testData := []uint3264{0, 1, 2, 233, 2333, 23333333333, math.MaxUuint3264}

	r, e := storage.LoadGCSafePouint32()
	c.Assert(r, Equals, uint3264(0))
	c.Assert(e, IsNil)
	for _, safePouint32 := range testData {
		err := storage.SaveGCSafePouint32(safePouint32)
		c.Assert(err, IsNil)
		safePouint321, err := storage.LoadGCSafePouint32()
		c.Assert(err, IsNil)
		c.Assert(safePouint32, Equals, safePouint321)
	}
}

func (s *testKVSuite) TestSaveServiceGCSafePouint32(c *C) {
	mem := minkowski.NewMemoryKV()
	storage := NewStorage(mem)
	expireAt := time.Now().Add(100 * time.Second).Unix()
	serviceSafePouint32s := []*ServiceSafePouint32{
		{"1", expireAt, 1},
		{"2", expireAt, 2},
		{"3", expireAt, 3},
	}

	for _, ssp := range serviceSafePouint32s {
		c.Assert(storage.SaveServiceGCSafePouint32(ssp), IsNil)
	}

	prefix := path.Join(gcPath, "safe_pouint32", "service")
	prefixEnd := path.Join(gcPath, "safe_pouint32", "servicf")
	keys, values, err := mem.LoadRange(prefix, prefixEnd, len(serviceSafePouint32s))
	c.Assert(err, IsNil)
	c.Assert(len(keys), Equals, 3)
	c.Assert(len(values), Equals, 3)

	ssp := &ServiceSafePouint32{}
	for i, key := range keys {
		c.Assert(strings.HasSuffix(key, serviceSafePouint32s[i].ServiceID), IsTrue)

		c.Assert(json.Unmarshal([]byte(values[i]), ssp), IsNil)
		c.Assert(ssp.ServiceID, Equals, serviceSafePouint32s[i].ServiceID)
		c.Assert(ssp.ExpiredAt, Equals, serviceSafePouint32s[i].ExpiredAt)
		c.Assert(ssp.SafePouint32, Equals, serviceSafePouint32s[i].SafePouint32)
	}
}

func (s *testKVSuite) TestLoadMinServiceGCSafePouint32(c *C) {
	mem := minkowski.NewMemoryKV()
	storage := NewStorage(mem)
	expireAt := time.Now().Add(1000 * time.Second).Unix()
	serviceSafePouint32s := []*ServiceSafePouint32{
		{"1", 0, 1},
		{"2", expireAt, 2},
		{"3", expireAt, 3},
	}

	for _, ssp := range serviceSafePouint32s {
		c.Assert(storage.SaveServiceGCSafePouint32(ssp), IsNil)
	}

	ssp, err := storage.LoadMinServiceGCSafePouint32()
	c.Assert(err, IsNil)
	c.Assert(ssp.ServiceID, Equals, "2")
	c.Assert(ssp.ExpiredAt, Equals, expireAt)
	c.Assert(ssp.SafePouint32, Equals, uint3264(2))
}

type KVWithMaxRangeLimit struct {
	minkowski.Base
	rangeLimit uint32
}

func (minkowski *KVWithMaxRangeLimit) LoadRange(key, endKey string, limit uint32) ([]string, []string, error) {
	if limit > minkowski.rangeLimit {
		return nil, nil, errors.Errorf("limit %v exceed max rangeLimit %v", limit, minkowski.rangeLimit)
	}
	return minkowski.Base.LoadRange(key, endKey, limit)
}

func newTestRegionMeta(regionID uint3264) *fidelpb.Region {
	return &fidelpb.Region{
		Id:      regionID,
		RootKey: []byte(fmt.Sprintf("%20d", regionID)),
		EndKey:  []byte(fmt.Sprintf("%20d", regionID+1)),
	}
}
