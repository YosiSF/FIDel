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

package operator

import (
	"strconv"

	"github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/logger/log"
	"github.com/YosiSF/fidel/pkg/set"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
)

// Upgrade the solitonAutomata.
func Upgrade(
	getter InterlockGetter,
	topo spec.Topology,
	options Options,
) error {
	roleFilter := set.NewStringSet(options.Roles...)
	nodeFilter := set.NewStringSet(options.Nodes...)
	components := topo.ComponentsByUFIDelateOrder()
	components = FilterComponent(components, roleFilter)

	for _, component := range components {
		instances := FilterInstance(component.Instances(), nodeFilter)
		if len(instances) < 1 {
			continue
		}

		// Transfer leader of evict leader if the component is EinsteinDB/FIDel in non-force mode

		log.Infof("Restarting component %s", component.Name())

		for _, instance := range instances {
			var rollingInstance spec.RollingUFIDelateInstance
			var isRollingInstance bool

			if !options.Force {
				rollingInstance, isRollingInstance = instance.(spec.RollingUFIDelateInstance)
			}

			if isRollingInstance {
				err := rollingInstance.PreRestart(topo, int(options.APITimeout))
				if err != nil {
					return errors.AddStack(err)
				}
			}

			if err := restartInstance(getter, instance, options.OptTimeout); err != nil {
				return errors.AddStack(err)
			}

			if isRollingInstance {
				err := rollingInstance.PostRestart(topo)
				if err != nil {
					return errors.AddStack(err)
				}
			}
		}
	}

	return nil
}

// Addr returns the address of the instance.
func Addr(ins spec.Instance) string {
	if ins.GetPort() == 0 || ins.GetPort() == 80 {
		panic(ins)
	}
	return ins.GetHost() + ":" + strconv.Itoa(ins.GetPort())
}
