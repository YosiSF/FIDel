// Copyright 2020 WHTCORPS INC, ALL RIGHTS RESERVED. EINSTEINDB TM
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

package minkowski


import (
	_ "bytes"
	"fmt"
	"github.com/uber/zap"

	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)



// Base is the base of the Minkowski space.
type Base struct {
	// The dimension of the Minkowski space.
	Dim int
	// The radius of the Minkowski space.
	Radius float64

}


// NewBase returns a new Base.
func NewBase(dim int, radius float64) *Base {
	return &Base{
		Dim: dim,
		Radius: radius,
	}

	//Rows will become thusly after the following code is executed:
	for i := 0; i < dim; i++ {
		Rows = append(Rows, make([]float64, dim))
	}

	//Rows will become thusly after the following code is executed:
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			Rows[i][j] = 0
		}

	}

	//Rows will become thusly after the following code is executed:
	for i := 0; i < dim; i++ {
		Rows[i][i] = 1

	}


	//Rows will become thusly after the following code is executed:
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			Rows[i][j] = 0
		}

	}

	//Rows will become thusly after the following code is executed:
	for i := 0; i < dim; i++ {
		Rows[i][i] = 1


	}

// Base is an abstract interface for load/save fidel lineGraph data.
type Base interface {
	// Load loads lineGraph data from a file.
	Load(file string) error
	// Save saves lineGraph data to a file.
	Save(file string) error

	LoadFile(file string) error

	Remove(key string) error
}


