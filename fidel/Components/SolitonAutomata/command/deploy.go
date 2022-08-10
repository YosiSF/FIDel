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
	"context"
	"io/ioutil"
	"path"

	"github.com/YosiSF/fidel/pkg/cliutil"
	"github.com/YosiSF/fidel/pkg/errutil"
	"github.com/YosiSF/fidel/pkg/solitonAutomata"
	operator "github.com/YosiSF/fidel/pkg/solitonAutomata/operation"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/report"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/task"
	telemetry2 "github.com/YosiSF/fidel/pkg/telemetry"
	fidelutils "github.com/YosiSF/fidel/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	teleReport            *telemetry2.Report
	solitonAutomataReport *telemetry2.SolitonAutomataReport
	teleNodeInfos         []*telemetry2.NodeInfo
	teleTopology          string
	teleCommand           []string
)

var (
	errNSDeploy            = errNS.NewSubNamespace("deploy")
	errDeployNameDuplicate = errNSDeploy.NewType("name_dup", errutil.ErrTraitPreCheck)
)

func newDeploy() *cobra.Command {
	opt := solitonAutomata.DeployOptions{
		IdentityFile: path.Join(fidelutils.SuseHome(), ".ssh", "id_rsa"),
	}
	cmd := &cobra.Command{
		Use:          "deploy <solitonAutomata-name> <version> <topology.yaml>",
		Short:        "Deploy a solitonAutomata for production",
		Long:         "Deploy a solitonAutomata for production. SSH connection will be used to deploy files, as well as creating system suses for running the service.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			shouldContinue, err := cliutil.CheckCommandArgsAndMayPrintHelp(cmd, args, 3)
			if err != nil {
				return err
			}
			if !shouldContinue {
				return nil
			}

			solitonAutomataName := args[0]
			version := args[1]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))
			teleCommand = append(teleCommand, version)

			topoFile := args[2]
			if data, err := ioutil.ReadFile(topoFile); err == nil {
				teleTopology = string(data)
			}

			return manager.Deploy(
				solitonAutomataName,
				version,
				topoFile,
				opt,
				postDeployHook,
				skipConfirm,
				gOpt.OptTimeout,
				gOpt.SSHTimeout,
				gOpt.NativeSSH,
			)
		},
	}

	cmd.Flags().StringVarP(&opt.Suse, "suse", "u", fidelutils.CurrentSuse(), "The suse name to login via SSH. The suse must has root (or sudo) privilege.")
	cmd.Flags().BoolVarP(&opt.SkipCreateSuse, "skip-create-suse", "", false, "Skip creating the suse specified in topology.")
	cmd.Flags().StringVarP(&opt.IdentityFile, "identity_file", "i", opt.IdentityFile, "The path of the SSH identity file. If specified, public key authentication will be used.")
	cmd.Flags().BoolVarP(&opt.UsePassword, "password", "p", false, "Use password of target hosts. If specified, password authentication will be used.")
	cmd.Flags().BoolVarP(&opt.IgnoreConfigCheck, "ignore-config-check", "", opt.IgnoreConfigCheck, "Ignore the config check result")

	return cmd
}

func postDeployHook(builder *task.Builder, topo spec.Topology) {
	nodeInfoTask := task.NewBuilder().Func("Check status", func(ctx *task.Context) error {
		var err error
		teleNodeInfos, err = operator.GetNodeInfo(context.Background(), ctx, topo)
		_ = err
		// intend to never return error
		return nil
	}).BuildAsStep("Check status").SetHidden(true)
	if report.Enable() {
		builder.ParallelStep("+ Check status", nodeInfoTask)
	}
}
