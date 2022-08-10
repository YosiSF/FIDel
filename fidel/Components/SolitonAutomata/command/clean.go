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
	operator "github.com/YosiSF/fidel/pkg/solitonAutomata/operation"
	"github.com/spf13/cobra"
)

func newCleanCmd() *cobra.Command {
	cleanOpt := operator.Options{}
	cleanALl := false

	cmd := &cobra.Command{
		Use: "clean <solitonAutomata-name>",
		Long: `Cleanup a specified solitonAutomata without destroying it.
You can retain some nodes and roles data when cleanup the solitonAutomata, eg:
    $ fidel solitonAutomata clean <solitonAutomata-name> --all
    $ fidel solitonAutomata clean <solitonAutomata-name> --log
    $ fidel solitonAutomata clean <solitonAutomata-name> --data
    $ fidel solitonAutomata clean <solitonAutomata-name> --all --ignore-role prometheus
    $ fidel solitonAutomata clean <solitonAutomata-name> --all --ignore-node 172.16.13.11:9000
    $ fidel solitonAutomata clean <solitonAutomata-name> --all --ignore-node 172.16.13.12`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			if cleanALl {
				cleanOpt.CleanuFIDelata = true
				cleanOpt.CleanupLog = true
			}

			if !(cleanOpt.CleanuFIDelata || cleanOpt.CleanupLog) {
				return cmd.Help()
			}

			return manager.CleanSolitonAutomata(solitonAutomataName, gOpt, cleanOpt, skipConfirm)
		},
	}

	cmd.Flags().StringArrayVar(&cleanOpt.RetainDataNodes, "ignore-node", nil, "Specify the nodes or hosts whose data will be retained")
	cmd.Flags().StringArrayVar(&cleanOpt.RetainDataRoles, "ignore-role", nil, "Specify the roles whose data will be retained")
	cmd.Flags().BoolVar(&cleanOpt.CleanuFIDelata, "data", false, "Cleanup data")
	cmd.Flags().BoolVar(&cleanOpt.CleanupLog, "log", false, "Cleanup log")
	cmd.Flags().BoolVar(&cleanALl, "all", false, "Cleanup both log and data")

	return cmd
}
