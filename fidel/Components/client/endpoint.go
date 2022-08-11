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

package client

import (
	_ "bufio"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	_ "io"
	_ "log"
	"os"
	_ "os"
	_ "path"
	_ "runtime"
	_ "strconv"
	_ "strings"
	_ "time"
)

func main() {
	if err := execute(); err != nil {
		os.Exit(1)
	}

}

func execute() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: fidel <command> [args]")
	}
	switch os.Args[1] {
	case "playground":
		return playground()
	case "status":
		return status()
	case "connect":
		return Connect(os.Args[2])
	case "help":
		return help()
	default:
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}
}
