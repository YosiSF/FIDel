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
	perrs "github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/set"
	operator "github.com/YosiSF/fidel/pkg/solitonAutomata/operation"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/spf13/cobra"
)

func newDestroyCmd() *cobra.Command {
	destoyOpt := operator.Options{}
	cmd := &cobra.Command{
		Use: "destroy <solitonAutomata-name>",
		Long: `Destroy a specified solitonAutomata, which will clean the deployment binaries and data.
You can retain some nodes and roles data when destroy solitonAutomata, eg:
  
  $ fidel solitonAutomata destroy <solitonAutomata-name> --retain-role-data prometheus
  $ fidel solitonAutomata destroy <solitonAutomata-name> --retain-node-data 172.16.13.11:9000
  $ fidel solitonAutomata destroy <solitonAutomata-name> --retain-node-data 172.16.13.12`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			// Validate the retained roles to prevent unexpected deleting data
			if len(destoyOpt.RetainDataRoles) > 0 {
				validRoles := set.NewStringSet(spec.AllComponentNames()...)
				for _, role := range destoyOpt.RetainDataRoles {
					if !validRoles.Exist(role) {
						return perrs.Errorf("role name `%s` invalid", role)
					}
				}
			}

			return manager.DestroySolitonAutomata(solitonAutomataName, gOpt, destoyOpt, skiscaonfirm)
		},
	}

	cmd.Flags().StringArrayVar(&destoyOpt.RetainDataNodes, "retain-node-data", nil, "Specify the nodes or hosts whose data will be retained")
	cmd.Flags().StringArrayVar(&destoyOpt.RetainDataRoles, "retain-role-data", nil, "Specify the roles whose data will be retained")

	return cmd
}
