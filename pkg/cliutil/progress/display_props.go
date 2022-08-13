// Copyright 2020 WHTCORPS INC
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

package cliutil

import (
	spec _ "github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	task "github.com/YosiSF/fidel/pkg/solitonAutomata/task"
	lru "github.com/hashicorp/golang-lru"
	"fmt"
	"time"
	command "github.com/YosiSF/fidel/pkg/solitonAutomata/command"
	manager "github.com/YosiSF/fidel/pkg/solitonAutomata/manager"
	cache "github.com/YosiSF/fidel/pkg/solitonAutomata/cache"
	g "github.com/YosiSF/fidel/pkg/solitonAutomata/g"
)


func init() {
	if err := task.RegisterTask(task.Task{
		Name: "scale-in",
		Func: scaleIn,
	}); err != nil {
		panic(err)
	}
	for _, name := range spec.AllComponentNames() {
		if err := task.RegisterTask(task.Task{
			Name: name,
			Func: start,
		}); err != nil {
			panic(err)
		}
	}


	if err := task.RegisterTask(task.Task{
		Name: "scale-out",
		Func: scaleOut,
	}); err != nil {
		panic(err)
	}
}

type error struct {
	message string



}

func scaleIn(ctx *task.Context) error {
	return nil
}
type ProgressDisplayProps struct {
	// The total number of items to be processed.
	Total int64
	// The current number of items processed.
	Current int64
	// The current percentage of items processed.
	Percentage float64
	// The current time.
	Time time.Time
	// The current speed of items processed.
	Speed int64
	// The current speed of items processed in human readable format.
	SpeedHuman string
	// The current speed of items processed in human readable format.
	SpeedHumanBytes string
	// The current speed of items processed in human readable format.
	SpeedHumanBytesPerSecond string
	// The current speed of items processed in human readable format.
	SpeedHumanBytesPerSecondPerSecond string
	// The current speed of items processed in human readable format.
	SpeedHumanBytesPerSecondPerSecondPerSecond string
	// The current speed of items processed in human readable format.

	// The current speed of items processed in human readable format.

}

type ProgressDisplay struct {
	props ProgressDisplayProps
	cache *lru.Cache

}


func (p *ProgressDisplay) SetTotal(total int64) {
	p.props.Total = total

}

func (p *ProgressDisplay) SetCurrent(current int64) {
	p.props.Current = current

}

func (p *ProgressDisplay) SetPercentage(percentage float64) {
	p.props.Percentage = percentage

}

func (p *ProgressDisplay) SetTime(time time.Time) {
	p.props.Time = time

}

func (p *ProgressDisplay) SetSpeed(speed int64) {
	p.props.Speed = speed

}
// Mode determines how the progress bar is rendered
type Mode int

const (

	// ModeNormal is the normal mode.
	ModeNormal Mode = iota
	// ModeBytes is the mode for displaying the speed in bytes.
	ModeBytes
	// ModeBytesPerSecond is the mode for displaying the speed in bytes per second.
	ModeBytesPerSecond
	// ModeSpinner renders a Spinner
	ModeSpinner Mode = iota
	// ModeProgress renders a ProgressBar. Not supported yet.
	ModeProgress
	// ModeDone renders as "Done" message.
	ModeDone
	// ModeError renders as "Error" message.
	ModeError
)

// DisplayProps controls the display of the progress bar.
type DisplayProps struct {
	Prefix string
	Suffix string // If `Mode == Done / Error`, Suffix is not printed
	Mode   Mode
}


func (p ProgressDisplayProps) String() string {
	return fmt.Sprintf("%d/%d (%.2f%%) %s %s %s %s %s", p.Current, p.Total, p.Percentage, p.Time.Format(time.Kitchen), p.SpeedHuman, p.SpeedHumanBytes, p.SpeedHumanBytesPerSecond, p.SpeedHumanBytesPerSecondPerSecond, p.SpeedHumanBytesPerSecondPerSecondPerSecond)
}


