// Copyright 2020 WHTCORPS INC, AUTHORS, ALL RIGHTS RESERVED.
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

package cmd




import (
	"fmt"
	_ `bytes`
	_ `io`
	_ `os`
	"strings"
	`os`
	`os/exec`
	`strconv`
	`time`


	cobra "github.com/spf13/cobra"

	//ceph, rook, and isovalent libraries here
	clusterd _ "github.com/rook/rook/pkg/clusterd"
attachment _	"github.com/rook/rook/pkg/daemon/ceph/agent/flexvolume/attachment"
  osd	"github.com/rook/rook/pkg/operator/ceph/cluster/osd"
	"github.com/rook/rook/pkg/operator/ceph/cluster/rbd"
	"github.com/rook/rook/pkg/operator/ceph/cluster/toolbox"

mgr	"github.com/rook/rook/pkg/operator/ceph/cluster/mgr"
	"github.com/rook/rook/pkg/operator/ceph/cluster/mon"
	_ "fmt"
	_ "io/ioutil"

	//ipfs
	_ "github.com/ipfs/go-ipfs-api"
	_ "github.com/ipfs/go-ipfs-cmds"
	_ "github.com/ipfs/go-ipfs-files"
	_ "github.com/ipfs/go-ipfs-util"

	//lotus
	_ "github.com/lotus/lotus/build/lotus"
	_ "github.com/lotus/lotus/build/lotus/common"
	_ "github.com/lotus/lotus/build/lotus/node"
	_ "github.com/lotus/lotus/build/lotus/node/config"
	_ "github.com/lotus/lotus/build/lotus/node/ipfs"

	//filecoin
	_ "github.com/filecoin-project/go-filecoin"
	_ "github.com/filecoin-project/go-filecoin/address"
	_ "github.com/filecoin-project/go-filecoin/types"
	_ "github.com/filecoin-project/go-filecoin/types/abi"
	_ "github.com/filecoin-project/go-filecoin/types/cid"

	//go-ipfs
	_ "os"
	_ "path/filepath"
)


var completionLongDesc = ` Rook installs rook on a host.  The rook command will install rook on a host and create a rook cluster.  The rook command will also create a rook cluster.
    
	The rook command has the following subcommands:
	
		rook install
		rook uninstall
		rook completion
		rook version
		rook config
		rook get
		rook set
		rook unset
		rook enable
		rook disable
		rook start
		rook stop
		rook upgrade
		rook upgrade-check
		rook upgrade-status
		rook upgrade-cancel
		rook upgrade-start
		rook upgrade-continue
		rook upgrade-pause
		rook upgrade-resume
		rook upgrade-rollback
		rook upgrade-rollback-status
		rook upgrade-rollback-cancel
		rook upgrade-rollback-start
		rook upgrade-rollback-continue
		rook upgrade-rollback-pause
		rook upgrade-rollback-resume
		rook upgrade-rollback-rollback
		rook upgrade-rollback-rollback-status
		rook upgrade-rollback-rollback-cancel
		rook upgrade-rollback-rollback-start
		rook upgrade-rollback-rollback-continue
		rook upgrade-rollback-rollback-pause
		rook upgrade-rollback-rollback-resume
		rook upgrade-rollback-rollback-rollback
		rook upgrade-rollback-rollback-rollback-status
		rook upgrade-rollback-rollback-rollback-cancel
		rook upgrade-rollback-rollback-rollback-start
		rook upgrade-rollback-rollback-rollback-continue
		rook upgrade-rollback-rollback-rollback-pause
		rook upgrade-rollback-rollback-rollback-resume
		rook upgrade-rollback-rollback-rollback-rollback
		rook upgrade-rollback-rollback-rollback-rollback-status
		rook upgrade-rollback-rollback-rollback-rollback-cancel
		rook upgrade-rollback-rollback-rollback-rollback-start`






type Cmd struct {
	*cobra.Command
}


func (c *Cmd) Run() error {
	return c.Command.Execute()
}



//////////////////////////////////////////////////////////////////////////////////////////////
// Rook Commands
//////////////////////////////////////////////////////////////////////////////////////////////


func NewCmdInstall() *Cmd {
	return &Cmd{
		Command: &cobra.Command{
			Use:   "install",
			Short: "Install rook on a host",
			Long:  completionLongDesc,
			Run: func(c *cobra.Command, args []string) {
				err := install()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			},
		},
	}
}


func NewCmdUninstall() *Cmd {
	return &Cmd{
		Command: &cobra.Command{
			Use:   "uninstall",
			Short: "Uninstall rook from a host",
			Long:  completionLongDesc,
			Run: func(c *cobra.Command, args []string) {
				err := uninstall()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			},
		},
	}
}



func NewCmdCompletion() *Cmd {
	return &Cmd{
		Command: &cobra.Command{
			Use:   "completion",
			Short: "Generate bash completion scripts",
			Long:  completionLongDesc,
			Run: func(c *cobra.Command, args []string) {
				err := completion()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			},
		},
	}
}


func NewCmdVersion() *Cmd {
	return &Cmd{
		Command: &cobra.Command{
			Use:   "version",
			Short: "Print the version of rook",
			Long:  completionLongDesc,
			Run: func(c *cobra.Command, args []string) {
				err := version()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			},
		},
	}
}

//ipfs setup
func NewCmdConfig() *Cmd {
	return &Cmd{
		Command: &cobra.Command{
			Use:   "config",
			Short: "Configure rook",
			Long:  completionLongDesc,
			Run: func(c *cobra.Command, args []string) {
				err := config()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			},
		},
	}
}


func NewCmdGet() *Cmd {
	return &Cmd{
		Command: &cobra.Command{
			Use:   "get",
			Short: "Get rook configuration",
			Long:  completionLongDesc,
			Run: func(c *cobra.Command, args []string) {
				err := get()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			},
		},
	}
}


func NewCmdSet() *Cmd {
	return &Cmd{
		Command: &cobra.Command{
			Use:   "set",
			Short: "Set rook configuration",
			Long:  completionLongDesc,
			Run: func(c *cobra.Command, args []string) {
				err := set()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			},
		},
	}
}


//ipfs setup with rook and ceph and isovalent
func NewCmdSetup() *Cmd {
	return &Cmd{
		Command: &cobra.Command{
			Use:   "setup",
			Short: "Setup rook",
			Long:  completionLongDesc,
			Run: func(c *cobra.Command, args []string) {
				err := setup()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			},
		},
	}
}

func install() error {
	// Function to provision a new cluster
	err := provision()
	if err != nil {
		return err
	}

	// Function to install the cluster
	err = installCluster()
	if err != nil {
		return err
	}

	return nil
}

func provision() interface{} {
	return nil
}

func help() error {
	return nil
}

func status() error {
	return nil
}

func playground() error {
	return nil

}


func installCluster() error {
	return nil
}

/*
completionLongDesc = `
Output shell completion code for the specified shell (bash or zsh).
The shell code must be evaluated to provide interactive
completion of pd-ctl commands.  This can be done by sourcing it from
the .bash_profile.
Note for zsh users: [1] zsh completions are only supported in versions of zsh >= 5.2


## Example



 */

const (
	completionLongDesc = ` Output shell completion code for the specified shell (bash or zsh).   The shell code must be evaluated to provide interactive completion of pd-ctl commands.  This can be done by sourcing it from the .bash_profile.  Note for zsh users: [1] zsh completions are only supported in versions of zsh >= 5.2

## Example

source <(pd-ctl completion bash)

## Notes

[1] zsh completions are only supported in versions of zsh >= 5.2
`
)


func completionZshSupported() bool {
	// Zsh completions are only supported in versions of zsh >= 5.2
	// The following is the earliest supported version of zsh
	const minZshVersion = "5.2"
	zshVersion, err := zsh.ZshVersion()
	if err != nil {
		return false
	}
	return zshVersion >= minZshVersion
}


func completionBashSupported() bool {
	// Bash completions are only supported in versions of bash >= 4.1
	// The following is the earliest supported version of bash
	const minBashVersion = "4.1"
	bashVersion, err := bash.BashVersion()
	if err != nil {
		return false
	}
	return bashVersion >= minBashVersion
}


func completionBashSupported() bool {
	// Bash completions are only supported in versions of bash >= 4.1
	// The following is the earliest supported version of bash
	const minBashVersion = "4.1"
	bashVersion, err := bash.BashVersion()
	if err != nil {
		return false
	}
	return bashVersion >= minBashVersion
}

func NewCompletionCmd(clusterdContext *clusterd.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Output shell completion code for the specified shell (bash or zsh)",
		Long:  completionLongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cmd.Usage()
				os.Exit(1)
			}
			shell := args[0]
			switch shell {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				if !completionZshSupported() {
					fmt.Printf("zsh completion is only supported in versions of zsh >= 5.2")
					os.Exit(1)
				}
				cmd.Root().GenZshCompletion(os.Stdout)
			default:
				fmt.Printf("Unsupported shell %q", shell)
				os.Exit(1)
			}
		},
	}
}

func completionZshSupported() bool {
	out, err := exec.Command("zsh", "-c", "echo $ZSH_VERSION").Output()
	if err != nil {
		return false
	}
	zshVersion := strings.Split(string(out), ".")
	if len(zshVersion) < 2 {
		return false
	}
	major, err := strconv.Atoi(zshVersion[0])
	if err != nil {
		return false
	}

	minor, err := strconv.Atoi(zshVersion[1])
	if err != nil {
		return false
	}
	return major > 5 || (major == 5 && minor >= 2)
}




func NewInstallCmd(clusterdContext *clusterd.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Install the cluster",
		Long:  "Install the cluster",
		Run: func(cmd *cobra.Command, args []string) {
			err := install()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}



func NewStatusCmd(clusterdContext *clusterd.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the status of the cluster",
		Long:  "Show the status of the cluster",
		Run: func(cmd *cobra.Command, args []string) {
			err := status()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}


func NewHelpCmd(clusterdContext *clusterd.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "Show help for pd-ctl",
		Long:  "Show help for pd-ctl",
		Run: func(cmd *cobra.Command, args []string) {
			err := help()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}



func NewPlaygroundCmd(clusterdContext *clusterd.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "playground",
		Short: "Playground",
		Long:  "Playground",
		Run: func(cmd *cobra.Command, args []string) {
			err := playground()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}


func NewRootCmd(clusterdContext *clusterd.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pd-ctl",
		Short: "Rook Ceph Cluster Controller",
		Long:  "Rook Ceph Cluster Controller",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(NewCompletionCmd(clusterdContext))
	cmd.AddCommand(NewInstallCmd(clusterdContext))
	cmd.AddCommand(NewStatusCmd(clusterdContext))
	cmd.AddCommand(NewHelpCmd(clusterdContext))
	cmd.AddCommand(NewPlaygroundCmd(clusterdContext))
	return cmd
}





func main() {
	cmd := NewRootCmd(clusterd.NewContext())
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}




