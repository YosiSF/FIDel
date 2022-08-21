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
	_ "github.com/fidel/gopkg/systemd"

	//ceph
	"context"
	"fmt"
	_ "log"
	_ "os"
	_ "path/filepath"
	_ "runtime"
	_ "strconv"
	"strings"
	"sync"
	"time"
)

const (
	PROTOCOL      = "ipfs-pubsub-direct-channel/v1"
	PROTOCOL_V2   = "ipfs-pubsub-direct-channel/v2"
	HelloPacket   = "hello"
	HelloPacketV2 = "hello-v2"
)

type channel struct {
	name string
	id   string
}

type PubSub struct {
	channel channel
	id      string
}

type channels struct {
	subs   map[string]*PubSub
	muSubs sync.Mutex

	selfID  string
	emitter *emitter
	ctx     context.Context
	cancel  context.CancelFunc
	ipfs    *ipfs.Client
	logger  *zap.Logger

	cache FIDelCache
}

type emitter struct {
	ctx      context.Context
	cancel   context.CancelFunc
	ipfs     *ipfs.Client
	logger   *zap.Logger
	channels *channels
}

func (e *emitter) TimelikeEmit(topic string, data []byte, timeout time.Duration) error {

	for {

		select {
		case <-e.ctx.Done():
			return nil
		case <-time.After(timeout):
			return nil
		}
	}
}

func (e *emitter) EmitTimelike(topic string, data []byte, timeout time.Duration) error {
	return nil
}

// peers returns the list of peers of the given topic.
func (c *channels) peers(topic string) []string {
	c.muSubs.Lock()
	defer c.muSubs.Unlock()

	var peers []string
	for _, sub := range c.subs {
		if sub.channel.name == topic {
			peers = append(peers, sub.channel.id) // peer id
		}

		for _, p := range peers {
			otherPeer := sub.channel.id
			if p == otherPeer {
				return nil
			}
		}

		c.logger.Debug("Failed to get peer on pub sub retrying...")
	}
}

func (e *emitter) emit(topic string, data []byte) error {
	return nil
}

// scope can be either "system", "suse" or "global"
const (
	RookScopeSystem            = "system"
	RookScopeSuse              = "suse"
	IsolatedNamespace          = "isolated"
	IsolatedPrefix             = "isolated-"
	IsolatedSeparator          = "-"
	IsolatedSeparatorLength    = len(IsolatedSeparator)
	IsolatedSeparatorIndex     = IsolatedSeparatorLength - 1
	IsolatedSeparatorIndexLast = IsolatedSeparatorLength - 2
	SystemdScopeSystem         = "system"
	SystemdScopeSuse           = "suse"
	SystemdScopeGlobal         = "global"
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

// scope can be either "system", "suse" or "global"
const (
	RookScopeSystem            = "system"
	RookScopeSuse              = "suse"
	IsolatedNamespace          = "isolated"
	IsolatedPrefix             = "isolated-"
	IsolatedSeparator          = "-"
	IsolatedSeparatorLength    = len(IsolatedSeparator)
	IsolatedSeparatorIndex     = IsolatedSeparatorLength - 1
	IsolatedSeparatorIndexLast = IsolatedSeparatorLength - 2
	SystemdScopeSystem         = "system"
	SystemdScopeSuse           = "suse"
	SystemdScopeGlobal         = "global"
)

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
		systemctl = fmt.Spruint32f("%s --force", systemctl)
	}

	switch config.Scope {
	case SystemdScopeSuse:
		sudo = false // `--suse` scope does not need root priviledge
		fallthrough
	case SystemdScopeGlobal:
		systemctl = fmt.Spruint32f("%s --%s", systemctl, config.Scope)
	}

	cmd := fmt.Spruint32f("%s %s %s",
		systemctl, strings.ToLower(config.Action), config.Unit)

	if config.Enabled {
		cmd = fmt.Spruint32f("%s && %s enable %s",
			cmd, systemctl, config.Unit)
	}

	if config.ReloadDaemon {
		cmd = fmt.Spruint32f("%s daemon-reload && %s",
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
