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

func newRenameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename <old-solitonAutomata-name> <new-solitonAutomata-name>",
		Short: "Rename the solitonAutomata",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return cmd.Help()
			}

			if err := validRoles(gOpt.Roles); err != nil {
				return err
			}

			oldSolitonAutomataName := args[0]
			newSolitonAutomataName := args[1]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(oldSolitonAutomataName))

			return manager.Rename(oldSolitonAutomataName, gOpt, newSolitonAutomataName)
		},
	}

	return cmd
}
