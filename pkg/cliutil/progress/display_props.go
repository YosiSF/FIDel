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
	"fmt"
	"time"
)

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

func (p ProgressDisplayProps) String() string {
	return fmt.Sprintf("%d/%d (%.2f%%) %s %s %s %s %s", p.Current, p.Total, p.Percentage, p.Time.Format(time.Kitchen), p.SpeedHuman, p.SpeedHumanBytes, p.SpeedHumanBytesPerSecond, p.SpeedHumanBytesPerSecondPerSecond, p.SpeedHumanBytesPerSecondPerSecondPerSecond)

}

// Mode determines how the progress bar is rendered
type Mode int

const (
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
