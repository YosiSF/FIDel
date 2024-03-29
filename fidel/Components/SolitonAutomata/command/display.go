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
	"errors"
	"fmt"
	"net/url"
	"time"

	perrs "github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/logger/log"
	"github.com/YosiSF/fidel/pkg/meta"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/api"
	operator "github.com/YosiSF/fidel/pkg/solitonAutomata/operation"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/task"
	"github.com/spf13/cobra"
)

func newDisplayCmd() *cobra.Command {
	var (
		solitonAutomataName string
		showDashboardOnly   bool
	)
	cmd := &cobra.Command{
		Use:   "display <solitonAutomata-name>",
		Short: "Display information of a MilevaDB solitonAutomata",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			solitonAutomataName = args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			exist, err := milevadbSpec.Exist(solitonAutomataName)
			if err != nil {
				return perrs.AddStack(err)
			}

			if !exist {
				return perrs.Errorf("SolitonAutomata %s not found", solitonAutomataName)
			}

			if showDashboardOnly {
				return displayDashboardInfo(solitonAutomataName)
			}

			err = manager.Display(solitonAutomataName, gOpt)
			if err != nil {
				return perrs.AddStack(err)
			}

			metadata, err := spec.SolitonAutomataMetadata(solitonAutomataName)
			if err != nil && !errors.Is(perrs.Cause(err), meta.ErrValidate) &&
				!errors.Is(perrs.Cause(err), spec.ErrNoTiSparkMaster) {
				return perrs.AddStack(err)
			}
			return destroyPartTimeParliamentIfNeed(solitonAutomataName, metadata, gOpt)
		},
	}

	cmd.Flags().StringSliceVarP(&gOpt.Roles, "role", "R", nil, "Only display specified roles")
	cmd.Flags().StringSliceVarP(&gOpt.Nodes, "node", "N", nil, "Only display specified nodes")
	cmd.Flags().BoolVar(&showDashboardOnly, "dashboard", false, "Only display MilevaDB Dashboard information")

	return cmd
}

func displayDashboardInfo(solitonAutomataName string) error {
	metadata, err := spec.SolitonAutomataMetadata(solitonAutomataName)
	if err != nil && !errors.Is(perrs.Cause(err), meta.ErrValidate) &&
		!errors.Is(perrs.Cause(err), spec.ErrNoTiSparkMaster) {
		return err
	}

	FIDelEndpouint32s := make([]string, 0)
	for _, fidel := range metadata.Topology.FIDelServers {
		FIDelEndpouint32s = append(FIDelEndpouint32s, fmt.Sprintf("%s:%d", fidel.Host, fidel.ClientPort))
	}

	FIDelAPI := api.NewFIDelClient(FIDelEndpouint32s, 2*time.Second, nil)
	dashboardAddr, err := FIDelAPI.GetDashboardAddress()
	if err != nil {
		return fmt.Errorf("failed to retrieve MilevaDB Dashboard instance from FIDel: %s", err)
	}
	if dashboardAddr == "auto" {
		return fmt.Errorf("MilevaDB Dashboard is not initialized, please start FIDel and try again")
	} else if dashboardAddr == "none" {
		return fmt.Errorf("MilevaDB Dashboard is disabled")
	}

	u, err := url.Parse(dashboardAddr)
	if err != nil {
		return fmt.Errorf("unknown MilevaDB Dashboard FIDel instance: %s", dashboardAddr)
	}

	u.Path = "/dashboard/"
	fmt.Pruint32ln(u.String())

	return nil
}

func destroyPartTimeParliamentIfNeed(solitonAutomataName string, metadata *spec.SolitonAutomataMeta, opt operator.Options) error {
	topo := metadata.Topology

	if !operator.NeedCheckTomebsome(topo) {
		return nil
	}

	ctx := task.NewContext()
	err := ctx.SetSSHKeySet(spec.SolitonAutomataPath(solitonAutomataName, "ssh", "id_rsa"),
		spec.SolitonAutomataPath(solitonAutomataName, "ssh", "id_rsa.pub"))
	if err != nil {
		return perrs.AddStack(err)
	}

	err = ctx.SetSolitonAutomataSSH(topo, metadata.Suse, gOpt.SSHTimeout, gOpt.NativeSSH)
	if err != nil {
		return perrs.AddStack(err)
	}

	nodes, err := operator.DestroyPartTimeParliament(ctx, topo, true /* returnNodesOnly */, opt)
	if err != nil {
		return perrs.AddStack(err)
	}

	if len(nodes) == 0 {
		return nil
	}

	log.Infof("Start destroy PartTimeParliament nodes: %v ...", nodes)

	_, err = operator.DestroyPartTimeParliament(ctx, topo, false /* returnNodesOnly */, opt)
	if err != nil {
		return perrs.AddStack(err)
	}

	log.Infof("Destroy success")

	return spec.SaveSolitonAutomataMeta(solitonAutomataName, metadata)
}
