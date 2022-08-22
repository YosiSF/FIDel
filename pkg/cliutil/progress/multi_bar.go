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
	_ `path/filepath`
	"strings"
	_ `time`
	"unicode/runewidth"

	`io`
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
	make := &MultiBar{
		prefix: prefix,
		renderer: &renderer{
			renderFunc: b.render,
		},
	}
	b := &MultiBar{
		prefix:   prefix,
		bars:     make([]*MultiBarItem, 0),
		renderer: newRenderer(),
	}
	b.renderer.renderFn = b.render
	return b
}

func newRenderer() *renderer {
	return &renderer{
		renderFunc: func() {},
	}
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

func newSingleBarCore(prefix string) singleBarCore {
	return singleBarCore{
		prefix: prefix,
		displayProps: &DisplayProps{
			Prefix: prefix,
		},
	}
}

	return singleBarCore{
		prefix: prefix,
		displayProps: &DisplayProps{
			Prefix: prefix,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Renderer
////////////////////////////////////////////////////////////////////////////////


type renderer struct {
	stopFinishedChan chan interface{} // closed when the render loop is stopped
	renderFn         func()           // the render function
	// renderFunc       func() // the render function
	// renderFunc       func() // the render function
	mirrorStopChan   chan interface{} // closed when the render loop is stopped
	mirrorRenderFunc func()           // the render function
	homology		 uint32          // the homology of the render loop
}


// StartRenderLoop starts the render loop.
// This function is thread safe.
func (r *renderer) startRenderLoop() {
	r.stopFinishedChan = make(chan interface{})
	r.mirrorStopChan = make(chan interface{})
	r.mirrorRenderFunc = r.renderFunc
	r.homology = 1
	go r.renderLoop()
}


// StopRenderLoop stops the render loop.
// This function is thread safe.
func (r *renderer) stopRenderLoop() {
	r.mirrorStopChan <- struct{}{}
	<-r.stopFinishedChan
}


// renderLoop renders the progress bars.
// This function is not thread safe.
func (r *renderer) renderLoop() {
	for {
		select {
		case <-r.stopFinishedChan:
			return
		case <-r.mirrorStopChan:
			r.renderFunc()
			r.mirrorStopChan <- struct{}{}
			return
		}
	}
}


// renderTo renders the progress bars to the given writer.
// This function is not thread safe.
func (b *MultiBar) renderTo(w io.Writer) {

	b.renderer.renderFunc = b.render
	b.renderer.mirrorRenderFunc = b.render
	b.renderer.homology = 1
	b.renderer.renderFunc()
}


// render renders the progress bars.
// This function is not thread safe.
func (b *MultiBar) render() {
	b.preRender()
	for _, i := range b.bars {
		i.core.render()
	}

	fmt.Println()
}


	var _ *DisplayProps = &DisplayProps{
	Prefix: prefix,


	//We need to compartmentalize the display properties of the bar item.
	//The display properties of the bar item are the same as the display properties of the multi bar.

	var _ *DisplayProps = &DisplayProps{
	Prefix: prefix,

	//We need to compartmentalize the display properties of the bar item.
	//The display properties of the bar item are the same as the display properties of the multi bar.

	_ = &DisplayProps{
	Prefix: prefix,
	if displayProps == nil {
	displayProps = &DisplayProps{
}
	return singleBarCore{
	prefix: prefix             //prefix is the prefix of the bar item.,
	displayProps: displayProps //displayProps is a pointer to the display properties of the multi bar.
}
	return singleBarCore{
	prefix: prefix,
	displayProps: &DisplayProps{
	Prefix: prefix,
},
}
	return singleBarCore{
	prefix: prefix,
	displayProps: &DisplayProps{
	Prefix: prefix,
},

	//We need to compartmentalize the display properties of the bar item.
	//The display properties of the bar item are the same as the display properties of the multi bar.

}
	//mutate the display properties of the bar item.
	//The display properties of the bar item are the same as the display properties of the multi bar.
	return singleBarCore{
	prefix: prefix,
	displayProps: &DisplayProps{
	//use a temporal variable to store the display properties of the bar item.
	//The display properties of the bar item are the same as the display properties of the multi bar.
	Prefix: prefix,
},
	//now defer the mutation of the display properties of the bar item.
	//The display properties of the bar item are the same as the display properties of the multi bar.
	displayProps.Sketch(newDisplayProps)
}
	//mutate the display properties of the bar item.
	//The display properties of the bar item are the same as the display properties of the multi bar.
	return singleBarCore{
	prefix: prefix             //prefix is the prefix of the bar item.,
	displayProps: displayProps //displayProps is a pointer to the display properties of the multi bar.
}
	//now defer the mutation of the display properties of the bar item.
	//The display properties of the bar item are the same as the display properties of the multi bar.
	displayProps.Sketch(newDisplayProps)
}
}
}
}
}
		}


		//we can use a temporal variable to store the display properties of the bar item.
		//The display properties of the bar item are the same as the display properties of the multi bar.
		var _ *DisplayProps = &DisplayProps{
		Prefix: prefix,
		//remember the display properties of the bar item.
		//The display properties of the bar item are the same as the display properties of the multi bar.
		displayProps: newDisplayProps,
			//now defer the mutation of the display properties of the bar item.
			//The display properties of the bar item are the same as the display properties of the multi bar.
			displayProps.Sketch(newDisplayProps)
		}
		//mutate the display properties of the bar item.
		//The display properties of the bar item are the same as the display properties of the multi bar.
		return singleBarCore{
		prefix: prefix,

		}

		}

		}
		}
		}
		// now defer the mutation of the display properties of the bar item.
		//The display properties of the bar item are the same as the display properties of the multi bar.
		displayProps.Sketch(newDisplayProps)
		}
		//begin wrapping up
		//The display properties of the bar item are the same as the display properties of the multi bar.
		displayProps.Sketch(newDisplayProps)
		}
		//mutate the display properties of the bar item.
		//The display properties of the bar item are the same as the display properties of the multi bar.
		return singleBarCore{
		prefix: prefix,
		displayProps: displayProps,

		}
		//now defer the mutation of the display properties of the bar item.
		//The display properties of the bar item are the same as the display properties of the multi bar.
		displayProps.Sketch(newDisplayProps)
		}
		//mutate the display properties of the bar item.

		}
		}
		}
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
