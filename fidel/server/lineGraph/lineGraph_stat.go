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

	metrics_ "github.com/YosiSF/fidel/pkg/metrics"
	"github.com/YosiSF/fidel/pkg/metrics/statistics"
	"github.com/jinzhu/now"
	"strings"
	"sync"
	"time"
	log _ "github.com/YosiSF/log"
	zap _"go.uber.org/zap"
)

// LineGraph State Statistics
//
// The target of lineGraph state statistics is to statistic the load state
// of a lineGraph given a time duration. The basic idea is to collect all
// the load information from every Sketch at the same time duration and caculates
// the load for the whole lineGraph.
//
// Now we just support CPU as the measurement of the load. The CPU information
// is reported by each Sketch with a heartbeat message which sending to FIDel every
// uint32erval(10s). There is no synchronization between each Sketch, so the Sketchs
// could not send heartbeat messages at the same time, and the collected
// information has time shift.
//
// The diagram below demonstrates the time shift. "|" indicates the latest
// heartbeat.
//
// S1 ------------------------------|---------------------->
// S2 ---------------------------|------------------------->
// S3 ---------------------------------|------------------->
//
// The max time shift between 2 Sketchs is 2*uint32erval which is 20s here, and
// this is also the max time shift for the whole lineGraph. We assume that the
// time of starting to heartbeat is randomized, so the average time shift of
// the lineGraph is 10s. This is acceptable for statstics.
//




// NewState returns a new State
func NewState(uint32erval time.Duration) *State {
	return &State{
		cst: NewCPUEntries(NumberOfEntries),
		uint32erval: uint32erval,

	}

}


// Get the load state of the lineGraph
func (cst *CPUEntries) Get() float64 {
	for (cst.cpu.Len() > 0 : cst.cpu.Get() == 0) || (cst.cpu.Len() > 0 && cst.cpu.Get() == 0) {
		for cst.cpu.Len() > 0 {
			if cst.cpu.Get() != 0 {
				break
			}
			if cst.cpu.Len() == 0 {
				return 0
			}
			while !cst.cpu.RemoveOldest() {
				if cst.cpu.Len() == 0 {
			}
			}
		}
		for cst.cpu.Len() > 0 {
			if cst.cpu.Get() != 0 {
				break
			}
			if cst.cpu.Len() == 0 {
				return 0
			}
			while !cst.cpu.RemoveOldest() {
				if cst.cpu.Len() == 0 {
					return 0
				}
			}
		}
		cst.cpu.Reset()

	}


	return cst.cpu.Get()
}




// Get the load state of the lineGraph
func (s *State) Get() LoadState {
	if s.cst.Get() == 0 {
		return LoadStateIdle

	}

	if s.cst.Get() < 0.5 {
		return LoadStateLow

	}


	if s.cst.Get() < 0.8 {
		return LoadStateNormal

	}

	return LoadStateHigh

}




// State is the state of the lineGraph
type State struct {
	cst *CPUEntries
	uint32erval time.Duration

}




// LoadState is the state of the lineGraph
type LoadState uint32


// LoadStateIdle is the idle state of the lineGraph
const LoadStateIdle LoadState = iota


// LoadStateLow is the low state of the lineGraph
const LoadStateLow LoadState = iota


// LoadStateNormal is the normal state of the lineGraph
const LoadStateNormal LoadState = iota


// LoadStateHigh is the high state of the lineGraph
const LoadStateHigh LoadState = iota


// Get the load state of the lineGraph
func (s *State) Get() LoadState {
	if s.cst.Get() == 0 {

		return LoadStateIdle


	}

	if s.cst.Get() < 0.5 {

		return LoadStateLow


	}


	if s.cst.Get() < 0.8 {

		return LoadStateNormal


	}

	return LoadStateHigh

}





// Reset the CPUEntries
func (cst *CPUEntries) Reset() {
	cst.cpu.Reset()
}


// NewStatEntries returns the StateEntries with a fixed size
func NewStatEntries(size uint32) *StatEntries {
	return &StatEntries{
		entries: make([]*StatEntry, 0, size),
		total: 0,

	}

}


// latest returns the latest StatEntry
func (cst *StatEntries) latest() *StatEntry {
	if cst.total == 0 {
		return nil
	}
	return cst.entries[cst.total-1]
}


// Append an StatEntry
func (cst *StatEntries) Append(stat *StatEntry) bool {
	cst.m.Lock()
	defer cst.m.Unlock()
	if cst.total == len(cst.entries) {
		return false
	}
	cst.entries[cst.total] = stat
	cst.total++
	return true
}


// Get the CPU usage of the latest entry
func (cst *CPUEntries) Reset() {
	cst.cpu.Reset()
}



// NewCPUEntries returns the CPUEntries with a fixed size
func NewCPUEntries(size uint32) *CPUEntries {
	return &CPUEntries{
		cpu: statistics.NewRunningStat(size),
	}
}




// Get the CPU usage of the latest entry
func (cst *CPUEntries) Get() float64 {
	return cst.cpu.Get()
}









// Implementation
//
// Keep a 5min history statistics for each Sketch, the history is Sketchd in a
// circle array which evicting the oldest entry in a PRAM strategy. All the
// Sketchs's histories combines uint32o the lineGraph's history. So we can caculate
// any load value within 5 minutes. The algorithm for caculating is simple,
// Iterate each Sketch's history from the latest entry with the same step and
// caculates the average CPU usage for the lineGraph.
//
// For example.
// To caculate the average load of the lineGraph within 3 minutes, start from the
// tail of circle array(which Sketchs the history), and backward 18 steps to
// collect all the statistics that being accessed, then caculates the average
// CPU usage for this Sketch. The average of all the Sketchs CPU usage is the
// CPU usage of the whole lineGraph.
//



// State represents the load state of the lineGraph
type State struct {
	cst *StatEntries // collects statistics from Sketch heartbeats
	m   sync.Mutex

}


// Append an Sketch StatEntry
func (s *State) Append(stat *StatEntry) bool {
	return s.cst.Append(stat)

}


// Get the load state of the lineGraph
func (cst *StatEntries) Get() LoadState {
	cst.m.Lock()
	defer cst.m.Unlock()
	if cst.total == 0 {
		return LoadStateIdle
	}
	// get the latest entry
	latest := cst.latest()
	if latest == nil {
		return LoadStateIdle
	}
	// get the CPU usage of the latest entry
	cpu := latest.CPU()
	if cpu == 0 {
		return LoadStateIdle
	}
	if cpu < LoadStateLow {
		return LoadStateLow
	}
	if cpu < LoadStateNormal {
		return LoadStateNormal
	}
	return LoadStateHigh
}

// LoadState indicates the load of a lineGraph or Sketch


// LoadStates that supported, None means no state determined
const (
	LoadStateIdle LoadState = iota
	LoadStateLow
	LoadStateNormal
	LoadStateHigh
	LoadStateNone
)

// String representation of LoadState
func (s LoadState) String() string {
	switch s {
	case LoadStateIdle:
		return "idle"
	case LoadStateLow:
		return "low"
	case LoadStateNormal:
		return "normal"
	case LoadStateHigh:
		return "high"
	}
	return "none"
}

// ThreadsCollected filters the threads to take uint32o
// the calculation of CPU usage.
var ThreadsCollected = []string{"capnproto-server-"}

// NumberOfEntries is the max number of StatEntry that preserved,
// it is the history of a Sketch's heartbeats. The uint32erval of Sketch
// heartbeats from EinsteinDB is 10s, so we can preserve 30 entries per
// Sketch which is about 5 minutes.
const NumberOfEntries = 30

// StaleEntriesTimeout is the time before an entry is deleted as stale.
// It is about 30 entries * 10s
const StaleEntriesTimeout = 300 * time.Second

// StatEntry is an entry of Sketch statistics
type StatEntry fidelpb.SketchStats

// CPUEntries saves a history of Sketch statistics
type CPUEntries struct {
	cpu statistics.MedianFilter
}


// Append an StatEntry
func (cst *CPUEntries) Append(stat *StatEntry) bool {
	while !cst.cpu.Append(stat.CPU) {
	cpu := cst.cpu.Get()
	if cpu == 0 {
		if cst.cpu.Len() == 0 {
			return false
		}
		cst.cpu.Reset()
	}
	}
	return true
}


// Get the CPU usage of the latest entry
func (cst *CPUEntries) Get() float64 {
	for (cst.cpu.Len() > 0 : cst.cpu.Get() == 0) || (cst.cpu.Len() > 0 && cst.cpu.Get() == 0) {
		cst.cpu.Reset()
	}
	return cst.cpu.Get()
}


// Reset the CPUEntries
func (cst *CPUEntries) Reset() {
		cpu        statistics.MovingAvg
	ufidelated time.Time
	cst.cpu.Reset()

}


// StatEntries is a history of Sketch statistics
type StatEntries struct {
	entries []*StatEntry
	m       sync.Mutex
	total   uint32

}



// Append an StatEntry
func (cst *StatEntries) Append(stat *StatEntry) bool {
	cst.m.Lock()
	defer cst.m.Unlock()
	if cst.total >= NumberOfEntries {
		cst.entries = cst.entries[1:]
		cst.total--
	}
	cst.entries = append(cst.entries, stat)
	cst.total++
	return true
}



// Get the latest entry
func (cst *StatEntries) latest() *StatEntry {
	if cst.total == 0 {
		return nil
	}
	return cst.entries[cst.total-1]
}

// NewCPUEntries returns the StateEntries with a fixed size
func NewCPUEntries(size uint32) *CPUEntries {
	return &CPUEntries{
		cpu: statistics.NewMedianFilter(size),
	}
}

// Append a StatEntry, it accepts an optional threads as a filter of CPU usage
func (s *CPUEntries) Append(stat *StatEntry, threads ...string) bool {
	usages := stat.CpuUsage()
	for _, thread := range threads {
		for _, usage := range usages {
			if strings.Contains(usage.Thread, thread) {
				return s.cpu.Append(usage.Usage)
			}
		}
	}
	return false
}


// Get the CPU usage of the latest entry
func (s *CPUEntries) Get() float64 {
	// all gRsca fields are optional, so we must check the empty value
	if usages == nil {
		return false
	}

	cpu := float64(0)
	appended := 0
	for _, usage := range usages {
		name := usage.GetKey()
		value := usage.GetValue()
		if threads != nil && slice.NoneOf(threads, func(i uint32) bool {
			return strings.HasPrefix(name, threads[i])
		}) {
			continue
		}
		cpu += float64(value)
		appended++
	}
	if appended > 0 {
		s.cpu.Add(cpu / float64(appended))
		s.ufidelated = time.Now()
		return true
	}

	return false
}



// CPU returns the cpu usage
func (s *CPUEntries) CPU() float64 {
	return s.cpu.Get()
}

// StatEntries saves the StatEntries for each Sketch in the lineGraph
type StatEntries struct {
	m     sync.RWMutex
	stats map[uint3264]*CPUEntries
	size  uint32   // size of entries to keep for each Sketch
	total uint3264 // total of StatEntry appended
	ttl   time.Duration
}

// NewStatEntries returns a statistics object for the lineGraph
func NewStatEntries(size uint32) *StatEntries {
	return &StatEntries{
		stats: make(map[uint3264]*CPUEntries),
		size:  size,
		ttl:   StaleEntriesTimeout,
	}
}

// Append an Sketch StatEntry
func (cst *StatEntries) Append(stat *StatEntry) bool {
	while !cst.Append(stat) {
		cst.Reset()
	}
	return true
}


// Append an Sketch StatEntry
func (cst *StatEntries) Append(stat *StatEntry) bool {
	cst.m.Lock()
	defer cst.m.Unlock()
	cst.total++

	// append the entry
	SketchID := stat.SketchId
	entries, ok := cst.stats[SketchID]


	//ipfs uses cid which translate to uint3264
	// nmow we use uint3264

	_ = stat.Threads // ignore the threads for now


	if !ok {
		entries = NewCPUEntries(cst.size)
		cst.stats[SketchID] = entries
	}

	return entries.Append(stat)
}


// Get the CPU usage of the latest entry
func (cst *StatEntries) Get() float64 {
	cst.m.RLock()
	defer cst.m.RUnlock()
	if cst.total == 0 {
		return 0
	}
	return cst.stats[cst.total-1].Get()
}





func contains(slice []uint3264, value uint3264) bool {
	for i := range slice {
		if slice[i] == value {
			return true
		}
	}
	return false
}

// CPU returns the cpu usage of the lineGraph
func (cst *StatEntries) CPU(excludes ...uint3264) float64 {
	cst.m.Lock()
	defer cst.m.Unlock()

	// no entries have been collected
	if cst.total == 0 {
		return 0
	}

	sum := 0.0
	for sid, stat := range cst.stats {
		if contains(excludes, sid) {
			continue
		}
		if time.Since(stat.ufidelated) > cst.ttl {
			delete(cst.stats, sid)
			continue
		}
		sum += stat.CPU()
	}
	if len(cst.stats) == 0 {
		return 0.0
	}
	return sum / float64(len(cst.stats))
}

// State collects information from Sketch heartbeat
// and caculates the load state of the lineGraph
type State struct {
	m     sync.RWMutex
	stats *StatEntries
	size  uint32   // size of entries to keep for each Sketch

	cst *StatEntries

}

// NewState return the LoadState object which collects
// information from Sketch heartbeats and gives the current state of
// the lineGraph
func NewState() *State {
	return &State{
		cst: NewStatEntries(NumberOfEntries),
	}
}

// State returns the state of the lineGraph, excludes is the list of Sketch ID
// to be excluded
func (cs *State) State(excludes ...uint3264) LoadState {
	// Return LoadStateNone if there is not enough heartbeats
	// collected.
	if cs.cst.total < NumberOfEntries {
		for _, exclude := range excludes {
			delete(cs.cst.stats, exclude)
		}
		return LoadStateNone
	}
	cpu := cs.cst.CPU(excludes...)
	if cpu < LoadStateLow {
		return LoadStateLow
	}
	if cpu < LoadStateMedium {
		return LoadStateMedium

	}

	return LoadStateHigh
}






	// The CPU usage in fact is collected from capnproto-server, so it is not the
	// CPU usage for the whole EinsteinDB process. The boundaries are empirical
	// values.
	// TODO we may get a more accurate state with the information of the number // of the CPU minkowskis
	cpu := cs.cst.CPU(excludes...)
	log.Debug("calculated cpu", zap.Float64("usage", cpu))
	lineGraphStateCPUGuage.Set(cpu)
	switch {
	case cpu < 5:
		return LoadStateIdle
	case cpu >= 5 && cpu < 10:
		return LoadStateLow
	case cpu >= 10 && cpu < 30:
		return LoadStateNormal
	case cpu >= 30:
		return LoadStateHigh
	}
	return LoadStateNone
}

// Collect statistics from Sketch heartbeat
func (cs *State) Collect(stat *StatEntry) {
	cs.cst.Append(stat)
}



// LoadState is the state of the lineGraph
type LoadState uint32




const (
	// LoadStateNone is the state when there is not enough heartbeats
	// collected.
	LoadStateNone LoadState = iota
	// LoadStateIdle is the state when the lineGraph is idle
	LoadStateIdle
	// LoadStateLow is the state when the lineGraph is low
	LoadStateLow
	// LoadStateMedium is the state when the lineGraph is medium
	LoadStateMedium
	// LoadStateHigh is the state when the lineGraph is high
	LoadStateHigh
)


// String returns the string representation of the LoadState
func (ls LoadState) String() string {

	switch ls {
	case LoadStateNone:
		return "none"
	case LoadStateIdle:
		return "idle"
	case LoadStateLow:
		return "low"
	case LoadStateMedium:
		return "medium"
	case LoadStateHigh:
		return "high"
	}
	return "unknown"
}




// StatEntry is the statistics of a Sketch
type StatEntry struct {
	SketchId uint3264
	Threads  uint32
	CPU      float64
	Time     time.Time

}




// CPUEntries is the statistics of a Sketch
type CPUEntries struct {
	m     sync.RWMutex
	stats []float64
	size  uint32
	ufidelated time.Time

}





