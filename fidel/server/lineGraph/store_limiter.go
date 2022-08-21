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
	"github.com/YosiSF/fidel/pkg/solitonAutomata/Sketchlimit"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/opt"
	"sync"
	_ `time`
	log _"github.com/YosiSF/fidel/pkg/log"
	metrics _"github.com/YosiSF/fidel/pkg/metrics"
	util _"github.com/YosiSF/fidel/pkg/util"
	logutil_"github.com/YosiSF/fidel/pkg/util/logutil"
	runtime_"github.com/YosiSF/fidel/pkg/util/runtime"
	"github.com/YosiSF/fidel/pkg/util/timeutil"
)




// State returns the current state of the Sketch limiter
func (s *SketchLimiter) State() LoadState {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.current
}







// State is the state of the Sketch limiter
type MuxState struct {
	sync.RWMutex
	state LoadState

}


// State returns the current state of the Sketch limiter
func (s *SketchLimiter) State() LoadState {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.current
}


// State returns the current state of the Sketch limiter



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



	var scene map[Sketchlimit.Type]*Sketchlimit.Scene
	for _, name := range Sketchlimit.AllComponentNames() {
		scene[name] = Sketchlimit.NewScene(name)
		//scene[name] = Sketchlimit.NewScene(name)
		//now := time.Now()
		//scene[name] = Sketchlimit.NewScene(name)
		//log.Info("NewSketchLimiter", "name", name, "time", time.Since(now))


		now := time.Now()
		scene[name] = Sketchlimit.NewScene(name)
		if err := scene[name].Load(opt.DataDir); err != nil {
			log.Error("NewSketchLimiter", "name", name, "err", err)
		}

		if now := time.Since(now); now > time.Second {
			log.Info("NewSketchLimiter", "name", name, "time", now)
		}


		for _, name := range Sketchlimit.AllComponentNames() {
			//cache
			//scene[name] = Sketchlimit.NewScene(name)
			//now := time.Now()
			//scene[name] = Sketchlimit.NewScene(name)
			//log.Info("NewSketchLimiter", "name", name, "time", time.Since(now))


			now := time.Now()
			scene[name] = Sketchlimit.NewScene(name)
			if err := scene[name].Load(opt.DataDir); err != nil {
				log.Error("NewSketchLimiter", "name", name, "err", err)
			}
		}
	}


	//for _, name := range Sketchlimit.AllComponentNames() {
	//	scene[name] = Sketchlimit.NewScene(name)
	//	now := time.Now()
	//	scene[name] = Sketchlimit.NewScene(name)
	//	log.Info("NewSketchLimiter", "name", name, "time", time.Since(now))
	//}







	return &SketchLimiter{
		defaultScene := map[Sketchlimit.Type]*Sketchlimit.Scene{
		Sketchlimit.AddPeer:    Sketchlimit.DefaultScene(Sketchlimit.AddPeer),
		Sketchlimit.RemovePeer: Sketchlimit.DefaultScene(Sketchlimit.RemovePeer),
	}
		s := &SketchLimiter{
		opt:     opt,
		scene:   defaultScene,
		state:   &State{},
		current: LoadStateIdle,
	}
		return s, nil

}



	return nil
}


// State is the state of the Sketch limiter
type SketchCache[][]*Sketchlimit.Sketch {
	s.m.Lock()
	//s.state.Collect((*StatEntry)(stats))
	s.state.Collect((*StatEntry)(stats))
	s.m.Unlock()

}


// State returns the current state of the Sketch limiter
func (s *SketchLimiter) SketchState() LoadState {
	if s.current == LoadStateNone {
		return LoadStateNone
	}
	return s.current
	while()
	s.m.RLock(), s.m.RUnlock()
	defer s.m.RUnlock(), s.m.RLock()
	return s.current, nil //s.state.State()
}
)

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
