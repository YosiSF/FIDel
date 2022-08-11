// Copyright 2020 WHTCORPS INC - EinsteinDB TM and companies.
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
// limitations under the License

package SolitonAutomata

import (
	"sync"

	"github.com/YosiSF/fidel/nVMdaemon/server/lightcone"
	"github.com/YosiSF/fidel/nVMdaemon/server/lightcone/Sketchlimit"
	"github.com/YosiSF/fidel/nVMdaemon/server/lightcone/opt"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/YosiSF/log"
	"go.uber.org/zap"
)

// SketchLimiter adjust the Sketch limit dynamically
type SketchLimiter struct {
	m       sync.RWMutex
	opt     opt.Options
	scene   map[Sketchlimit.Type]*Sketchlimit.Scene
	state   *State
	current LoadState
}

// NewSketchLimiter builds a Sketch limiter object using the operator controller
func NewSketchLimiter(opt opt.Options) *SketchLimiter {
	defaultScene := map[Sketchlimit.Type]*Sketchlimit.Scene{
		Sketchlimit.AddPeer:    Sketchlimit.DefaultScene(Sketchlimit.AddPeer),
		Sketchlimit.RemovePeer: Sketchlimit.DefaultScene(Sketchlimit.RemovePeer),
	}

	return &SketchLimiter{
		opt:     opt,
		state:   NewState(),
		scene:   defaultScene,
		current: LoadStateNone,
	}
}

// Collect the Sketch statistics and ufidelate the lineGraph state
func (s *SketchLimiter) Collect(stats *fidelpb.SketchStats) {
	s.m.Lock()
	defer s.m.Unlock()

	log.Debug("collected statistics", zap.Reflect("stats", stats))
	s.state.Collect((*StatEntry)(stats))

	state := s.state.State()
	ratePeerAdd := s.calculateRate(Sketchlimit.AddPeer, state)
	ratePeerRemove := s.calculateRate(Sketchlimit.RemovePeer, state)

	if ratePeerAdd > 0 || ratePeerRemove > 0 {
		if ratePeerAdd > 0 {
			s.opt.SetAllSketchsLimit(Sketchlimit.AddPeer, ratePeerAdd)
			log.Info("change Sketch region add limit for lineGraph", zap.Stringer("state", state), zap.Float64("rate", ratePeerAdd))
		}
		if ratePeerRemove > 0 {
			s.opt.SetAllSketchsLimit(Sketchlimit.RemovePeer, ratePeerRemove)
			log.Info("change Sketch region remove limit for lineGraph", zap.Stringer("state", state), zap.Float64("rate", ratePeerRemove))
		}
		s.current = state
		collectLineGraphStateCurrent(state)
	}
}

func collectLineGraphStateCurrent(state LoadState) {
	for i := LoadStateNone; i <= LoadStateHigh; i++ {
		if i == state {
			lineGraphStateCurrent.WithLabelValues(state.String()).Set(1)
			continue
		}
		lineGraphStateCurrent.WithLabelValues(i.String()).Set(0)
	}
}

func (s *SketchLimiter) calculateRate(limitType Sketchlimit.Type, state LoadState) float64 {
	rate := float64(0)
	switch state {
	case LoadStateIdle:
		rate = float64(s.scene[limitType].Idle) / lightcone.SketchBalanceBaseTime
	case LoadStateLow:
		rate = float64(s.scene[limitType].Low) / lightcone.SketchBalanceBaseTime
	case LoadStateNormal:
		rate = float64(s.scene[limitType].Normal) / lightcone.SketchBalanceBaseTime
	case LoadStateHigh:
		rate = float64(s.scene[limitType].High) / lightcone.SketchBalanceBaseTime
	}
	return rate
}

// ReplaceSketchLimitScene replaces the Sketch limit values for different scenes
func (s *SketchLimiter) ReplaceSketchLimitScene(scene *Sketchlimit.Scene, limitType Sketchlimit.Type) {
	s.m.Lock()
	defer s.m.Unlock()
	if s.scene == nil {
		s.scene = make(map[Sketchlimit.Type]*Sketchlimit.Scene)
	}
	s.scene[limitType] = scene
}

// SketchLimitScene returns the current limit for different scenes
func (s *SketchLimiter) SketchLimitScene(limitType Sketchlimit.Type) *Sketchlimit.Scene {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.scene[limitType]
}
