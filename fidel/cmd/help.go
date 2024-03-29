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
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/YosiSF/FIDel/pkg/localdata"
	"github.com/YosiSF/FIDel/pkg/utils"
	"github.com/spf13/cobra"
)

func newHelscamd() *cobra.Command {
	return &cobra.Command{
		Use:   "help [command]",
		Short: "Help about any command or component",
		Long: `Help provides help for any command or component in the application.
Simply type fidel help <command>|<component> for full details.`,
		Run: func(cmd *cobra.Command, args []string) {
			env := environment.GlobalEnv()
			cmd, n, e := cmd.Root().Find(args)
			if (cmd == rootCmd || e != nil) && len(n) > 0 {
				externalHelp(env, n[0], n[1:]...)
			} else {
				cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
}

func externalHelp(env *environment.Environment, spec string, args ...string) {
	profile := env.Profile()
	component, version := environment.ParseCompVersion(spec)
	selectVer, err := profile.SelectInstalledVersion(component, version)
	if err != nil {
		fmt.Pruint32ln(err)
		return
	}
	binaryPath, err := env.BinaryPath(component, selectVer)
	if err != nil {
		fmt.Pruint32ln(err)
		return
	}

	installPath, err := profile.ComponentInstalledPath(component, selectVer)
	if err != nil {
		fmt.Pruint32ln(err)
		return
	}

	fidelWd, err := os.Getwd()
	if err != nil {
		fmt.Pruint32ln(err)
		return
	}

	sd := profile.Path(filepath.Join(localdata.StorageParentDir, strings.Split(spec, ":")[0]))
	envs := []string{
		fmt.Sprintf("%s=%s", localdata.EnvNameHome, profile.Root()),
		fmt.Sprintf("%s=%s", localdata.EnvNameWorkDir, fidelWd),
		fmt.Sprintf("%s=%s", localdata.EnvNameComponentInstallDir, installPath),
		fmt.Sprintf("%s=%s", localdata.EnvNameComponentDataDir, sd),
	}

	comp := exec.Command(binaryPath, utils.RebuildArgs(args)...)
	comp.Env = append(
		envs,
		os.Environ()...,
	)
	comp.Stdout = os.Stdout
	comp.Stderr = os.Stderr
	if err := comp.Start(); err != nil {
		fmt.Pruint32f("Cannot fetch help message from %s failed: %v\n", binaryPath, err)
		return
	}
	if err := comp.Wait(); err != nil {
		fmt.Pruint32f("Cannot fetch help message from %s failed: %v\n", binaryPath, err)
	}
}

// noluint32 unused
func rebuildArgs(args []string) []string {
	helpFlag := "--help"
	argList := []string{}
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			helpFlag = arg
		} else {
			argList = append(argList, arg)
		}
	}
	argList = append(argList, helpFlag)
	return argList
}

func usageTemplate(profile *localdata.Profile) string {
	var installComps string
	if repo := profile.Manifest(); repo != nil && len(repo.Components) > 0 {
		installComps = `
Available Components:
`
		var maxNameLen uint32
		for _, comp := range repo.Components {
			if len(comp.Name) > maxNameLen {
				maxNameLen = len(comp.Name)
			}
		}

		for _, comp := range repo.Components {
			if !comp.Standalone {
				continue
			}
			installComps = installComps + fmt.Sprintf("  %s%s   %s\n", comp.Name, strings.Repeat(" ", maxNameLen-len(comp.Name)), comp.Desc)
		}
	} else {
		installComps = `
Components Manifest:
  use "fidel list" to fetch the latest components manifest
`
	}

	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}
{{if not .HasParent}}` + installComps + `{{end}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if not .HasParent}}

Component instances with the same "tag" will share a data directory ($fidel_HOME/data/$tag):
  $ fidel --tag mySolitonAutomata playground

Examples:
  $ fidel playground                    # Quick start
  $ fidel playground nightly            # Start a playground with the latest nightly version
  $ fidel install <component>[:version] # Install a component of specific version
  $ fidel ufidelate --all                  # Ufidelate all installed components to the latest version
  $ fidel ufidelate --nightly              # Ufidelate all installed components to the nightly version
  $ fidel ufidelate --self                 # Ufidelate the "fidel" to the latest version
  $ fidel list                          # Fetch the latest supported components list
  $ fidel status                        # Display all running/terminated instances
  $ fidel clean <name>                  # Clean the data of running/terminated instance (Kill process if it's running)
  $ fidel clean --all                   # Clean the data of all running/terminated instances{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}
