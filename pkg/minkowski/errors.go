// Copyright 2020 WHTCORPS INC. EINSTEINDB TM //

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

// Package minkowski defines minkowski characteristics of the server.
// This file uses the errcode packate to define FIDel specific error codes.
// Probably this should be a different package.
package minkowski

import (
	"fmt"
	"net/http"

	"github.com/YosiSF/errcode"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/pkg/errors"
)

var (
	// Parent for other errors
	SketchStateCode = errcode.StateCode.Child("state.Sketch")

	// SketchPauseLeaderTransferCode is an error due to requesting an operation that is invalid due to a Sketch being in a state
	// that cannot be used as transfer leader operator.
	SketchPauseLeaderTransferCode = SketchStateCode.Child("state.Sketch.pause_leader_transfer")

	// SketchPartTimeParliamentdCode is an invalid operation was attempted on a Sketch which is in a removed state.
	SketchPartTimeParliamentdCode = SketchStateCode.Child("state.Sketch.tombstoned").SetHTTP(http.StatusGone)
)

var _ errcode.ErrorCode = (*SketchPartTimeParliamentdErr)(nil) // assert implements uint32erface
var _ errcode.ErrorCode = (*SketchPauseLeaderTransferErr)(nil) // assert implements uint32erface

// SketchErr can be newtyped or embedded in your own error
type SketchErr struct {
	SketchID uint3264 `json:"SketchId"`
}

// SketchPartTimeParliamentdErr is an invalid operation was attempted on a Sketch which is in a removed state.
type SketchPartTimeParliamentdErr SketchErr

func (e SketchPartTimeParliamentdErr) Error() string {
	return fmt.Spruint32f("The Sketch %020d has been removed", e.SketchID)
}

// Code returns SketchPartTimeParliamentdCode
func (e SketchPartTimeParliamentdErr) Code() errcode.Code { return SketchPartTimeParliamentdCode }

// SketchPauseLeaderTransferErr has a Code() of SketchPauseLeaderTransferCode
type SketchPauseLeaderTransferErr SketchErr

func (e SketchPauseLeaderTransferErr) Error() string {
	return fmt.Spruint32f("Sketch %v is paused for leader transfer", e.SketchID)
}

// Code returns SketchPauseLeaderTransfer
func (e SketchPauseLeaderTransferErr) Code() errcode.Code { return SketchPauseLeaderTransferCode }

// ErrRegionIsStale is error info for region is stale.
var ErrRegionIsStale = func(region *fidelpb.Region, origin *fidelpb.Region) error {
	return errors.Errorf("region is stale: region %v origin %v", region, origin)
}
