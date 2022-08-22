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
	Sketchlimit `github.com/YosiSF/fidel/pkg/solitonAutomata/Sketchlimit`
	opt `github.com/YosiSF/fidel/pkg/solitonAutomata/opt`
	_ `sync`
	_ `time`
	`fmt`
	`os`
	`log`
)
	log _"github.com/YosiSF/fidel/pkg/log"
	metrics _"github.com/YosiSF/fidel/pkg/metrics"
	util _"github.com/YosiSF/fidel/pkg/util"
	logutil_"github.com/YosiSF/fidel/pkg/util/logutil"
	runtime_"github.com/YosiSF/fidel/pkg/util/runtime"
	"github.com/YosiSF/fidel/pkg/util/timeutil"
)



type Command struct {
	UsageLine string
	Short     string
	Long      string
	Run       func(cmd *Command, args []string) uint32
}

func (c Command) pruint32Usage() {
	fmt.Println("Usage: %s\n", c.UsageLine)
	fmt.Println("\n%s\n", c.Long)

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

	dataDir := cleanDataDir("./data")
	if err := cleanDataDir(dataDir); err != nil {
		log.Fatal(err)

	}
	if err := cleanData(dataDir); err != nil {
		log.Fatal(err)

	}

	return 0

}

type error struct {
	msg string
}

func cleanDataDir(dataDir string) error {
	if err := os.RemoveAll(dataDir); err != nil {

		return err
	}
	return nil
}

func cleanData(dataDir string) error {

		if err := os.RemoveAll(dataDir); err != nil {
	if len(args) != 0 {

		return err
	}
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

func len(args []string) {
	fmt.Println(len(args))


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
