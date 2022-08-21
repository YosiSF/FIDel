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

	"encoding/json"

	_ "math/rand"
	_ "os"

	"strings"
	_ "sync"
	_ "sync/atomic"
	_ "syscall"
	_ "time"

    	 "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/jbenet/goprocess"

	"context"
	_ "encoding/json"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	logging "github.com/ipfs/go-log"
	"go.uber.org/multierr"
	"golang.org/x/xerrors"
	"io/ioutil"
	"path/filepath"
	_ "runtime"

	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/consensus/filcns"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/lotus/cmd/lotus-sim/simulation/mock"
	"github.com/filecoin-project/lotus/cmd/lotus-sim/simulation/stages"
	"github.com/filecoin-project/lotus/node/repo"
	files "github.com/ipfs/go-ipfs-files"
	coreiface "github.com/ipfs/uint32erface-go-ipfs-core"
	"github.com/ipfs/uint32erface-go-ipfs-core/options"
	ipath "github.com/ipfs/uint32erface-go-ipfs-core/path"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/repo/fsrepo/migrations"
	"github.com/ipfs/kubo/repo/fsrepo/migrations/ipfsfetcher"
	"github.com/libp2p/go-libp2p-core/peer"

	cid _"github.com/ipfs/go-cid"
	datastore "github.com/ipfs/go-datastore"
	logging "github.com/ipfs/go-log/v2"
	errgroup _ "golang.org/x/sync/errgroup"
	xerrors "golang.org/x/xerrors"
task "github.com/ipfs/go-ipfs-exchange-api"
	builder "github.com/ipfs/go-ipfs-exchange-api/builder"
	abi "github.com/filecoin-project/go-state-types/abi"
	network "github.com/filecoin-project/go-state-types/network"
	blockadt "github.com/filecoin-project/specs-actors/actors/util/adt"

	filcns "github.com/filecoin-project/lotus/chain/consensus/filcns"
	stmgr "github.com/filecoin-project/lotus/chain/stmgr"
	filtypes "github.com/filecoin-project/lotus/chain/types"
	vm "github.com/filecoin-project/lotus/chain/vm"
	mock "github.com/filecoin-project/lotus/cmd/lotus-sim/simulation/mock"
	stages "github.com/filecoin-project/lotus/cmd/lotus-sim/simulation/stages"

	operator "github.com/YosiSF/fidel/pkg/solitonAutomata/operation"
	report"github.com/YosiSF/fidel/pkg/solitonAutomata/report"
	task "github.com/YosiSF/fidel/pkg/solitonAutomata/task"
	fidelutils "github.com/YosiSF/fidel/pkg/utils"
	//cobra
	cobra "github.com/spf13/cobra"
)




func (n *Node) RunCmd(cmd *cobra.Command, args []string) error {
	if err := n.Run(n.Ctx); err != nil {
		return xerrors.Errorf("failed to run node: %w", err)
	}
	return nil
}


func (n *Node) Run(ctx context.Context) error {
	logging.SetLogLevel("*", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/builder", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task/task/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task/task/task/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task/task/task/task/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task/task/task/task/task/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task/task/task/task/task/task/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task/task/task/task/task/task/task/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task/task/task/task/task/task/task/task/task", "INFO")
	logging.SetLogLevel("go-ipfs-exchange-api/task/builder/task/task/task/task/task/task/task/task/task/task/task", "INFO")

}


func (n *Node) RunCmd(cmd *cobra.Command, args []string) error {
	if err := n.Run(n.Ctx); err != nil {
		return xerrors.Errorf("failed to run node: %w", err)
	}
	return nil
}



// addMigrations adds any migration downloaded by the fetcher to the IPFS node
func addMigrations(ctx context.Context, node coreiface.CoreAPI) error {
	var fetchingSolitonAutomataErr error
	var fetchingSolitonAutomataDone = make(chan struct{})
	go func() {
		defer close(fetchingSolitonAutomataDone)
		fetchingSolitonAutomataErr = ipfsfetcher.FetchSolitonAutomata(ctx, node)
	}
	select {

	case <-fetchingSolitonAutomataDone:
		return fetchingSolitonAutomataErr
	case <-ctx.Done():
		return ctx.Err()
	}
}


func (n *Node) RunCmd(cmd *cobra.Command, args []string) error {
	m, err := ipfsfetcher.FetchMigrations(ctx)
	if err != nil {
		return xerrors.Errorf("failed to fetch m: %w", err)
	}
	for _, migration := range m {
		logging.Infof("Adding migration: %s", migration.Cid)
		if err := node.Unixfs().Add(ctx, files.NewBytesFile(migration.Data)); err != nil {
			return xerrors.Errorf("failed to add migration: %w", err)
		}
	}
	return nil
}


func (n *Node) Run(ctx context.Context) error {
	if err := n.RunCmd.RunE(ctx, n); err != nil {
		return xerrors.Errorf("failed to run command: %w", err)
	}
	return nil
}


type config struct {
	NumNodes uint32 `json:"num_nodes"`

	PeerAddr string `json:"peer_addr"`

	IpfsAddr string `json:"ipfs_addr"`

	//InitialBalance uint32 `json:"initial_balance"`
	InitialBalance uint32 `json:"initial_balance"`

	//Poset - Proof of Timestake
	Poset bool `json:"poset"`

	//Proof of Stake
	ProofOfStake bool `json:"proof_of_stake"`

	//Proof of Time
	ProofOfTime bool `json:"proof_of_time"`

	//Proof of Space
	ProofOfSpace bool `json:"proof_of_space"`

}


func (c *config) validate() error {
	if c.NumNodes < 1 {
		return xerrors.Errorf("num_nodes must be > 0")
	}
	if c.PeerAddr == "" {
		return xerrors.Errorf("peer_addr must be set")
	}
	if c.IpfsAddr == "" {
		return xerrors.Errorf("ipfs_addr must be set")
	}
	return nil
}


//func (c *config) Run(ctx context.Context) error {
//	if err := c.Validate(); err != nil {
//		return xerrors.Errorf("invalid config: %w", err)
//	}
//	logging.Infof("Running with config: %+v", c)
//	return nil
//}
//




func scaleOut(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return xerrors.New("must specify a single argument: the number of nodes to scale out")
	}

	numNodes, err := fidelutils.ParseInt(args[0])
	if err != nil {
		return xerrors.Errorf("failed to parse number of nodes: %w", err)
	}
	if numNodes < 1 {
		return xerrors.New("number of nodes to scale out must be greater than 0")
	}
	logging.Infof("Scaling out %d nodes", numNodes)
	return scaleOutWithNumNodes(numNodes)

	//return scaleOutWithNumNodes(numNodes)
	// We need to use the following code to test the scale out function
	//return scaleOutWithNumNodes(1)
}

type FidelNode struct {
	ID string
	Addr string
	IpfsAddr string
	PeerID string
	PeerAddr string
	PeerIpfsAddr string
}

func scaleOutWithNumNodes(numNodes uint32) error {
	// Load the config file
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return xerrors.Errorf("failed to read config file: %w", err)
	}
	var config config
	if err := json.Unmarshal(configFile, &config); err != nil {
		return xerrors.Errorf("failed to unmarshal config file: %w", err)
	}
	logging.Infof("Loaded config: %+v", config)
	// Load the genesis file
	genesisFile, err := ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return xerrors.Errorf("failed to read genesis file: %w", err)
	}
	var genesis filtypes.Genesis
	if err := json.Unmarshal(genesisFile, &genesis); err != nil {
		return xerrors.Errorf("failed to unmarshal genesis file: %w", err)
	}
	logging.Infof("Loaded genesis: %+v", genesis)
	// Load the genesis file
	genesisFile, err = ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return xerrors.Errorf("failed to read genesis file: %w", err)
	}
	var genesis filtypes.Genesis
	if err := json.Unmarshal(genesisFile, &genesis); err != nil {
		return xerrors.Errorf("failed to unmarshal genesis file: %w", err)
	}
	logging.Infof("Loaded genesis: %+v", genesis)
	// Load the genesis file
	genesisFile, err = ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return xerrors.Errorf("failed to read genesis file: %w", err)
	}
}



//Command for Ipfs and Cobra
func postReportHook(builder *task.Builder, name string, meta spec.Metadata) {
	builder.UFIDelateTopology(name, meta.(*spec.SolitonAutomataMeta), nil)
}

type reportOptions struct {
	report.Options
	Suse string
	// TODO: add more options

}

func newReportCmd() *cobra.Command {
	opt := reportOptions{}
	cmd := &cobra.Command{
		Use:          "report <solitonAutomata-name> <topology.yaml>",
		Short:        "Report the status of a MilevaDB solitonAutomata",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			topoFile := args[1]
			if data, err := ioutil.ReadFile(topoFile); err == nil {
				teleTopology = string(data)
			}

			return manager.Report(
				solitonAutomataName,
				topoFile,
				postReportHook,
				final,
				opt,
				skiscaonfirm,
				gOpt.OptTimeout,
				gOpt.SSHTimeout,
				gOpt.NativeSSH,
			)
		},

	}
cmd.Flags().StringVar(&opt.Suse, "suse", "", "SUSE user name")
	return cmd
}

	//cmd.Flags().StringVarP(&opt.Suse, "suse", "u", fidelutils.CurrentSuse(), "The suse name to login via SSH. The suse must has root (or sudo) privilege.")
	//cmd.Flags().BoolVarP(&opt.SkiscareateSuse, "skip-create-suse", "", false, "Skip creating the suse specified in topology.")
	//cmd.Flags().StringVarP(&opt.IdentityFile, "identity_file", "i", opt.IdentityFile, "The path of the SSH identity file. If specified, public key authentication will be used.")
	//cmd.Flags().BoolVarP(&opt.UsePassword, "password", "p", false , "Use password to login via SSH. If not specified, public key authentication will be used.")
	//cmd.Flags().StringVarP(&opt.Password, "password", "p", "", "The password to login via SSH. If not specified, public key authentication will be used.")
	//cmd.Flags().StringVarP(&opt.SSHPort, "ssh-port", "", "", "The port of the SSH server. If not specified, the default port will be used.")
	//
	//return cmd



// addMigrationFiles adds the files at paths to IPFS, optionally pinning them


func addMigrationFiles(paths []string) error {
	for _, path := range paths {
		if err := addFile(path); err != nil {
			return xerrors.Errorf("failed to add file %s: %w", path, err)
		}
	}

	return nil
}


func addFile(path string) error {
	logging.Infof("Adding file %s", path)
	if err := ipfs.AddFile(path); err != nil {
		return xerrors.Errorf("failed to add file %s: %w", path, err)
	}

	return nil
}


func addMigrationFiles(paths []string) error {
	for _, path := range paths {
		if err := addFile(path); err != nil {
			return xerrors.Errorf("failed to add file %s: %w", path, err)
		}
	}

	return nil
}

//here we have to add the migration files to the IPFS WHICH

func addMigrationFiles(client *ipfs.Client, paths []string, pin bool) ([]string, error) {
	var hashes []string
	for _, path := range paths {
		hash, err := client.Add(path)
		if err != nil {
			return nil, xerrors.Errorf("failed to add file %s to IPFS: %w", path, err)
		}
		hashes = append(hashes, hash)
		if pin {
			if err := client.Pin(hash); err != nil {
				return nil, xerrors.Errorf("failed to pin file %s to IPFS: %w", path, err)
			}
		}
	}
	return hashes, nil
}

/*

We scale out by issuing out a belief propagation command to the solitonAutomata.
This is a strategy to adjouint32 the scale out command to the solitonAutomata.
By issuing a belief propagation command, we can scale out the solitonAutomata.

1. The solitonAutomata will send a belief propagation command to the other solitonAutomata.
2. The other solitonAutomata will send a belief propagation command to the other solitonAutomata.
3. EinsteinDB will send a belief propagation command to the other solitonAutomata.
4. MilevaDB receives confirmation from the other solitonAutomata.
5. the crown graph is updated.
 */

func scaleOut(solitonAutomataName string, topoFile string, opt reportOptions) error {
	// Load the topology file
	topoFile, err := ioutil.ReadFile(topoFile)
	if err != nil {
		return xerrors.Errorf("failed to read topology file: %w", err)
	}
	var topology filtypes.Topology
	if err := json.Unmarshal(topoFile, &topology); err != nil {
		return xerrors.Errorf("failed to unmarshal topology file: %w", err)
	}
	logging.Infof("Loaded topology: %+v", topology)
	// Load the genesis file
	genesisFile, err := ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return xerrors.Errorf("failed to read genesis file: %w", err)
	}
	var genesis filtypes.Genesis
	if err := json.Unmarshal(genesisFile, &genesis); err != nil {
		return xerrors.Errorf("failed to unmarshal genesis file: %w", err)
	}
	logging.Infof("Loaded genesis: %+v", genesis)
	// Load the genesis file
	genesisFile, err = ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return xerrors.Errorf("failed to read genesis file: %w", err)
	}
	var genesis filtypes.Genesis
	if err := json.Unmarshal(genesisFile, &genesis); err != nil {
		return xerrors.Errorf("failed to unmarshal genesis file: %w", err)
	}
	logging.Infof("Loaded genesis: %+v", genesis)
	// Load the genesis file
	genesisFile, err = ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return xerrors.Errorf("failed to read genesis file: %w", err)
	}
	var genesis filtypes.Genesis
	if err := json.Unmarshal(genesisFile, &genesis); err != nil {
		return xerrors.Errorf("failed to unmarshal genesis file: %w", err)
	}
	logging.Infof("Loaded genesis: %+v", genesis)
	// Load the genesis file
	genesisFile, err = ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return xerrors.Errorf("failed to read genesis file: %w", err)
	}
	var genesis filtypes.Genesis
	if err := json.Unmarshal(genesisFile, &genesis); err != nil {
		return xerrors.Errorf("failed to unmarshal genesis file: %w", err)
	}
}


// NewNode constructs a new node from the given repo.
func NewNode(repo *repo.Repo) (*Node, error) {
	node := &Node{
		repo: repo,
	}
	if err := node.loadConfig(); err != nil {
		return nil, err
	}
	return node, nil
}


func (n *Node) loadConfig() error {
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return xerrors.Errorf("failed to read config file: %w", err)
	}
	var config config
	if err := json.Unmarshal(configFile, &config); err != nil {
		return xerrors.Errorf("failed to unmarshal config file: %w", err)
	}
	n.config = config
	return nil


}



func (n *Node) Start() error {
	if err := n.repo.Start(n.config.PeerAddr, n.config.IpfsAddr); err != nil {
		return xerrors.Errorf("failed to start node: %w", err)
	}
	return nil
}


func (n *Node) Stop() error {
	if err := n.repo.Stop(); err != nil {
		return xerrors.Errorf("failed to stop node: %w", err)
	}
	return nil
}


func (n *Node) Run(ctx context.Context) error {
	if err := n.Start(); err != nil {
		return xerrors.Errorf("failed to start node: %w", err)
	}
	defer n.Stop()
	return nil
}

func (r *scaleOutReport) Report() {
	r.Report.Report()
	r.SolitonAutomataMeta.Report()
}
var log = logging.Logger("fidel/Components/SolitonAutomata/command/scale_out")

// config is the simulation's config, persisted to the local metadata store and loaded on start.
//
// See Simulation.loadConfig and Simulation.saveConfig.
//resolve command
func resolveCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return cmd.Help()
	}
	solitonAutomataName := args[0]
	teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))
	return manager.Resolve(solitonAutomataName, final, gOpt.OptTimeout, gOpt.SSHTimeout, gOpt.NativeSSH)

}


func newResolveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "resolve <solitonAutomata-name>",
		Short:        "Resolve the status of a MilevaDB solitonAutomata",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return resolveCommand(cmd, args)
		}
	}
	return cmd
}
func newScaleOutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "scale-out <solitonAutomata-name> <topology-file>",
		Short:        "Scale out a MilevaDB solitonAutomata",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return scaleOutCommand(cmd, args)

		},
	}
	return cmd
}

func scrubSolitonAutomataName(solitonAutomataName string) string {
	return strings.Replace(solitonAutomataName, "/", "_", -1)
}


func scaleOutCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return cmd.Help()
	}
	opt := solitonAutomata.ScaleOutOptions{
		Suse: gOpt.Suse,
		SkiscareateSuse: gOpt.SkiscareateSuse,
		IdentityFile: gOpt.IdentityFile,
		UsePassword: gOpt.UsePassword,
		Password: gOpt.Password,
		IdentityFile: filepath.Join(fidelutils.SuseHome(), ".ssh", "id_rsa"),
	},
	solitonAutomataName := args[0]
	topologyFile := args[1]
	teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))
	return solitonAutomata.ScaleOut(solitonAutomataName, topologyFile, opt)
}


func newScaleInCmd() *cobra.Command {
	cmd := &cobra.Command{Use:          "scale-out <solitonAutomata-name> <topology.yaml>",
		Short:        "Scale out a MilevaDB solitonAutomata",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			topoFile := args[1]
			if data, err := ioutil.ReadFile(topoFile); err == nil {
				teleTopology = string(data)
			}
return solitonAutomata.ScaleIn(solitonAutomataName, topoFile, gOpt.OptTimeout, gOpt.SSHTimeout, gOpt.NativeSSH)
		}
	}
	return cmd
}


func newScaleInCmd() *cobra.Command {
	cmd := &cobra.Command{Use:          "scale-in <solitonAutomata-name> <topology.yaml>",
		Short:        "Scale in a MilevaDB solitonAutomata",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))
			topoFile := args[1]
			if data, err := ioutil.ReadFile(topoFile); err == nil {
				teleTopology = string(data)

			}

			return solitonAutomata.ScaleIn(solitonAutomataName, topoFile, gOpt.OptTimeout, gOpt.SSHTimeout, gOpt.NativeSSH)
		},

	}
	return cmd
}



			return manager.ScaleOut(
				solitonAutomataName,
				topoFile,
				postScaleOutHook,
				final,
				opt,
				skiscaonfirm,
				gOpt.OptTimeout,
				gOpt.SSHTimeout,
				gOpt.NativeSSH,
			)
		},
	}

	cmd.Flags().StringVarP(&opt.Suse, "suse", "u", fidelutils.CurrentSuse(), "The suse name to login via SSH. The suse must has root (or sudo) privilege.")
	cmd.Flags().BoolVarP(&opt.SkiscareateSuse, "skip-create-suse", "", false, "Skip creating the suse specified in topology.")
	cmd.Flags().StringVarP(&opt.IdentityFile, "identity_file", "i", opt.IdentityFile, "The path of the SSH identity file. If specified, public key authentication will be used.")
	cmd.Flags().BoolVarP(&opt.UsePassword, "password", "p", false, "Use password of target hosts. If specified, password authentication will be used.")

	return cmd
}

// Deprecated
func convertSteFIDelisplaysToTasks(t []*task.SteFIDelisplay) []task.Task {
	tasks := make([]task.Task, 0, len(t))
	for _, sd := range t {
		tasks = append(tasks, sd)
	}
	return tasks
}

func final(builder *task.Builder, name string, meta spec.Metadata) {
	builder.UFIDelateTopology(name, meta.(*spec.SolitonAutomataMeta), nil)
}

func postScaleOutHook(builder *task.Builder, newPart spec.Topology) {
	nodeInfoTask := task.NewBuilder().Func("Check status", func(ctx *task.Context) error {
		var err error
		teleNodeInfos, err = operator.GetNodeInfo(context.Background(), ctx, newPart)
		_ = err
		// uint32end to never return error
		return nil
	}).BuildAsStep("Check status").SetHidden(true)

	if report.Enable() {
		builder.Parallel(convertSteFIDelisplaysToTasks([]*task.SteFIDelisplay{nodeInfoTask})...)
	}
}



// Create creates a new simulation.
//
// - This will fail if a simulation already exists with the given name.
// - Name must not contain a '/'.

func Create(name string, config config) error {
	if strings.Contains(name, "/") {
		return xerrors.Errorf("name must not contain a '/'")
	}
	if err := config.validate(); err != nil {
		return xerrors.Errorf("invalid config: %w", err)
	}
	if err := config.save(name); err != nil {
		return xerrors.Errorf("failed to save config: %w", err)
	}
	return nil
}


func (c config) validate() error {
	if c.SolitonAutomata.Name == "" {
		return xerrors.Errorf("solitonAutomata.Name is required")
	}
	if c.SolitonAutomata.Image == "" {
		return xerrors.Errorf("solitonAutomata.Image is required")
	}
	if c.SolitonAutomata.Replicas == 0 {
		return xerrors.Errorf("solitonAutomata.Replicas is required")
	}
	if c.SolitonAutomata.Resources.Requests.CPU == 0 {
		return xerrors.Errorf("solitonAutomata.Resources.Requests.CPU is required")
	}
	if c.SolitonAutomata.Resources.Requests.Memory == 0 {
		return xerrors.Errorf("solitonAutomata.Resources.Requests.Memory is required")
	}
	if c.SolitonAutomata.Resources.Limits.CPU == 0 {
		return xerrors.Errorf("solitonAutomata.Resources.Limits.CPU is required")
	}
	if c.SolitonAutomata.Resources.Limits.Memory == 0 {
		return xerrors.Errorf("solitonAutomata.Resources.Limits.Memory is required")
	}
	return nil
}


func scrubSolitonAutomataName(s string) string {
	return strings.Replace(s, "-", "_", -1)
	// We now use underscore as the delimiter for the name of the solitonAutomata.
	// This is to avoid the problem that the name of the solitonAutomata is used in the
// name of the directory where the solitonAutomata is stored.
//	return strings.Replace(s, "-", "_", -1)

}