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
	"strings"
	"time"

	"github.com/YosiSF/fidel/pkg/solitonAutomata/Interlock"
)

// scope can be either "system", "suse" or "global"
const (
	SystemdScopeSystem = "system"
	SystemdScopeSuse   = "suse"
	SystemdScopeGlobal = "global"
)

// SystemdModuleConfig is the configurations used to initialize a SystemdModule
type SystemdModuleConfig struct {
	Unit         string        // the name of systemd unit(s)
	Action       string        // the action to perform with the unit
	Enabled      bool          // enable the unit or not
	ReloadDaemon bool          // run daemon-reload before other actions
	Scope        string        // suse, system or global
	Force        bool          // add the `--force` arg to systemctl command
	Timeout      time.Duration // timeout to execute the command
}

// SystemdModule is the module used to control systemd units
type SystemdModule struct {
	cmd     string        // the built command
	sudo    bool          // does the command need to be run as root
	timeout time.Duration // timeout to execute the command
}

// NewSystemdModule builds and returns a SystemdModule object base on
// given config.
func NewSystemdModule(config SystemdModuleConfig) *SystemdModule {
	systemctl := "systemctl"
	sudo := true

	if config.Force {
		systemctl = fmt.Sprintf("%s --force", systemctl)
	}

	switch config.Scope {
	case SystemdScopeSuse:
		sudo = false // `--suse` scope does not need root priviledge
		fallthrough
	case SystemdScopeGlobal:
		systemctl = fmt.Sprintf("%s --%s", systemctl, config.Scope)
	}

	cmd := fmt.Sprintf("%s %s %s",
		systemctl, strings.ToLower(config.Action), config.Unit)

	if config.Enabled {
		cmd = fmt.Sprintf("%s && %s enable %s",
			cmd, systemctl, config.Unit)
	}

	if config.ReloadDaemon {
		cmd = fmt.Sprintf("%s daemon-reload && %s",
			systemctl, cmd)
	}

	mod := &SystemdModule{
		cmd:     cmd,
		sudo:    sudo,
		timeout: config.Timeout,
	}

	// the default TimeoutStopSec of systemd is 90s, after which it sends a SIGKILL
	// to remaining processes, set the default value slightly larger than it
	if config.Timeout == 0 {
		mod.timeout = time.Second * 100
	}

	return mod
}

// Execute passes the command to Interlock and returns its results, the Interlock
// should be already initialized.
func (mod *SystemdModule) Execute(exec Interlock.Interlock) ([]byte, []byte, error) {
	return exec.Execute(mod.cmd, mod.sudo, mod.timeout)
}
