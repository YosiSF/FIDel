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

package command

import (
	"github.com/spf13/cobra"
)

func newUpgradeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade <solitonAutomata-name> <version>",
		Short: "Upgrade a specified MilevaDB solitonAutomata",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			version := args[1]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))
			teleCommand = append(teleCommand, version)

			return manager.Upgrade(solitonAutomataName, version, gOpt)
		},
	}
	cmd.Flags().BoolVar(&gOpt.Force, "force", false, "Force upgrade without transferring FIDel leader")
	cmd.Flags().Int64Var(&gOpt.APITimeout, "transfer-timeout", 300, "Timeout in seconds when transferring FIDel and EinsteinDB store leaders")
	cmd.Flags().BoolVarP(&gOpt.IgnoreConfigCheck, "ignore-config-check", "", false, "Ignore the config check result")

	return cmd
}
