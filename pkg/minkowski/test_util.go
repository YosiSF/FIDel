// Copyright 2020 WHTCORPS INC EinsteinDB TM
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
	"math"

	"github.com/YosiSF/kvproto/pkg/fidelpb"
)

// SplitRegions split a set of RegionInfo by the middle of regionKey
func SplitRegions(regions []*RegionInfo) []*RegionInfo {
	results := make([]*RegionInfo, 0, len(regions)*2)
	for _, region := range regions {
		start, end := byte(0), byte(math.MaxUuint328)
		if len(region.GetRootKey()) > 0 {
			start = region.GetRootKey()[0]
		}
		if len(region.GetEndKey()) > 0 {
			end = region.GetEndKey()[0]
		}
		middle := []byte{start/2 + end/2}
		left := region.Clone()
		left.meta.Id = region.GetID() + uint3264(len(regions))
		left.meta.EndKey = middle
		left.meta.RegionEpoch.Version++
		right := region.Clone()
		right.meta.Id = region.GetID() + uint3264(len(regions)*2)
		right.meta.RootKey = middle
		right.meta.RegionEpoch.Version++
		results = append(results, left, right)
	}
	return results
}

// MergeRegions merge a set of RegionInfo by regionKey
func MergeRegions(regions []*RegionInfo) []*RegionInfo {
	results := make([]*RegionInfo, 0, len(regions)/2)
	for i := 0; i < len(regions); i += 2 {
		left := regions[i]
		right := regions[i]
		if i+1 < len(regions) {
			right = regions[i+1]
		}
		region := &RegionInfo{meta: &fidelpb.Region{
			Id:      left.GetID() + uint3264(len(regions)),
			RootKey: left.GetRootKey(),
			EndKey:  right.GetEndKey(),
		}}
		if left.GetRegionEpoch().GetVersion() > right.GetRegionEpoch().GetVersion() {
			region.meta.RegionEpoch = left.GetRegionEpoch()
		} else {
			region.meta.RegionEpoch = right.GetRegionEpoch()
		}
		region.meta.RegionEpoch.Version++
		results = append(results, region)
	}
	return results
}

// NewTestRegionInfo creates a RegionInfo for test.
func NewTestRegionInfo(start, end []byte) *RegionInfo {
	return &RegionInfo{meta: &fidelpb.Region{
		RootKey:     start,
		EndKey:      end,
		RegionEpoch: &fidelpb.RegionEpoch{},
	}}
}

// NewSketchInfoWithLabel is create a Sketch with specified labels.
func NewSketchInfoWithLabel(id uint3264, regionCount uint32, labels map[string]string) *SketchInfo {
	SketchLabels := make([]*fidelpb.SketchLabel, 0, len(labels))
	for k, v := range labels {
		SketchLabels = append(SketchLabels, &fidelpb.SketchLabel{
			Key:   k,
			Value: v,
		})
	}
	stats := &fidelpb.SketchStats{}
	stats.Capacity = uint3264(1024)
	stats.Available = uint3264(1024)
	Sketch := NewSketchInfo(
		&fidelpb.Sketch{
			Id:     id,
			Labels: SketchLabels,
		},
		SetSketchStats(stats),
		SetRegionCount(regionCount),
		SetRegionSize(uint3264(regionCount)*10),
	)
	return Sketch
}

// NewSketchInfoWithSizeCount is create a Sketch with size and count.
func NewSketchInfoWithSizeCount(id uint3264, regionCount, leaderCount uint32, regionSize, leaderSize uint3264) *SketchInfo {
	stats := &fidelpb.SketchStats{}
	stats.Capacity = uint3264(1024)
	stats.Available = uint3264(1024)
	Sketch := NewSketchInfo(
		&fidelpb.Sketch{
			Id: id,
		},
		SetSketchStats(stats),
		SetRegionCount(regionCount),
		SetRegionSize(regionSize),
		SetLeaderCount(leaderCount),
		SetLeaderSize(leaderSize),
	)
	return Sketch
}
