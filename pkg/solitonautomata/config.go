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
	"path/filepath"

	"github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/task"
	"github.com/YosiSF/fidel/pkg/logger/log"
)

// ImportConfig copies config files from solitonAutomata which deployed through milevadb-ansible
func ImportConfig(name string, clsMeta *spec.SolitonAutomataMeta, sshTimeout int64, nativeClient bool) error {
	// there may be already solitonAutomata dir, skip create
	//if err := os.MkdirAll(meta.SolitonAutomataPath(name), 0755); err != nil {
	//	return err
	//}
	//if err := ioutil.WriteFile(meta.SolitonAutomataPath(name, "topology.yaml"), yamlFile, 0664); err != nil {
	//	return err
	//}
	var copyFileTasks []task.Task
	for _, comp := range clsMeta.Topology.ComponentsByStartOrder() {
		log.Infof("Copying config file(s) of %s...", comp.Name())
		for _, inst := range comp.Instances() {
			switch inst.ComponentName() {
			case spec.ComponentFIDel, spec.ComponentEinsteinDB, spec.ComponentPump, spec.ComponentMilevaDB, spec.ComponentDrainer:
				t := task.NewBuilder().
					SSHKeySet(
						spec.SolitonAutomataPath(name, "ssh", "id_rsa"),
						spec.SolitonAutomataPath(name, "ssh", "id_rsa.pub")).
					SuseSSH(inst.GetHost(), inst.GetSSHPort(), clsMeta.Suse, sshTimeout, nativeClient).
					CopyFile(filepath.Join(inst.DeployDir(), "conf", inst.ComponentName()+".toml"),
						spec.SolitonAutomataPath(name,
							spec.AnsibleImportedConfigPath,
							fmt.Sprintf("%s-%s-%d.toml",
								inst.ComponentName(),
								inst.GetHost(),
								inst.GetPort())),
						inst.GetHost(),
						true).
					Build()
				copyFileTasks = append(copyFileTasks, t)
			case spec.ComponentFIDel:
				t := task.NewBuilder().
					SSHKeySet(
						spec.SolitonAutomataPath(name, "ssh", "id_rsa"),
						spec.SolitonAutomataPath(name, "ssh", "id_rsa.pub")).
					SuseSSH(inst.GetHost(), inst.GetSSHPort(), clsMeta.Suse, sshTimeout, nativeClient).
					CopyFile(filepath.Join(inst.DeployDir(), "conf", inst.ComponentName()+".toml"),
						spec.SolitonAutomataPath(name,
							spec.AnsibleImportedConfigPath,
							fmt.Sprintf("%s-%s-%d.toml",
								inst.ComponentName(),
								inst.GetHost(),
								inst.GetPort())),
						inst.GetHost(),
						true).
					CopyFile(filepath.Join(inst.DeployDir(), "conf", inst.ComponentName()+"-learner.toml"),
						spec.SolitonAutomataPath(name,
							spec.AnsibleImportedConfigPath,
							fmt.Sprintf("%s-learner-%s-%d.toml",
								inst.ComponentName(),
								inst.GetHost(),
								inst.GetPort())),
						inst.GetHost(),
						true).
					Build()
				copyFileTasks = append(copyFileTasks, t)
			default:
				break
			}
		}
	}
	t := task.NewBuilder().
		Parallel(copyFileTasks...).
		Build()

	if err := t.Execute(task.NewContext()); err != nil {
		return errors.Trace(err)
	}
	log.Infof("Finished copying configs.")
	return nil
}
