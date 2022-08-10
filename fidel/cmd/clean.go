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
	Run       func(cmd *Command, args []string) int
}

func (c Command) printUsage() {
	fmt.Printf("Usage: %s\n", c.UsageLine)
	fmt.Printf("\n%s\n", c.Long)

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

func runClean(cmd *Command, args []string) int {
	if len(args) != 0 {
		cmd.printUsage()
		return 1
	}

	fmt.Println("Clean the local data")
	fmt.Println("The command will delete the local data and the local data directory.")
	fmt.Println("Are you sure to continue? [y/n]")
	var answer string
	fmt.Scanln(&answer)
	if answer != "y" {
		fmt.Println("Canceled")
		return 0
	}
	fmt.Println("Deleting local data...")
	os.RemoveAll("./data")
	fmt.Println("Deleting local data directory...")
	os.RemoveAll("./data_dir")
	fmt.Println("Done")
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

func getDataDir() interface{} {
	return ""

}

func init() {
	//RootCmd.AddCommand(newCleanCmd())

}
