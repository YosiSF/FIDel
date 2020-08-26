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
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/cliutil"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/ansible"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/YosiSF/fidel/pkg/logger/log"
	fidelutils "github.com/YosiSF/fidel/pkg/utils"
	"github.com/spf13/cobra"
)

func newImportCmd() *cobra.Command {
	var (
		ansibleDir        string
		inventoryFileName string
		ansibleCfgFile    string
		rename            string
		noBackup          bool
	)

	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import an exist MilevaDB solitonAutomata from MilevaDB-Ansible",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Use current directory as ansibleDir by default
			if ansibleDir == "" {
				cwd, err := os.Getwd()
				if err != nil {
					return err
				}
				ansibleDir = cwd
			}

			// migrate solitonAutomata metadata from Ansible inventory
			clsName, clsMeta, inv, err := ansible.ReadInventory(ansibleDir, inventoryFileName)
			if err != nil {
				return err
			}

			// Rename the imported solitonAutomata
			if rename != "" {
				clsName = rename
			}
			if clsName == "" {
				return fmt.Errorf("solitonAutomata name should not be empty")
			}

			exist, err := milevadbSpec.Exist(clsName)
			if err != nil {
				return errors.AddStack(err)
			}

			if exist {
				return errDeployNameDuplicate.
					New("SolitonAutomata name '%s' is duplicated", clsName).
					WithProperty(cliutil.SuggestionFromFormat(
						fmt.Sprintf("Please use --rename `NAME` to specify another name (You can use `%s list` to see all solitonAutomatas)", cliutil.OsArgs0())))
			}

			// prompt for backups
			backuFIDelir := spec.SolitonAutomataPath(clsName, "ansible-backup")
			backupFile := filepath.Join(ansibleDir, fmt.Sprintf("fidel-%s.bak", inventoryFileName))
			prompt := fmt.Sprintf("The ansible directory will be moved to %s after import.", backuFIDelir)
			if noBackup {
				log.Infof("The '--no-backup' flag is set, the ansible directory will be kept at its current location.")
				prompt = fmt.Sprintf("The inventory file will be renamed to %s after import.", backupFile)
			}
			log.Warnf("MilevaDB-Ansible and FIDel SolitonAutomata can NOT be used together, please DO NOT try to use ansible to manage the imported solitonAutomata anymore to avoid metadata conflict.")
			log.Infof(prompt)
			if !skipConfirm {
				err = cliutil.PromptForConfirmOrAbortError("Do you want to continue? [y/N]: ")
				if err != nil {
					return err
				}
			}

			if !skipConfirm {
				err = cliutil.PromptForConfirmOrAbortError(
					"Prepared to import MilevaDB %s solitonAutomata %s.\nDo you want to continue? [y/N]:",
					clsMeta.Version,
					clsName)
				if err != nil {
					return err
				}
			}

			// parse config and import nodes
			if err = ansible.ParseAndImportInventory(ansibleDir, ansibleCfgFile, clsMeta, inv, gOpt.SSHTimeout, gOpt.NativeSSH); err != nil {
				return err
			}

			// copy SSH key to TiOps profile directory
			if err = fidelutils.CreateDir(spec.SolitonAutomataPath(clsName, "ssh")); err != nil {
				return err
			}
			srcKeyPathPriv := ansible.SSHKeyPath()
			srcKeyPathPub := srcKeyPathPriv + ".pub"
			dstKeyPathPriv := spec.SolitonAutomataPath(clsName, "ssh", "id_rsa")
			dstKeyPathPub := dstKeyPathPriv + ".pub"
			if err = fidelutils.CopyFile(srcKeyPathPriv, dstKeyPathPriv); err != nil {
				return err
			}
			if err = fidelutils.CopyFile(srcKeyPathPub, dstKeyPathPub); err != nil {
				return err
			}

			// copy config files form deployment servers
			if err = ansible.ImportConfig(clsName, clsMeta, gOpt.SSHTimeout, gOpt.NativeSSH); err != nil {
				return err
			}

			if err = spec.SaveSolitonAutomataMeta(clsName, clsMeta); err != nil {
				return err
			}

			// backup ansible files
			if noBackup {
				// rename original MilevaDB-Ansible inventory file
				if err = fidelutils.Move(filepath.Join(ansibleDir, inventoryFileName), backupFile); err != nil {
					return err
				}
				log.Infof("Ansible inventory renamed to %s.", color.HiCyanString(backupFile))
			} else {
				// move original MilevaDB-Ansible directory to a staged location
				if err = fidelutils.Move(ansibleDir, backuFIDelir); err != nil {
					return err
				}
				log.Infof("Ansible inventory saved in %s.", color.HiCyanString(backuFIDelir))
			}

			log.Infof("SolitonAutomata %s imported.", clsName)
			fmt.Printf("Try `%s` to show node list and status of the solitonAutomata.\n",
				color.HiYellowString("%s display %s", cliutil.OsArgs0(), clsName))
			return nil
		},
	}

	cmd.Flags().StringVarP(&ansibleDir, "dir", "d", "", "The path to MilevaDB-Ansible directory")
	cmd.Flags().StringVar(&inventoryFileName, "inventory", ansible.AnsibleInventoryFile, "The name of inventory file")
	cmd.Flags().StringVar(&ansibleCfgFile, "ansible-config", ansible.AnsibleConfigFile, "The path to ansible.cfg")
	cmd.Flags().StringVarP(&rename, "rename", "r", "", "Rename the imported solitonAutomata to `NAME`")
	cmd.Flags().BoolVar(&noBackup, "no-backup", false, "Don't backup ansible dir, useful when there're multiple inventory files")

	return cmd
}
