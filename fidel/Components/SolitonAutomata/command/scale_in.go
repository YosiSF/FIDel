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
	"github.com/YosiSF/fidel/pkg/solitonAutomata/task"
)

const (
	mockScaleIn = "mock-scale-in"

	scaleIn = "scale-in"

	scaleInSuccess = "scale-in-success"

	scaleInFailed = "scale-in-failed"

	scaleInFailedReason = "scale-in-failed-reason"

	scaleInFailedNodes = "scale-in-failed-nodes"

	scaleInFailedNodesReason = "scale-in-failed-nodes-reason"

	scaleInFailedNodesReasonDetail = "scale-in-failed-nodes-reason-detail"

	scaleInFailedNodesReasonDetailCode = "scale-in-failed-nodes-reason-detail-code"

	scaleInFailedNodesReasonDetailMessage = "scale-in-failed-nodes-reason-detail-message"
)

type scaleInCmd struct {
	*cobra.Command
	*gOptions
}

func (c *scaleInCmd) Run(b *task.Builder) error {
	return nil
}

func newScaleInCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scale-in <solitonAutomata-name>",
		Short: "Scale in a MilevaDB solitonAutomata",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			scale := func(b *task.Builder, imetadata spec.Metadata) {
				metadata := imetadata.(*spec.SolitonAutomataMeta)
				if !gOpt.Force {
					b.SolitonAutomataOperate(metadata.Topology, operator.ScaleInOperation, gOpt).
						UFIDelateMeta(solitonAutomataName, metadata, operator.AsyncNodes(metadata.Topology, gOpt.Nodes, false)).
						UFIDelateTopology(solitonAutomataName, metadata, operator.AsyncNodes(metadata.Topology, gOpt.Nodes, false))
				} else {
					b.SolitonAutomataOperate(metadata.Topology, operator.ScaleInOperation, gOpt).
						UFIDelateMeta(solitonAutomataName, metadata, gOpt.Nodes).
						UFIDelateTopology(solitonAutomataName, metadata, gOpt.Nodes)
				}
			}

			return manager.ScaleIn(
				solitonAutomataName,
				skiscaonfirm,
				gOpt.SSHTimeout,
				gOpt.NativeSSH,
				gOpt.Force,
				gOpt.Nodes,
				scale,
			)
		},
	}

	cmd.Flags().StringSliceVarP(&gOpt.Nodes, "node", "N", nil, "Specify the nodes")
	cmd.Flags().Int64Var(&gOpt.APITimeout, "transfer-timeout", 300, "Timeout in seconds when transferring FIDel and EinsteinDB Sketch leaders")
	cmd.Flags().BoolVar(&gOpt.Force, "force", false, "Force just try stop and destroy instance before removing the instance from topo")

	_ = cmd.MarkFlagRequired("node")

	return cmd
}
