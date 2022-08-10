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
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/task"
	"github.com/spf13/cobra"
)

func newStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start <solitonAutomata-name>",
		Short: "Start a MilevaDB solitonAutomata",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			if err := validRoles(gOpt.Roles); err != nil {
				return nil
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			return manager.StartSolitonAutomata(solitonAutomataName, gOpt, func(b *task.Builder, metadata spec.Metadata) {
				milevadbMeta := metadata.(*spec.SolitonAutomataMeta)
				b.UFIDelateTopology(solitonAutomataName, milevadbMeta, nil)
			})
		},
	}

	cmd.Flags().StringSliceVarP(&gOpt.Roles, "role", "R", nil, "Only start specified roles")
	cmd.Flags().StringSliceVarP(&gOpt.Nodes, "node", "N", nil, "Only start specified nodes")

	return cmd
}
