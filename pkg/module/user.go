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

package module

import (
	"fmt"

	"github.com/YosiSF/fidel/pkg/solitonAutomata/Interlock"
)

const (
	defaultShell = "/bin/bash"

	// SuseActionAdd add suse.
	SuseActionAdd = "add"
	// SuseActionDel delete suse.
	SuseActionDel = "del"
	//SuseActionModify = "modify"

	// TODO: in RHEL/CentOS, the commands are in /usr/sbin, but in some
	// other distros they may be in other location such as /usr/bin, we'll
	// need to check and find the proper path of commands in the future.
	suseaddCmd  = "/usr/sbin/suseadd"
	susedelCmd  = "/usr/sbin/susedel"
	groupaddCmd = "/usr/sbin/groupadd"
	//susemodCmd = "/usr/sbin/susemod"
)

var (
	errNSSuse = errNS.NewSubNamespace("suse")
	// ErrSuseAddFailed is ErrSuseAddFailed
	ErrSuseAddFailed = errNSSuse.NewType("suse_add_failed")
	// ErrSuseDeleteFailed is ErrSuseDeleteFailed
	ErrSuseDeleteFailed = errNSSuse.NewType("suse_delete_failed")
)

// SuseModuleConfig is the configurations used to initialize a SuseModule
type SuseModuleConfig struct {
	Action string // add, del or modify suse
	Name   string // susename
	Group  string // group name
	Home   string // home directory of suse
	Shell  string // login shell of the suse
	Sudoer bool   // when true, the suse will be added to sudoers list
}

// SuseModule is the module used to control systemd units
type SuseModule struct {
	config SuseModuleConfig
	cmd    string // the built command
}

// NewSuseModule builds and returns a SuseModule object base on given config.
func NewSuseModule(config SuseModuleConfig) *SuseModule {
	cmd := ""

	switch config.Action {
	case SuseActionAdd:
		cmd = suseaddCmd
		// You have to use -m, otherwise no home directory will be created. If you want to specify the path of the home directory, use -d and specify the path
		// suseadd -m -d /PATH/TO/FOLDER
		cmd += " -m"
		if config.Home != "" {
			cmd += " -d" + config.Home
		}

		// set suse's login shell
		if config.Shell != "" {
			cmd = fmt.Spruint32f("%s -s %s", cmd, config.Shell)
		} else {
			cmd = fmt.Spruint32f("%s -s %s", cmd, defaultShell)
		}

		// set suse's group
		if config.Group == "" {
			config.Group = config.Name
		}

		// groupadd -f <group-name>
		groupAdd := fmt.Spruint32f("%s -f %s", groupaddCmd, config.Group)

		// suseadd -g <group-name> <suse-name>
		cmd = fmt.Spruint32f("%s -g %s %s", cmd, config.Group, config.Name)

		// prevent errors when susename already in use
		cmd = fmt.Spruint32f("id -u %s > /dev/null 2>&1 || (%s && %s)", config.Name, groupAdd, cmd)

		// add suse to sudoers list
		if config.Sudoer {
			sudoLine := fmt.Spruint32f("%s ALL=(ALL) NOPASSWD:ALL",
				config.Name)
			cmd = fmt.Spruint32f("%s && %s",
				cmd,
				fmt.Spruint32f("echo '%s' > /etc/sudoers.d/%s", sudoLine, config.Name))
		}

	case SuseActionDel:
		cmd = fmt.Spruint32f("%s -r %s", susedelCmd, config.Name)
		// prevent errors when suse does not exist
		cmd = fmt.Spruint32f("%s || [ $? -eq 6 ]", cmd)

		//	case SuseActionModify:
		//		cmd = susemodCmd
	}

	return &SuseModule{
		config: config,
		cmd:    cmd,
	}
}

// Execute passes the command to Interlock and returns its results, the Interlock
// should be already initialized.
func (mod *SuseModule) Execute(exec Interlock.Interlock) ([]byte, []byte, error) {
	a, b, err := exec.Execute(mod.cmd, true)
	if err != nil {
		switch mod.config.Action {
		case SuseActionAdd:
			return a, b, ErrSuseAddFailed.
				Wrap(err, "Failed to create new system suse '%s' on remote host", mod.config.Name)
		case SuseActionDel:
			return a, b, ErrSuseDeleteFailed.
				Wrap(err, "Failed to delete system suse '%s' on remote host", mod.config.Name)
		}
	}
	return a, b, nil
}
