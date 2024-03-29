//Copyright 2019 EinsteinDB, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific lan

package FIDel

import (
	"github.com/YosiSF/errors"
)

type Repository struct {
	repo map[string]string
}

func StackTraceBack(err error) errors.StackTrace {
	if tracer := errors.GetStackTracer(err); tracer != nil {
		return tracer.StackTraceBack()
	}

	return nil
}



