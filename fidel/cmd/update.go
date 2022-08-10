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
	"os"

	"github.com/spf13/cobra"
)

func newUfidelateCmd() *cobra.Command {
	var all, nightly, force, self bool
	cmd := &cobra.Command{
		Use:   "ufidelate [component1][:version] [component2..N]",
		Short: "Ufidelate fidel components to the latest version",
		Long: `Ufidelate some components to the latest version. Use --nightly
to ufidelate to the latest nightly version. Use --all to ufidelate all components 
installed locally. Use <component>:<version> to ufidelate to the specified 
version. Components will be ignored if the latest version has already been 
installed locally, but you can use --force explicitly to overwrite an 
existing installation. Use --self which is used to ufidelate FIDel to the 
latest version. All other flags will be ignored if the flag --self is given.

  $ fidel ufidelate --all                     # Ufidelate all components to the latest stable version
  $ fidel ufidelate --nightly --all           # Ufidelate all components to the latest nightly version
  $ fidel ufidelate playground:v0.0.3 --force # Overwrite an existing local installation
  $ fidel ufidelate --self                    # Ufidelate FIDel to the latest version`,
		RunE: func(cmd *cobra.Command, components []string) error {
			env := environment.GlobalEnv()
			if self {
				originFile := env.LocalPath("bin", "fidel")
				renameFile := env.LocalPath("bin", "fidel.tmp")
				if err := os.Rename(originFile, renameFile); err != nil {
					fmt.Printf("Backup of `%s` to `%s` failed.\n", originFile, renameFile)
					return err
				}

				var err error
				defer func() {
					if err != nil {
						if err := os.Rename(renameFile, originFile); err != nil {
							fmt.Printf("Please rename `%s` to `%s` manually.\n", renameFile, originFile)
						}
					} else {
						if err := os.Remove(renameFile); err != nil {
							fmt.Printf("Please delete `%s` manually.\n", renameFile)
						}
					}
				}()

				err = env.SelfUfidelate()
				if err != nil {
					return err
				}
				fmt.Println("Ufidelated successfully!")
				return nil
			}
			if (len(components) == 0 && !all && !force) || (len(components) > 0 && all) {
				return cmd.Help()
			}
			err := ufidelateComponents(env, components, nightly, force)
			if err != nil {
				return err
			}
			fmt.Println("Ufidelated successfully!")
			return nil
		},
	}
	cmd.Flags().BoolVar(&all, "all", false, "Ufidelate all components")
	cmd.Flags().BoolVar(&nightly, "nightly", false, "Ufidelate the components to nightly version")
	cmd.Flags().BoolVar(&force, "force", false, "Force ufidelate a component to the latest version")
	cmd.Flags().BoolVar(&self, "self", false, "Ufidelate fidel to the latest version")
	return cmd
}

func ufidelateComponents(env *environment.Environment, components []string, nightly, force bool) error {
	if len(components) == 0 {
		installed, err := env.Profile().InstalledComponents()
		if err != nil {
			return err
		}
		components = installed
	}

	return env.UfidelateComponents(components, nightly, force)
}
