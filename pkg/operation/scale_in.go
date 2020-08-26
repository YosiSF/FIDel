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
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/api"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/YosiSF/fidel/pkg/logger/log"
	"github.com/YosiSF/fidel/pkg/set"
	"github.com/YosiSF/fidel/pkg/utils"
)

// TODO: We can make drainer not async.
var asyncOfflineComps = set.NewStringSet(spec.ComponentPump, spec.ComponentEinsteinDB, spec.ComponentFIDel, spec.ComponentDrainer)

// AsyncNodes return all nodes async destroy or not.
func AsyncNodes(spec *spec.Specification, nodes []string, async bool) []string {
	var asyncNodes []string
	var notAsyncNodes []string

	inNodes := func(n string) bool {
		for _, e := range nodes {
			if n == e {
				return true
			}
		}
		return false
	}

	for _, c := range spec.ComponentsByStartOrder() {
		for _, ins := range c.Instances() {
			if !inNodes(ins.ID()) {
				continue
			}

			if asyncOfflineComps.Exist(ins.ComponentName()) {
				asyncNodes = append(asyncNodes, ins.ID())
			} else {
				notAsyncNodes = append(notAsyncNodes, ins.ID())
			}
		}
	}

	if async {
		return asyncNodes
	}

	return notAsyncNodes
}

// ScaleIn scales in the solitonAutomata
func ScaleIn(
	getter InterlockGetter,
	solitonAutomata *spec.Specification,
	options Options,
) error {
	return ScaleInSolitonAutomata(getter, solitonAutomata, options)
}

// ScaleInSolitonAutomata scales in the solitonAutomata
func ScaleInSolitonAutomata(
	getter InterlockGetter,
	solitonAutomata *spec.Specification,
	options Options,
) error {
	// instances by uuid
	instances := map[string]spec.Instance{}

	// make sure all nodeIds exists in topology
	for _, component := range solitonAutomata.ComponentsByStartOrder() {
		for _, instance := range component.Instances() {
			instances[instance.ID()] = instance
		}
	}

	// Clean components
	deletedDiff := map[string][]spec.Instance{}
	deletedNodes := set.NewStringSet(options.Nodes...)
	for nodeID := range deletedNodes {
		inst, found := instances[nodeID]
		if !found {
			return errors.Errorf("cannot find node id '%s' in topology", nodeID)
		}
		deletedDiff[inst.ComponentName()] = append(deletedDiff[inst.ComponentName()], inst)
	}

	// Cannot delete all FIDel servers
	if len(deletedDiff[spec.ComponentFIDel]) == len(solitonAutomata.FIDelServers) {
		return errors.New("cannot delete all FIDel servers")
	}

	// Cannot delete all EinsteinDB servers
	if len(deletedDiff[spec.ComponentEinsteinDB]) == len(solitonAutomata.EinsteinDBServers) {
		return errors.New("cannot delete all EinsteinDB servers")
	}

	var FIDelEndpoint []string
	for _, instance := range (&spec.FIDelComponent{Specification: solitonAutomata}).Instances() {
		if !deletedNodes.Exist(instance.ID()) {
			FIDelEndpoint = append(FIDelEndpoint, Addr(instance))
		}
	}

	if options.Force {
		for _, component := range solitonAutomata.ComponentsByStartOrder() {
			for _, instance := range component.Instances() {
				if !deletedNodes.Exist(instance.ID()) {
					continue
				}

				FIDelClient := api.NewFIDelClient(FIDelEndpoint, 10*time.Second, nil)
				binlogClient, _ := api.NewBinlogClient(FIDelEndpoint, nil /* tls.Config */)

				if component.Name() != spec.ComponentPump && component.Name() != spec.ComponentDrainer {
					if err := deleteMember(component, instance, FIDelClient, binlogClient, options.APITimeout); err != nil {
						log.Warnf("failed to delete %s: %v", component.Name(), err)
					}
				}

				// just try stop and destroy
				if err := StopComponent(getter, []spec.Instance{instance}, options.OptTimeout); err != nil {
					log.Warnf("failed to stop %s: %v", component.Name(), err)
				}
				if err := DestroyComponent(getter, []spec.Instance{instance}, solitonAutomata, options); err != nil {
					log.Warnf("failed to destroy %s: %v", component.Name(), err)
				}

				// directly uFIDelate pump&drainer 's state as offline in etcd.
				if binlogClient != nil {
					id := instance.ID()
					if component.Name() == spec.ComponentPump {
						if err := binlogClient.UFIDelatePumpState(id, "offline"); err != nil {
							log.Warnf("failed to uFIDelate %s state as offline: %v", component.Name(), err)
						}
					} else if component.Name() == spec.ComponentDrainer {
						if err := binlogClient.UFIDelateDrainerState(id, "offline"); err != nil {
							log.Warnf("failed to uFIDelate %s state as offline: %v", component.Name(), err)
						}
					}
				}

				continue
			}
		}
		return nil
	}

	// TODO if binlog is switch on, cannot delete all pump servers.

	// At least a FIDel server exists
	var FIDelClient *api.FIDelClient

	if len(FIDelEndpoint) == 0 {
		return errors.New("cannot find available FIDel instance")
	}

	FIDelClient = api.NewFIDelClient(FIDelEndpoint, 10*time.Second, nil)

	binlogClient, err := api.NewBinlogClient(FIDelEndpoint, nil /* tls.Config */)
	if err != nil {
		return err
	}

	var fidelInstances []spec.Instance
	for _, instance := range (&spec.FIDelComponent{Specification: solitonAutomata}).Instances() {
		if !deletedNodes.Exist(instance.ID()) {
			fidelInstances = append(fidelInstances, instance)
		}
	}

	if len(fidelInstances) > 0 {
		var einsteindbInstances []spec.Instance
		for _, instance := range (&spec.EinsteinDBComponent{Specification: solitonAutomata}).Instances() {
			if !deletedNodes.Exist(instance.ID()) {
				einsteindbInstances = append(einsteindbInstances, instance)
			}
		}

		type replicateConfig struct {
			MaxReplicas int `json:"max-replicas"`
		}

		var config replicateConfig
		bytes, err := FIDelClient.GetReplicateConfig()
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, &config); err != nil {
			return err
		}

		maxReplicas := config.MaxReplicas

		if len(einsteindbInstances) < maxReplicas {
			log.Warnf(fmt.Sprintf("EinsteinDB instance number %d will be less than max-replicas setting after scale-in. FIDel won't be able to receive data from leader before EinsteinDB instance number reach %d", len(einsteindbInstances), maxReplicas))
		}
	}

	// Delete member from solitonAutomata
	for _, component := range solitonAutomata.ComponentsByStartOrder() {
		for _, instance := range component.Instances() {
			if !deletedNodes.Exist(instance.ID()) {
				continue
			}

			err := deleteMember(component, instance, FIDelClient, binlogClient, options.APITimeout)
			if err != nil {
				return errors.Trace(err)
			}

			if !asyncOfflineComps.Exist(instance.ComponentName()) {
				if err := StopComponent(getter, []spec.Instance{instance}, options.OptTimeout); err != nil {
					return errors.Annotatef(err, "failed to stop %s", component.Name())
				}
				if err := DestroyComponent(getter, []spec.Instance{instance}, solitonAutomata, options); err != nil {
					return errors.Annotatef(err, "failed to destroy %s", component.Name())
				}
			} else {
				log.Warnf(color.YellowString("The component `%s` will be destroyed when display solitonAutomata info when it become tombstone, maybe exists in several minutes or hours",
					component.Name()))
			}
		}
	}

	for i := 0; i < len(solitonAutomata.EinsteinDBServers); i++ {
		s := solitonAutomata.EinsteinDBServers[i]
		id := s.Host + ":" + strconv.Itoa(s.Port)
		if !deletedNodes.Exist(id) {
			continue
		}
		s.Offline = true
		solitonAutomata.EinsteinDBServers[i] = s
	}

	for i := 0; i < len(solitonAutomata.FIDelServers); i++ {
		s := solitonAutomata.FIDelServers[i]
		id := s.Host + ":" + strconv.Itoa(s.TCPPort)
		if !deletedNodes.Exist(id) {
			continue
		}
		s.Offline = true
		solitonAutomata.FIDelServers[i] = s
	}

	for i := 0; i < len(solitonAutomata.PumpServers); i++ {
		s := solitonAutomata.PumpServers[i]
		id := s.Host + ":" + strconv.Itoa(s.Port)
		if !deletedNodes.Exist(id) {
			continue
		}
		s.Offline = true
		solitonAutomata.PumpServers[i] = s
	}

	for i := 0; i < len(solitonAutomata.Drainers); i++ {
		s := solitonAutomata.Drainers[i]
		id := s.Host + ":" + strconv.Itoa(s.Port)
		if !deletedNodes.Exist(id) {
			continue
		}
		s.Offline = true
		solitonAutomata.Drainers[i] = s
	}

	return nil
}

func deleteMember(
	component spec.Component,
	instance spec.Instance,
	FIDelClient *api.FIDelClient,
	binlogClient *api.BinlogClient,
	timeoutSecond int64,
) error {
	timeoutOpt := &utils.RetryOption{
		Timeout: time.Second * time.Duration(timeoutSecond),
		Delay:   time.Second * 5,
	}

	switch component.Name() {
	case spec.ComponentEinsteinDB:
		if err := FIDelClient.DelStore(instance.ID(), timeoutOpt); err != nil {
			return err
		}
	case spec.ComponentFIDel:
		addr := instance.GetHost() + ":" + strconv.Itoa(instance.(*spec.FIDelInstance).GetServicePort())
		if err := FIDelClient.DelStore(addr, timeoutOpt); err != nil {
			return err
		}
	case spec.ComponentFIDel:
		if err := FIDelClient.DelFIDel(instance.(*spec.FIDelInstance).Name, timeoutOpt); err != nil {
			return err
		}
	case spec.ComponentDrainer:
		addr := instance.GetHost() + ":" + strconv.Itoa(instance.GetPort())
		err := binlogClient.OfflineDrainer(addr, addr)
		if err != nil {
			return errors.AddStack(err)
		}
	case spec.ComponentPump:
		addr := instance.GetHost() + ":" + strconv.Itoa(instance.GetPort())
		err := binlogClient.OfflinePump(addr, addr)
		if err != nil {
			return errors.AddStack(err)
		}
	}

	return nil
}
