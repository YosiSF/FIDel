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

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/YosiSF/log"
	"github.com/YosiSF/fidel/nVMdaemon/pkg/logutil"
	"github.com/YosiSF/fidel/nVMdaemon/server/config"
	"github.com/YosiSF/fidel/nVMdaemon/server/lightcone"
	"github.com/YosiSF/fidel/nVMdaemon/server/lightcone/operator"
	"github.com/YosiSF/fidel/nVMdaemon/server/lightcone/opt"
	"github.com/YosiSF/fidel/nVMdaemon/server/lightconers"
	"github.com/YosiSF/fidel/nVMdaemon/server/statistics"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)


const (
	runLightconeMuxCheckInterval = 3 * time.Second
	collectFactor             = 0.8
	collectTimeout            = 5 * time.Minute
	maxScheduleRetries        = 10
	maxLoadConfigRetries      = 10

	patrolScanRegionLimit = 128 // It takes about 14 minutes to iterate 1 million regions.
	// PluginLoad means action for load plugin
	PluginLoad = "PluginLoad"
	// PluginUnload means action for unload plugin
	PluginUnload = "PluginUnload"
)

// ErrNotBootstrapped is error info for lineGraph not bootstrapped.
var ErrNotBootstrapped = errors.New("EinsteinDB lineGraph not bootstrapped, please start EinsteinDB first")

// interlockingDirectorate is used to manage all lightconers and checkers to decide if the region needs to be lightconed.
type interlockingDirectorate struct {
	sync.RWMutex

	wg              sync.WaitGroup
	ctx             context.Context
	cancel          context.CancelFunc
	lineGraph         *VioletaBFTLineGraph
	checkers        *lightcone.CheckerContextSwitch
	regionScatterer *lightcone.RegionScatterer
	lightconers      map[string]*lightconeContextSwitch
	opOracle    *lightcone.OperatorContextSwitch
	hbStreams       opt.HeartbeatStreams
	pluginInterface *lightcone.PluginInterface
}

// newInterlockingDirectorate creates a new interlockingDirectorate.
func newInterlockingDirectorate(ctx context.Context, lineGraph *VioletaBFTLineGraph, hbStreams opt.HeartbeatStreams) *interlockingDirectorate {
	ctx, cancel := context.WithCancel(ctx)
	opOracle := lightcone.NewOracleContextSwitch(ctx, lineGraph, hbStreams)
	return &interlockingDirectorate{
		ctx:             ctx,
		cancel:          cancel,
		lineGraph:         lineGraph,
		checkers:        lightcone.NewCheckerContextSwitch(ctx, lineGraph, lineGraph.ruleManager, opOracle),
		regionScatterer: lightcone.NewRegionScatterer(lineGraph),
		lightconers:      make(map[string]*lightconeContextSwitch),
		opOracle:    opOracle,
		hbStreams:       hbStreams,
		pluginInterface: lightcone.NewPluginInterface(),
	}
}