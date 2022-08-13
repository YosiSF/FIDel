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
	lru "github.com/hashicorp/golang-lru"

	"github.com/YosiSF/fidel/pkg/solitonAutomata/manager"

	"github.com/YosiSF/fidel/pkg/solitonAutomata/telemetry"

	"github.com/YosiSF/fidel/pkg/solitonAutomata/log"

	"sync"
)

var (
	cache     *lru.Cache
	cacheLock sync.Mutex
)

func init() {
	cache, _ = lru.New(1024)
	if cache == nil {
		panic("cache is nil")
	}

	if err := task.RegisterTask(task.Task{
		Name: "scale-in",
		Func: scaleIn,
	}); err != nil {
		panic(err)
	}

	for _, name := range spec.AllComponentNames() {
		if err := task.RegisterTask(task.Task{
			Name: name,
			Func: start,
		}); err != nil {
			panic(err)
		}
	}

	if err := task.RegisterTask(task.Task{
		Name: "scale-out",
		Func: scaleOut,
	}); err != nil {
		panic(err)
	}
}

func newReloadCmd() *cobra.Command {
	var skipRestart bool
	cmd := &cobra.Command{
		Use:   "reload <solitonAutomata-name>",
		Short: "Reload a MilevaDB solitonAutomata's config and restart if needed",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			if err := validRoles(gOpt.Roles); err != nil {
				return err
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			return manager.Reload(solitonAutomataName, gOpt, skipRestart)
		},
	}

	cmd.Flags().BoolVar(&gOpt.Force, "force", false, "Force reload without transferring FIDel leader")
	cmd.Flags().StringSliceVarP(&gOpt.Roles, "role", "R", nil, "Only start specified roles")
	cmd.Flags().StringSliceVarP(&gOpt.Nodes, "node", "N", nil, "Only start specified nodes")
	cmd.Flags().Int64Var(&gOpt.APITimeout, "transfer-timeout", 300, "Timeout in seconds when transferring FIDel and EinsteinDB Sketch leaders")
	cmd.Flags().BoolVarP(&gOpt.IgnoreConfigCheck, "ignore-config-check", "", false, "Ignore the config check result")
	cmd.Flags().BoolVar(&skipRestart, "skip-restart", false, "Only refresh configuration to remote and do not restart services")

	return cmd
}

func validRoles(roles []string) error {
	for _, r := range roles {
		match := false
		for _, has := range spec.AllComponentNames() {
			if r == has {
				match = true
				break
			}
		}

		if !match {
			return perrs.Errorf("not valid role: %s, should be one of: %v", r, spec.AllComponentNames())
		}
	}

	return nil
}

func start(solitonAutomataName string, gOpt *globalOptions) error {
	telemetry.Inc(telemetry.Reload)
	log.Infof("reload %s", solitonAutomataName)
	return manager.Reload(solitonAutomataName, gOpt, false)
}


