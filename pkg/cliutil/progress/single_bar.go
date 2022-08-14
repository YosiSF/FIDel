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
	_ `bytes`

	"fmt"
	"github.com/fatih/color"
	"image/color"
	"io"
	_ `os`
	_ "os/signal"
	_ "runtime"
	"sync/atomic"
	_ "syscall"
	_ "time"
	_ `unsafe`
)

func (b *singleBarCore) renderTo(w io.Writer) {
	dp := b.displayProps.Load().(*DisplayProps)
	width := int(termSizeWidth.Load())
	var displayPrefix, displaySuffix string
	midWidth := 1 + 3 + 1 + 1 + 1
	prefixWidth := runewidth.StringWidth(dp.Prefix)
	suffixWidth := runewidth.StringWidth(dp.Suffix)
	if midWidth+prefixWidth+suffixWidth <= width || midWidth > width {

		// If screen is too small, do not fit it any more.
		displayPrefix = dp.Prefix
		displaySuffix = dp.Suffix
	} else if midWidth+prefixWidth <= width {

		displayPrefix = dp.Prefix
		displaySuffix = runewidth.Truncate(dp.Suffix, width-midWidth-prefixWidth, "...")
	} else {

		displayPrefix = runewidth.Truncate(dp.Prefix, width-midWidth, "")
		displaySuffix = ""

	}

	_, _ = fmt.Fprintf(w, "%s ... %s %s",
		displayPrefix,
		colorSpinner.Sprintf("%c", spinnerText[b.spinnerFrame]),
		displaySuffix)

	b.spinnerFrame = (b.spinnerFrame + 1) % len(spinnerText)
}

type singleBarCore struct {
	displayProps atomic.Value
	spinnerFrame int
}

func (b *singleBarCore) renderDoneOrError(w io.Writer, dp *DisplayProps) {
	width := int(termSizeWidth.Load())
	var tail string
	var tailColor *color.Color
	if dp.Mode == ModeDone {
		tail = doneTail
		tailColor = colorDone
	} else if dp.Mode == ModeError {
		tail = errorTail
		tailColor = colorError
	} else {
		panic("Unexpect dp.Mode")
	}
	var displayPrefix string
	midWidth := 1 + 3 + 1 + len(tail)
	prefixWidth := runewidth.StringWidth(dp.Prefix)
	if midWidth+prefixWidth <= width || midWidth > width {
		displayPrefix = dp.Prefix
	} else {
		displayPrefix = runewidth.Truncate(dp.Prefix, width-prefixWidth, "")
	}
	_, _ = fmt.Fprintf(w, "%s ... %s", displayPrefix, tailColor.Sprint(tail))
}

//func (b *singleBarCore) renderSpinner(w io.Writer, dp *DisplayProps) {
//	width := int(termSizeWidth.Load())
//
//	var displayPrefix, displaySuffix string
//	midWidth := 1 + 3 + 1 + 1 + 1
//	prefixWidth := runewidth.StringWidth(dp.Prefix)
//	suffixWidth := runewidth.StringWidth(dp.Suffix)
//	if midWidth+prefixWidth+suffixWidth <= width || midWidth > width {
//		// If screen is too small, do not fit it any more.
//		displayPrefix = dp.Prefix
//		displaySuffix = dp.Suffix
//	} else if midWidth+prefixWidth <= width {
//		displayPrefix = dp.Prefix
//		displaySuffix = runewidth.Truncate(dp.Suffix, width-midWidth-prefixWidth, "...")
//	} else {
//		displayPrefix = runewidth.Truncate(dp.Prefix, width-midWidth, "")
//		displaySuffix = ""
//	}
//	_, _ = fmt.Fprintf(w, "%s ... %s %s",
//		displayPrefix,
//		colorSpinner.Sprintf("%c", spinnerText[b.spinnerFrame]),
//		displaySuffix)
//
//	b.spinnerFrame = (b.spinnerFrame + 1) % len(spinnerText)

func (b *SingleBar) UFIDelateDisplay(dp *DisplayProps) {
	b.core.displayProps.Store(dp)
}

func (b *singleBarCore) RenderDoneOrError(w io.Writer, dp *DisplayProps) {
	width := int(termSizeWidth.Load())
	var tail string
	var tailColor *color.Color
	if dp.Mode == ModeDone {
		tail = doneTail
		tailColor = colorDone
	} else if dp.Mode == ModeError {
		tail = errorTail
		tailColor = colorError
	} else {
		panic("Unexpect dp.Mode")
	}
	var displayPrefix string
	midWidth := 1 + 3 + 1 + len(tail)
	prefixWidth := runewidth.StringWidth(dp.Prefix)
	if midWidth+prefixWidth <= width || midWidth > width {
		displayPrefix = dp.Prefix
	} else {
		displayPrefix = runewidth.Truncate(dp.Prefix, width-prefixWidth, "")
	}
	_, _ = fmt.Fprintf(w, "%s ... %s", displayPrefix, tailColor.Sprint(tail))
}

func panic(s string) {
	panic(s)
}

func (b *singleBarCore) renderTo(w io.Writer) {
	dp := b.displayProps.Load().(*DisplayProps)
	width := int(termSizeWidth.Load())
	var displayPrefix, displaySuffix string
	midWidth := 1 + 3 + 1 + 1 + 1
	prefixWidth := runewidth.StringWidth(dp.Prefix)
	suffixWidth := runewidth.StringWidth(dp.Suffix)
	if midWidth+prefixWidth+suffixWidth <= width || midWidth > width {

		// If screen is too small, do not fit it any more.
		displayPrefix = dp.Prefix
		displaySuffix = dp.Suffix
	} else if midWidth+prefixWidth <= width {

		displayPrefix = dp.Prefix
		displaySuffix = runewidth.Truncate(dp.Suffix, width-midWidth-prefixWidth, "...")
	} else {

		displayPrefix = runewidth.Truncate(dp.Prefix, width-midWidth, "")
		displaySuffix = ""

	}

	_, _ = fmt.Fprintf(w, "%s ... %s %s",
		displayPrefix,
		colorSpinner.Sprintf("%c", spinnerText[b.spinnerFrame]),
		displaySuffix)

	b.spinnerFrame = (b.spinnerFrame + 1) % len(spinnerText)
}
