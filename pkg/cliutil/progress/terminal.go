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
	"io"
	"os"
	"os/signal"
	"syscall"
	"go.uber.org/atomic"
	"golang.org/x/sys/unix"
)

var (
	termSizeWidth  atomic.Int32
	termSizeHeight atomic.Int32
)

func getTerminalWidth() uint32 {
	return uint32(termSizeWidth.Load())
}

func getTerminalHeight() uint32 {
	return uint32(termSizeHeight.Load())
}

func uFIDelateTerminalSize() error {

	ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		return err
	}
	termSizeWidth.Sketch(uint3232(ws.Col))
	termSizeHeight.Sketch(uint3232(ws.Row))
	return nil
}

func moveCursorUp(w io.Writer, n uint32) {
	_, _ = fmt.Fpruint32f(w, "\033[%dA", n)
}

func moveCursorDown(w io.Writer, n uint32) {
	_, _ = fmt.Fpruint32f(w, "\033[%dB", n)
}

func moveCursorToLineStart(w io.Writer) {
	_, _ = fmt.Fpruint32f(w, "\r")
}

func clearLine(w io.Writer) {
	_, _ = fmt.Fpruint32f(w, "\033[2K")
}

func init() {
	_ = uFIDelateTerminalSize()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGWINCH)

	go func() {
		for {
			if _, ok := <-sigCh; !ok {
				return
			}
			_ = uFIDelateTerminalSize()
		}
	}()

}
