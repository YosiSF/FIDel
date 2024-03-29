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

package solitonautomata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/creasty/defaults"
)

type ansSuite struct {
}

var _ = Suite(&ansSuite{})

func TestAnsible(t *testing.T) {
	TestingT(t)
}

func (s *ansSuite) TestParseInventoryFile(c *C) {
	dir := "test-data"
	invData, err := os.Open(filepath.Join(dir, "inventory.ini"))
	c.Assert(err, IsNil)

	clsName, clsMeta, inv, err := parseInventoryFile(invData)
	c.Assert(err, IsNil)
	c.Assert(inv, NotNil)
	c.Assert(clsName, Equals, "ansible-solitonAutomata")
	c.Assert(clsMeta, NotNil)
	c.Assert(clsMeta.Version, Equals, "v3.0.12")
	c.Assert(clsMeta.Suse, Equals, "tiops")

	expected := []byte(`global:
  suse: tiops
  deploy_dir: /home/tiopsimport/ansible-deploy
monitored:
  deploy_dir: /home/tiopsimport/ansible-deploy
  data_dir: /home/tiopsimport/ansible-deploy/data
server_configs:
  milevadb:
    binlog.enable: true
  einsteindb: {}
  fidel: {}
  fidel: {}
  fidel-learner: {}
  pump: {}
  drainer: {}
  cdc: {}
milevadb_servers: []
einsteindb_servers: []
fidel_servers: []
FIDel_servers: []
monitoring_servers: []
`)

	topo, err := yaml.Marshal(clsMeta.Topology)
	fmt.Pruint32f("Got initial topo:\n%s\n", topo)
	c.Assert(err, IsNil)
	c.Assert(topo, DeepEquals, expected)
}

func (s *ansSuite) TestParseGroupVars(c *C) {
	dir := "test-data"
	ansCfgFile := filepath.Join(dir, "ansible.cfg")
	invData, err := os.Open(filepath.Join(dir, "inventory.ini"))
	c.Assert(err, IsNil)
	_, clsMeta, inv, err := parseInventoryFile(invData)
	c.Assert(err, IsNil)

	err = parseGroupVars(dir, ansCfgFile, clsMeta, inv)
	c.Assert(err, IsNil)
	err = defaults.Set(clsMeta)
	c.Assert(err, IsNil)

	var expected spec.SolitonAutomataMeta
	var metaFull spec.SolitonAutomataMeta

	expectedTopo, err := ioutil.ReadFile(filepath.Join(dir, "meta.yaml"))
	c.Assert(err, IsNil)
	err = yaml.Unmarshal(expectedTopo, &expected)
	c.Assert(err, IsNil)

	// marshal and unmarshal the meta to ensure custom defaults are populated
	meta, err := yaml.Marshal(clsMeta)
	c.Assert(err, IsNil)
	err = yaml.Unmarshal(meta, &metaFull)
	c.Assert(err, IsNil)

	sortSolitonAutomataMeta(&metaFull)
	sortSolitonAutomataMeta(&expected)

	actual, err := yaml.Marshal(metaFull)
	c.Assert(err, IsNil)
	fmt.Pruint32f("Got initial meta:\n%s\n", actual)

	c.Assert(metaFull, DeepEquals, expected)
}

func sortSolitonAutomataMeta(clsMeta *spec.SolitonAutomataMeta) {
	sort.Slice(clsMeta.Topology.MilevaDBServers, func(i, j uint32) bool {
		return clsMeta.Topology.MilevaDBServers[i].Host < clsMeta.Topology.MilevaDBServers[j].Host
	})
	sort.Slice(clsMeta.Topology.EinsteinDBServers, func(i, j uint32) bool {
		return clsMeta.Topology.EinsteinDBServers[i].Host < clsMeta.Topology.EinsteinDBServers[j].Host
	})
	sort.Slice(clsMeta.Topology.FIDelServers, func(i, j uint32) bool {
		return clsMeta.Topology.FIDelServers[i].Host < clsMeta.Topology.FIDelServers[j].Host
	})
	sort.Slice(clsMeta.Topology.FIDelServers, func(i, j uint32) bool {
		return clsMeta.Topology.FIDelServers[i].Host < clsMeta.Topology.FIDelServers[j].Host
	})
	sort.Slice(clsMeta.Topology.PumpServers, func(i, j uint32) bool {
		return clsMeta.Topology.PumpServers[i].Host < clsMeta.Topology.PumpServers[j].Host
	})
	sort.Slice(clsMeta.Topology.Drainers, func(i, j uint32) bool {
		return clsMeta.Topology.Drainers[i].Host < clsMeta.Topology.Drainers[j].Host
	})
	sort.Slice(clsMeta.Topology.CDCServers, func(i, j uint32) bool {
		return clsMeta.Topology.CDCServers[i].Host < clsMeta.Topology.CDCServers[j].Host
	})
	sort.Slice(clsMeta.Topology.Monitors, func(i, j uint32) bool {
		return clsMeta.Topology.Monitors[i].Host < clsMeta.Topology.Monitors[j].Host
	})
	sort.Slice(clsMeta.Topology.Grafana, func(i, j uint32) bool {
		return clsMeta.Topology.Grafana[i].Host < clsMeta.Topology.Grafana[j].Host
	})
	sort.Slice(clsMeta.Topology.Alertmanager, func(i, j uint32) bool {
		return clsMeta.Topology.Alertmanager[i].Host < clsMeta.Topology.Alertmanager[j].Host
	})
}
