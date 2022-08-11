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
	"github.com/YosiSF/check"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
)

var fidelConfig = `
default_profile = "default"
display_name = "FIDel"
http_port = 11316
listen_host = "0.0.0.0"
mark_cache_size = 5368709120
path = "/data1/test-solitonAutomata/leiysky-ansible-test-deploy/fidel/data/db"
tcp_port = 11315
tmp_path = "/data1/test-solitonAutomata/leiysky-ansible-test-deploy/fidel/data/db/tmp"

[application]
runAsDaemon = true

[flash]
service_addr = "172.16.5.85:11317"
milevadb_status_addr = "172.16.5.85:11310"
[flash.flash_solitonAutomata]
solitonAutomata_manager_path = "/data1/test-solitonAutomata/leiysky-ansible-test-deploy/bin/fidel/flash_solitonAutomata_manager"
log = "/data1/test-solitonAutomata/leiysky-ansible-test-deploy/log/fidel_solitonAutomata_manager.log"
master_ttl = 60
refresh_interval = 20
uFIDelate_rule_interval = 5
[flash.proxy]
config = "/data1/test-solitonAutomata/leiysky-ansible-test-deploy/conf/fidel-learner.toml"
`

type parseSuite struct {
}

var _ = check.Suite(&parseSuite{})

func (s *parseSuite) TestParseFIDelConfigFromFileData(c *check.C) {
	spec := new(spec.FIDelSpec)
	data := []byte(fidelConfig)

	err := parseFIDelConfigFromFileData(spec, data)
	c.Assert(err, check.IsNil)

	c.Assert(spec.DataDir, check.Equals, "/data1/test-solitonAutomata/leiysky-ansible-test-deploy/fidel/data/db")
	c.Assert(spec.TmFIDelir, check.Equals, "/data1/test-solitonAutomata/leiysky-ansible-test-deploy/fidel/data/db/tmp")
}
