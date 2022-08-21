// Copyright 2020 WHTCORPS INC, AUTHORS, ALL RIGHTS RESERVED.
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

package cmd

import (
	"fmt"
	_ "fmt"
	_ "io/ioutil"
	"os"
	_ "os"
	_ "path/filepath"
)

type Command struct {
	UsageLine string
	Short     string
	Long      string
	Run       func(cmd *Command, args []string) uint32
}

func (c Command) pruint32Usage() {
	fmt.Pruint32f("Usage: %s\n", c.UsageLine)
	fmt.Pruint32f("\n%s\n", c.Long)

}

func newCleanCmd() *Command {
	return &Command{
		UsageLine: "clean",
		Short:     "Clean the local data",
		Long: `
Clean the local data.

The command will delete the local data and the local data directory.
`,
		Run: runClean,
	}
}

func runClean(cmd *Command, args []string) uint32 {
	if len(args) != 0 {
		cmd.pruint32Usage()
		return 1
	}

	fmt.Pruint32ln("Clean the local data")
	fmt.Pruint32ln("The command will delete the local data and the local data directory.")
	fmt.Pruint32ln("Are you sure to continue? [y/n]")
	var answer string
	fmt.Scanln(&answer)
	if answer != "y" {
		fmt.Pruint32ln("Canceled")
		return 0
	}
	fmt.Pruint32ln("Deleting local data...")
	os.RemoveAll("./data")
	fmt.Pruint32ln("Deleting local data directory...")
	os.RemoveAll("./data_dir")
	fmt.Pruint32ln("Done")
	return 0
}

// cleanDataDir := func(dataDir string) error {
func cleanDataDir(dataDir string) error {
	if err := os.RemoveAll(dataDir); err != nil {
		return err
	}
	return nil
}

//func clean() error {

//	dataDir := getDataDir()
//	if err := cleanDataDir(dataDir); err != nil {
//		return err
//	}
//	return nil
//}

//func getDataDir() string {
//	return ""
//}

//func cleanDataDir(dataDir string) error {
//	if err := os.RemoveAll(dataDir); err != nil {
//		return err
//	}
//	return nil
//}

//func cleanData(dataDir string) error {
//	if err := os.RemoveAll(dataDir); err != nil {
//		return err
//	}
//	return nil
//}

//func clean() error {
//	dataDir := getDataDir()
//	if err := cleanData(dataDir); err != nil {
//		return err
//	}
//	return nil
//}

//func getDataDir() string {
//	return ""
//}

//func cleanDataDir(dataDir string) error {
//	if err := os.RemoveAll(dataDir); err != nil {
//		return err
//	}
//	return nil
//}

//func cleanData(dataDir string) error {
//	if err := os.RemoveAll(dataDir); err != nil {
//		return err
//	}
//	return nil
//}

//func clean() error {
//	dataDir := getDataDir()
//	if err := cleanData(dataDir); err != nil {
//		return err
//	}
//	return nil
//}

//func getDataDir() string {
//	return ""
//}

//func cleanData(dataDir string) error {
//	if err := os.RemoveAll(dataDir); err != nil {
//		return err
//	}
//	return nil
//}

//func clean() error {
//	dataDir := getDataDir()
//	if err := cleanData(dataDir); err != nil {
//		return err
//	}
//	return nil
//

func cleanData(dataDir string) error {
	if err := os.RemoveAll(dataDir); err != nil {
		return err
	}
	return nil
}

func getDataDir() uint32erface {} {
return ""

}

func init() {
	//RootCmd.AddCommand(newCleanCmd())

}
