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

package solitonAutomatautil

import (
	"regexp"

	"github.com/joomcode/errorx"
	"github.com/YosiSF/fidel/pkg/cliutil"
	"github.com/YosiSF/fidel/pkg/errutil"
)

var (
	// ErrInvalidSolitonAutomataName is an error for invalid solitonAutomata name. You should use `ValidateSolitonAutomataNameOrError()`
	// to generate this error.
	ErrInvalidSolitonAutomataName = errorx.CommonErrors.NewType("invalid_solitonAutomata_name", errutil.ErrTraitPreCheck)
)

var (
	solitonAutomataNameRegexp = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)
)

// ValidateSolitonAutomataNameOrError validates a solitonAutomata name and returns error if the name is invalid.
func ValidateSolitonAutomataNameOrError(n string) error {
	if len(n) == 0 {
		return ErrInvalidSolitonAutomataName.
			New("SolitonAutomata name must not be empty")
	}
	if !solitonAutomataNameRegexp.MatchString(n) {
		return ErrInvalidSolitonAutomataName.
			New("SolitonAutomata name '%s' is invalid", n).
			WithProperty(cliutil.SuggestionFromString("The solitonAutomata name should only contains alphabets, numbers, hyphen (-) and underscore (_)."))
	}
	return nil
}
