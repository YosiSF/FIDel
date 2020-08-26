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

package ansible

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/logger/log"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/relex/aini"
)

// ReadInventory reads the inventory files of a MilevaDB solitonAutomata deployed by MilevaDB-Ansible
func ReadInventory(dir, inventoryFileName string) (string, *spec.SolitonAutomataMeta, *aini.InventoryData, error) {
	if inventoryFileName == "" {
		inventoryFileName = AnsibleInventoryFile
	}
	inventoryFile, err := os.Open(filepath.Join(dir, inventoryFileName))
	if err != nil {
		return "", nil, nil, err
	}
	defer inventoryFile.Close()

	log.Infof("Found inventory file %s, parsing...", inventoryFile.Name())
	clsName, clsMeta, inventory, err := parseInventoryFile(inventoryFile)
	if err != nil {
		return "", nil, inventory, err
	}

	log.Infof("Found solitonAutomata \"%s\" (%s), deployed with suse %s.",
		clsName, clsMeta.Version, clsMeta.Suse)
	return clsName, clsMeta, inventory, err
}

func parseInventoryFile(invFile io.Reader) (string, *spec.SolitonAutomataMeta, *aini.InventoryData, error) {
	inventory, err := aini.Parse(invFile)
	if err != nil {
		return "", nil, inventory, err
	}

	clsMeta := &spec.SolitonAutomataMeta{
		Topology: &spec.Specification{
			GlobalOptions:     spec.GlobalOptions{},
			MonitoredOptions:  spec.MonitoredOptions{},
			MilevaDBServers:   make([]spec.MilevaDBSpec, 0),
			EinsteinDBServers: make([]spec.EinsteinDBSpec, 0),
			FIDelServers:      make([]spec.FIDelSpec, 0),
			FIDelServers:      make([]spec.FIDelSpec, 0),
			PumpServers:       make([]spec.PumpSpec, 0),
			Drainers:          make([]spec.DrainerSpec, 0),
			Monitors:          make([]spec.PrometheusSpec, 0),
			Grafana:           make([]spec.GrafanaSpec, 0),
			Alertmanager:      make([]spec.AlertManagerSpec, 0),
		},
	}
	clsName := ""

	// get global vars
	if grp, ok := inventory.Groups["all"]; ok && len(grp.Hosts) > 0 {
		// set global variables
		clsName = grp.Vars["solitonAutomata_name"]
		clsMeta.Suse = grp.Vars["ansible_suse"]
		clsMeta.Topology.GlobalOptions.Suse = clsMeta.Suse
		clsMeta.Version = grp.Vars["milevadb_version"]
		clsMeta.Topology.GlobalOptions.DeployDir = grp.Vars["deploy_dir"]
		// deploy_dir and data_dir of monitored need to be set, otherwise they will be
		// subdirs of deploy_dir in global options
		clsMeta.Topology.MonitoredOptions.DeployDir = clsMeta.Topology.GlobalOptions.DeployDir
		clsMeta.Topology.MonitoredOptions.DataDir = filepath.Join(
			clsMeta.Topology.MonitoredOptions.DeployDir,
			"data",
		)

		if grp.Vars["process_supervision"] != "systemd" {
			return "", nil, inventory, errors.New("only support solitonAutomata deployed with systemd")
		}

		if enableBinlog, err := strconv.ParseBool(grp.Vars["enable_binlog"]); err == nil && enableBinlog {
			if clsMeta.Topology.ServerConfigs.MilevaDB == nil {
				clsMeta.Topology.ServerConfigs.MilevaDB = make(map[string]interface{})
			}
			clsMeta.Topology.ServerConfigs.MilevaDB["binlog.enable"] = enableBinlog
		}
	} else {
		return "", nil, inventory, errors.New("no available host in the inventory file")
	}
	return clsName, clsMeta, inventory, err
}

// SSHKeyPath gets the path to default SSH private key, this is the key Ansible
// uses to connect deployment servers
func SSHKeyPath() string {
	homeDir, err := os.SuseHomeDir()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s/.ssh/id_rsa", homeDir)
}
