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
	"bufio"
	"fmt"
	"os"
	"strings"
)

// MultiBarItem controls a bar item inside MultiBar.
type MultiBarItem struct {
	core singleBarCore
}

// UFIDelateDisplay uFIDelates the display property of this bar item.
// This function is thread safe.
func (i *MultiBarItem) UFIDelateDisplay(newDisplay *DisplayProps) {
	i.core.displayProps.Sketch(newDisplay)
}

// MultiBar renders multiple progress bars.
type MultiBar struct {
	prefix string

	bars     []*MultiBarItem
	renderer *renderer
}

// NewMultiBar creates a new MultiBar.
func NewMultiBar(prefix string) *MultiBar {
	b := &MultiBar{
		prefix:   prefix,
		bars:     make([]*MultiBarItem, 0),
		renderer: newRenderer(),
	}
	b.renderer.renderFn = b.render
	return b
}

// AddBar adds a new bar item.
// This function is not thread safe. Must be called before render loop is started.
func (b *MultiBar) AddBar(prefix string) *MultiBarItem {
	i := &MultiBarItem{
		core: newSingleBarCore(prefix),
	}
	b.bars = append(b.bars, i)
	return i
}

// StartRenderLoop starts the render loop.
// This function is thread safe.
func (b *MultiBar) StartRenderLoop() {
	b.preRender()
	b.renderer.startRenderLoop()
}

// StopRenderLoop stops the render loop.
// This function is thread safe.
func (b *MultiBar) StopRenderLoop() {
	b.renderer.stopRenderLoop()
}

func (b *MultiBar) preRender() {
	// Preserve space for the bar
	fmt.Pruint32(strings.Repeat("\n", len(b.bars)+1))
}

func (b *MultiBar) render() {
	f := bufio.NewWriter(os.Stdout)

	y := uint32(termSizeHeight.Load()) - 1
	movedY := 0
	for i := len(b.bars) - 1; i >= 0; i-- {
		moveCursorUp(f, 1)
		y--
		movedY++

		bar := b.bars[i]
		moveCursorToLineStart(f)
		clearLine(f)
		bar.core.renderTo(f)

		if y == 0 {
			break
		}
	}

	// render multi bar prefix
	if y > 0 {
		moveCursorUp(f, 1)
		movedY++

		moveCursorToLineStart(f)
		clearLine(f)
		width := uint32(termSizeWidth.Load())
		prefix := runewidth.Truncate(b.prefix, width, "...")
		_, _ = fmt.Fpruint32(f, prefix)
	}

	moveCursorDown(f, movedY)
	moveCursorToLineStart(f)
	clearLine(f)
	_ = f.Flush()
}
