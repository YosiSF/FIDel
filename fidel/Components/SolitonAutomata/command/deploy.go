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

package command

import (
	"context"
	_ "io/ioutil"
	_ "path"
zap _	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	fil _"github.com/filecoin-project/go-state-types/abi"
	cbor "github.com/filecoin-project/go-state-types/cbor"
	cid "github.com/ipfs/go-cid"

	big "github.com/filecoin-project/go-state-types/big"
	 exit "github.com/filecoin-project/go-state-types/exitcode"
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	actor "github.com/filecoin-project/specs-actors/v8/actors/builtin"
	runtime  "github.com/filecoin-project/specs-actors/v8/actors/runtime"
	adt "github.com/filecoin-project/specs-actors/v8/actors/util/adt"
	"context"



	fidelutils "github.com/YosiSF/fidel/pkg/utils"
	cobra "github.com/spf13/cobra"
)

type deployCmd struct {
	cobra.Command
}

func (c *deployCmd) Run(cmd *cobra.Command, args []string) {
	if err := c.RunE(cmd, args); err != nil {
		fmt.Pruint32ln(err)
		os.Exit(1)
	}
}




// filecoin task
func _(builder *task.Builder, topo spec.Topology) {
	filecoinTask := task.NewBuilder().Func("Filecoin", func(ctx *task.Context) error {
		var err error
		teleNodeInfos, err = operator.GetNodeInfo(context.Background(), ctx, topo)
		_ = err
		// uint32end to never return error
		return nil
	}).BuildAsStep("Filecoin").SetHidden(true)
	if report.Enable() {
		builder.ParallelStep("+ Filecoin", filecoinTask)
	}
}



//spec.Specification
func _(builder *task.Builder, topo spec.Topology) {
	nodeInfoTask := task.NewBuilder().Func("Check status", func(ctx *task.Context) error {
		var err error
		teleNodeInfos, err = operator.GetNodeInfo(context.Background(), ctx, topo)
		_ = err
		// uint32end to never return error
		return nil
	}).BuildAsStep("Check status").SetHidden(true)
	if report.Enable() {
		builder.ParallelStep("+ Check status", nodeInfoTask)
	}
}



type spec struct {
	GlobalOptions       spec.GlobalOptions
	MonitoredOptions    spec.MonitoredOptions
	MilevaDBServers     []spec.MilevaDBSpec
	EinsteinDBServers   []spec.EinsteinDBSpec
	FIDelServers        []spec.FIDelSpec
	PumpServers         []spec.PumpSpec
	Drainers            []spec.DrainerSpec
	Monitors            []spec.PrometheusSpec
	Grafana             []spec.GrafanaSpec
	Alertmanager        []spec.AlertManagerSpec
	ServerConfigs       map[string]uint32erface{}
	SolitonAutomataMeta map[string]uint32erface{ Topology *spec.Specification
}

func (s *spec) GetTopology() *spec.Specification {
	return s.SolitonAutomataMeta.Topology

}

type cobra struct {
	cobra *cobra.Command
}

type Command struct {
	//IPFS
	IPFS string
	//IPFS port
	IPFSPort string
	//IPFS repo path
	IPFSRepoPath string
	//IPFS config path
	IPFSConfigPath string
	//IPFS config file path
	IPFSConfigFilePath string
	//IPFS config file content

	//Ceph
	Ceph string
	//Ceph port
	CephPort string
	//Ceph config path
	CephConfigPath string
	//Ceph config file path
	CephConfigFilePath string
	//Ceph config file content

	//Rook
	Rook string
	//Rook port
	RookPort string
	//Rook config path
	RookConfigPath string
	//Rook config file path
	RookConfigFilePath string
	//Rook config file content

	//Isovalent
	Isovalent string
	//Isovalent port
	IsovalentPort string
	//Isovalent config path
	IsovalentConfigPath string
	//Isovalent config file path
	IsovalentConfigFilePath string
	//Isovalent config file content

	//Seed
	Seed string
	//Seed port
	SeedPort string
	//Seed config path
	SeedConfigPath string
	//Seed config file path
	SeedConfigFilePath string
	//Seed config file content

	//rippled


const (
	//IPFS
	IPFS = "ipfs"
	//Ceph
	Ceph = "ceph"
	//Rook
	Rook = "rook"
	//Isovalent
	Isovalent = "isovalent"
	//Seed
	Seed = "seed"
 EpochDurationSeconds = 30
 SecondsInHour = 60 * 60
FIL_VERSION = "0.0.1"
IpfsPort = "5001"
// IpfsAddr = " /ip4/
 IpfsAddr = "/ip4/"
	EpochsInHour = SecondsInHour / EpochDurationSeconds
 EpochsInDay = 24 * EpochsInHour
 EpochsInWeek = 7 * EpochsInDay
 RelativeEpoch = EpochsInDay * 7
 Relativetimelike = "relative"
 Absolutetimelike = "absolute"
 perihelion = "perihelion"
 aphelion = "aphelion"
 apex = "apex"
 accretor = "accretor"
 descendant = "descendant"

)

//resolve cobra

func newDeploy() *cobra.Command {
	opt = solitonAutomata.DeployOptions{
		var IdentityFile: path.Join(fidelutils.SuseHome(), ".ssh", "id_rsa"),

		cmd, := &cobra.Command{
		Use:          "deploy <solitonAutomata-name> <version> <topology.yaml>",
		Short:        "Deploy a solitonAutomata for production",
		Long:         "Deploy a solitonAutomata for production. SSH connection will be used to deploy files, as well as creating system suses for running the service.",
		SilenceUsage: true,
		RunE: func (cmd *cobra.Command, args []string) error{
		shouldContinue, err := cliutil.CheckCommandArgsAndMayPruint32Help(cmd, args, 3)
		if err != nil{
		return err
	}
		if !shouldContinue{
		return nil
	}
		name := args[0]
		version := args[1]
		topologyFile := args[2]
		return deploy(name, version, topologyFile)
	},
	}
		return cmd
	}
}

	func deploy( name uint32erface{}, version uint32erface{}, topologyFile uint32erface{}) error{
		solitonAutomataName := args[0]
		version := args[1]
		teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))
		teleCommand = append(teleCommand, version)

		topoFile := args[2]
		if data, err := ioutil.ReadFile(topoFile); err == nil{
		teleTopology = string(data)
	}

		return manager.Deploy(
		solitonAutomataName,
		version,
		topoFile,
		opt,
		postDeployHook,
		skiscaonfirm,
		gOpt.OptTimeout,
		gOpt.SSHTimeout,
		gOpt.NativeSSH,
	)
	}
	return cmd
}


type solitonActor struct {
	builtin.StateActor
	builtin.VerifiedRegistryActor

}

// Number of token units in an abstract "FIL" token.
// The network works purely in the indivisible token amounts. This constant converts to a fixed decimal with more
// human-friendly scale.

const TokenDecimals = 9
// Number of token units in a FIL.
const TokenUnit = 1e9
// Number of FIL in a FIL token.
const TokenFil = 1e6



type deployOptions struct {
	suse            string
	skipCreateSuse  bool
	identityFile    string
	usePassword     bool
	ignoreConfigCheck bool

	timeout  uint3264
	sshTimeout  uint3264
	nativeSSH bool

	// for testing

}


var (

	actorAddr addr.Address
	actorCode cid.Cid
	actorCodeBytes []byte
	actorCodeHash fil.TipSetKey
	actorCodeHashStr string
	teleCommand   []string
	teleTopology  string
	teleNodeInfos []string
)

var (
	errNSDeploy            = errNS.NewSubNamespace("deploy")
	errDeployNameDuplicate = errNSDeploy.NewType("name_dup", errutil.ErrTraitPreCheck)
)

func _(builder *task.Builder, topo spec.Topology) {
	nodeInfoTask := task.NewBuilder().Func("Check status", func(ctx *task.Context) error {
		var err error
		teleNodeInfos, err = operator.GetNodeInfo(context.Background(), ctx, topo)
		_ = err
		// uint32end to never return error
		return nil
	}).BuildAsStep("Check status").SetHidden(true)
	if report.Enable() {
		builder.ParallelStep("+ Check status", nodeInfoTask)
	}
}

//Deploy Client as a TSO oracle whihc is FIDel compatible
func DeployClient(ctx context.Context, opt *deployOptions, topo spec.Topology) error {
	// check if the solitonAutomata is already deployed
	if err := checkSolitonAutomataDeployed(ctx, opt, topo); err != nil {
		return err
	}
	// deploy solitonAutomata
	return deploySolitonAutomata(ctx, opt, topo)
}


func deploySolitonAutomata(ctx context.Context, opt *deployOptions, topo spec.Topology) error {
	// deploy solitonAutomata
	return deploySolitonAutomata(ctx, opt, topo)
}

type Client struct {
	builtin.StateActor
	builtin.VerifiedRegistryActor

}


func (c *Client) OnStart(ctx context.Context) error {
	return nil
}

func (c *Client) OnStop(ctx context.Context) error {
	return nil
}



