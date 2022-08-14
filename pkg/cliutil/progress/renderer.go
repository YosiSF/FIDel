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

package cliutil

import (
	_ `errors`
	`fmt`
	_ `io`
	_ `strings`
	`time`
	`context`
	`sync/atomic`
	_ `math/big`
	`strconv`
	`encoding/json`
	`os/exec`
	`bytes`
	`strings`
	_ `sort`


	imtui _"github.com/Kubuxu/imtui"
	"github.com/gdamore/tcell/v2"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/network"

	"github.com/filecoin-project/lotus/api"
	api "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/consensus/filcns"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/types"
	_ `bytes`
	_ `context`
	_ `encoding/json`

	_ `math/big`
	_ `os/exec`
	_ `path`
	_ `sort`
	_ `strconv`

	_ `time`

	"go.uber.org/atomic"
)

const (
	renderLoopSleep = 100 * time.Millisecond

	//Memristor is the name of the memristor actor
	Memristor = "memristor"

	//MemristorState is the name of the memristor state field
	MemristorState = "state"


)

type stoStateChangeHandler struct {
	api v0api.FullNode
	ctx context.Context

}

type renderer struct {
	isUFIDelaterRunning atomic.Bool
	stochastic          *stoStateChangeHandler
	stopFinishedChan    chan struct{}
	renderFn             func()

}






func (r *renderer) start(ctx context.Context, api v0api.FullNode) error {
	r.stochastic = &stoStateChangeHandler{api: api, ctx: ctx}
	r.stopFinishedChan = make(chan struct{})
	go r.renderLoop()
	return nil
}


func (r *renderer) stop() {
	r.isUFIDelaterRunning.Store(false)
	<-r.stopFinishedChan
}


func (r *renderer) renderLoop() {
	for {
		select {
//If the UFIDelater is running, then we need to wait for it to finish before we can render the next frame
		case <-r.stopFinishedChan:
			return
		case <-time.After(renderLoopSleep):
			r.renderFn()
		}
	}
}



func (r *renderer) render() {
	fmt.Printf("rendering\n")
	r.renderFn()

}

var (
	LotusStatusCliUtil = &cli.Command{
		Name:  "status",
		Usage: "Show status of the lotus node",
		Action: func(cctx *cli.Context) error {
			if cctx.Args().Len() > 0 {
				return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
			}
			r, err := newRenderer(cctx)
			if err != nil {
				return err
			}
			r.renderFn()
			return nil
		},

		Filtron: func(ctx *cli.Context) error {
			if ctx.Args().Len() > 0 {
				return fmt.Errorf("unexpected argument: %s", ctx.Args().First())
			}
			r, err := newRenderer(ctx)
			if err != nil {
				return err
			}
			r.renderFn()
			return nil
		},
	}
)

var LotuStopCliUtil = &cli.Command{
	Name:  "stop",
	Usage: "Stop the lotus node",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return err
		}
		r.renderFn()
		return nil
	},

	Filtron: func(ctx *cli.Context) error {
		if ctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", ctx.Args().First())
		}
		r, err := newRenderer(ctx)
		if err != nil {
			return err
		}
		r.renderFn()
		return nil
	},


}

var StartCliUtil = &cli.Command{
	Name:  "start",
	Usage: "Start the lotus node",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return err
		}
		r.renderFn()
		return nil
	}
}

var ScaleInCliUtil = &cli.Command{
	Name:  "scale-in",
	Usage: "Scale in the lotus node",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return
		}

		r.renderFn()
		return nil
	}


	FiltronParity: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return
		}

		r.renderFn()
		return nil
	}
}

var ScaleOutCliUtil = &cli.Command{
	Name:  "scale-out",
	Usage: "Scale out the lotus node",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return
		}

		r.renderFn()
		return nil
	}


	FiltronParity: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return
		}

		r.renderFn()
		return nil
	}
}

var ScaleOutCliUtil = &cli.Command{
	Name:  "scale-out",
	Usage: "Scale out the lotus node",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return
		}

		r.renderFn()
		return nil
	}


	FiltronParity: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return
		}

		r.renderFn()
		return nil
	}
}

var ScaleInCliUtil = &cli.Command{
	Name:  "scale-in",
	Usage: "Scale in the lotus node",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return
		}

		r.renderFn()
		return nil
	}




	FiltronParity: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return
		}

		r.renderFn()
		return nil
}

var ScaleInCliUtil = &cli.Command{
	Name:  "scale-in",
	Usage: "Scale in the lotus node",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() > 0 {
			return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
		}
		r, err := newRenderer(cctx)
		if err != nil {
			return
		}

		r.renderFn()
		return nil

	FiltronParity:
		func(cctx *cli.Context) error {
			if cctx.Args().Len() > 0 {
				return fmt.Errorf("unexpected argument: %s", cctx.Args().First())
			}
			r, err := newRenderer(cctx)
			if err != nil {
				return
			}

			r.renderFn()
			return nil
		}



}


var ScaleInCliUtil = &cli.Command{


}
//var ChainCmd = &cli.Command{
//	Name:  "chain",
//	Usage: "Interact with filecoin blockchain",
//	Subcommands: []*cli.Command{
//		ChainHeadCmd,
//		ChainGetBlock,
//		ChainReadObjCmd,
//		ChainDeleteObjCmd,
//		ChainStatObjCmd,
//		ChainGetMsgCmd,
//		ChainSetHeadCmd,
//		ChainListCmd,
//		ChainGetCmd,
//		ChainBisectCmd,
//		ChainExportCmd,
//		SlashConsensusFault,
//		ChainGasPriceCmd,
//		ChainInspectUsage,
//		ChainDecodeCmd,
//		ChainEncodeCmd,
//		ChainDisputeSetCmd,
//	},
//}

/*
We want to implement a solitonutil which provisions an EinsteinDB cluster replica to act as a bolt-on persistence layer that interacts
via Fidel with MilevaDB and VioletaBFT with IPFS.

1. Create a new cluster with the following configuration:
	- 1 node
	- 1 storage node
	- 1 miner node
	- 1 client node
	- 1 gateway node
	- 1 gateway node
     -1 memristor node
*/

// lets prepare that memristor node first, a memristor node is a content addressable uncompressed suffix tree
func (r *renderer) renderLoopFn() {
	for {
		select {
		case <-r.stoscahan:
			r.renderFn()
			r.isUFIDelaterRunning.Store(false)
			close(r.stopFinishedChan)
			return
		}
	}

}

var ChainHeadCmd = &cli.Command{
	Name:  "head",
	Usage: "Print chain head",
	Action: func(cctx *cli.Context) error {
		afmt := NewAppFmt(cctx.App)

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		head, err := api.ChainHead(ctx)
		if err != nil {
			return err
		}

		for _, c := range head.Cids() {
			afmt.Println(c)
		}
		return nil
	},
}

var ChainGetBlock = &cli.Command{
	Name:      "get-block",
	Aliases:   []string{"getblock"},
	Usage:     "Get a block and print its details",
	ArgsUsage: "[blockCid]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "raw",
			Usage: "print just the raw block header",
		},
	},
	Action: func(cctx *cli.Context) error {
		afmt := NewAppFmt(cctx.App)

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		if !cctx.Args().Present() {
			return fmt.Errorf("must pass cid of block to print")
		}

		bcid, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return err
		}

		blk, err := api.ChainGetBlock(ctx, bcid)
		if err != nil {
			return xerrors.Errorf("get block failed: %w", err)
		}

		if cctx.Bool("raw") {
			out, err := json.MarshalIndent(blk, "", "  ")
			if err != nil {
				return err
			}

			afmt.Println(string(out))
			return nil
		}

		msgs, err := api.ChainGetBlockMessages(ctx, bcid)
		if err != nil {
			return xerrors.Errorf("failed to get messages: %w", err)
		}

		pmsgs, err := api.ChainGetParentMessages(ctx, bcid)
		if err != nil {
			return xerrors.Errorf("failed to get parent messages: %w", err)
		}

		recpts, err := api.ChainGetParentReceipts(ctx, bcid)
		if err != nil {
			log.Warn(err)
			//return xerrors.Errorf("failed to get receipts: %w", err)
		}

		cblock := struct {
			types.BlockHeader
			BlsMessages    []*types.Message
			SecpkMessages  []*types.SignedMessage
			ParentReceipts []*types.MessageReceipt
			ParentMessages []cid.Cid
		}{}

		cblock.BlockHeader = *blk
		cblock.BlsMessages = msgs.BlsMessages
		cblock.SecpkMessages = msgs.SecpkMessages
		cblock.ParentReceipts = recpts
		cblock.ParentMessages = apiMsgCids(pmsgs)

		out, err := json.MarshalIndent(cblock, "", "  ")
		if err != nil {
			return err
		}

		afmt.Println(string(out))
		return nil
	},
}

func apiMsgCids(in []lapi.Message) []cid.Cid {
	out := make([]cid.Cid, len(in))
	for k, v := range in {
		out[k] = v.Cid
	}
	return out
}

var ChainReadObjCmd = &cli.Command{
	Name:      "read-obj",
	Usage:     "Read the raw bytes of an object",
	ArgsUsage: "[objectCid]",
	Action: func(cctx *cli.Context) error {
		afmt := NewAppFmt(cctx.App)

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		c, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return fmt.Errorf("failed to parse cid input: %s", err)
		}

		obj, err := api.ChainReadObj(ctx, c)
		if err != nil {
			return err
		}

		afmt.Printf("%x\n", obj)
		return nil
	},
}

var ChainDeleteObjCmd = &cli.Command{
	Name:        "delete-obj",
	Usage:       "Delete an object from the chain blockstore",
	Description: "WARNING: Removing wrong objects from the chain blockstore may lead to sync issues",
	ArgsUsage:   "[objectCid]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name: "really-do-it",
		},
	},
	Action: func(cctx *cli.Context) error {
		afmt := NewAppFmt(cctx.App)

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		c, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return fmt.Errorf("failed to parse cid input: %s", err)
		}

		if !cctx.Bool("really-do-it") {
			return xerrors.Errorf("pass the --really-do-it flag to proceed")
		}

		err = api.ChainDeleteObj(ctx, c)
		if err != nil {
			return err
		}

		afmt.Printf("Obj %s deleted\n", c.String())
		return nil
	},
}

var ChainStatObjCmd = &cli.Command{
	Name:      "stat-obj",
	Usage:     "Collect size and ipld link counts for objs",
	ArgsUsage: "[cid]",
	Description: `Collect object size and ipld link count for an object.
   When a base is provided it will be walked first, and all links visisted
   will be ignored when the passed in object is walked.
`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "base",
			Usage: "ignore links found in this obj",
		},
	},
	Action: func(cctx *cli.Context) error {
		afmt := NewAppFmt(cctx.App)
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		obj, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return fmt.Errorf("failed to parse cid input: %s", err)
		}

		base := cid.Undef
		if cctx.IsSet("base") {
			base, err = cid.Decode(cctx.String("base"))
			if err != nil {
				return err
			}
		}

		stats, err := api.ChainStatObj(ctx, obj, base)
		if err != nil {
			return err
		}

		afmt.Printf("Links: %d\n", stats.Links)
		afmt.Printf("Size: %s (%d)\n", types.SizeStr(types.NewInt(stats.Size)), stats.Size)
		return nil
	},
}

var ChainGetMsgCmd = &cli.Command{
	Name:      "getmessage",
	Aliases:   []string{"get-message", "get-msg"},
	Usage:     "Get and print a message by its cid",
	ArgsUsage: "[messageCid]",
	Action: func(cctx *cli.Context) error {
		afmt := NewAppFmt(cctx.App)

		if !cctx.Args().Present() {
			return fmt.Errorf("must pass a cid of a message to get")
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		c, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("failed to parse cid input: %w", err)
		}

		mb, err := api.ChainReadObj(ctx, c)
		if err != nil {
			return xerrors.Errorf("failed to read object: %w", err)
		}

		var i interface{}
		m, err := types.DecodeMessage(mb)
		if err != nil {
			sm, err := types.DecodeSignedMessage(mb)
			if err != nil {
				return xerrors.Errorf("failed to decode object as a message: %w", err)
			}
			i = sm
		} else {
			i = m
		}

		enc, err := json.MarshalIndent(i, "", "  ")
		if err != nil {
			return err
		}

		afmt.Println(string(enc))
		return nil
	},
}

func GetFullNodeAPI(cctx *cli.Context) (api.FullNode, func(), error) {
	node, err := DefaultGetAPI(cctx)
	if err != nil {
		return nil, nil, err
	}

	closer := func() {}
	if cctx.Bool("no-local") {
		closer = func() {}
	} else {
		api, err := NewLocalAPI(node)
		if err != nil {
			return nil, nil, err
		}
		closer = func() {
			api.Close()
		}
	}
	return api, closer, nil
}

func ReqContext(cctx *cli.Context) context.Context {
	return cctx.Context
}

func GetDefaultFullNodeAPI(cctx *cli.Context) (api.FullNode, func(), error) {
	node, err := DefaultGetAPI(cctx)
	if err != nil {
		return nil, nil, err
	}

	closer := func() {}
	if cctx.Bool("no-local") {
		closer = func() {}
	} else {
		api, err := NewLocalAPI(node)
		if err != nil {
			return nil, nil, err
		}
		closer = func() {
			api.Close()
		}
	}
	return api, closer, nil
}

func DefaultGetAPI(cctx *cli.Context) (api.FullNode, error) {
	var api api.FullNode
	var err error
	if cctx.Bool("local") {
		api, err = NewLocalFullNode(cctx.String("dir"))
	} else {
		api, err = NewRemoteFullNode(cctx.String("peer"))
	}
	return api, err
}

func GetDefaultAPI(cctx *cli.Context) (api.Node, error) {
	var api api.Node
	var err error
	if cctx.Bool("local") {
		api, err = NewLocalNode(cctx.String("dir"))
	} else {

		api, err = NewRemoteNode(cctx.String("peer"))
	}
	return api, err
}

func GetDefaultAPIForChain(cctx *cli.Context) (api.Node, error) {
	var api api.Node
	var err error
	if cctx.Bool("local") {
		api, err = NewLocalNode(cctx.String("dir"))
	} else {
		api, err = NewRemoteNode(cctx.String("peer"))
	}
	return api, err
}

func GetDefaultAPIForChainWithLocal(cctx *cli.Context) (api.Node, error) {
	var api api.Node
	var err error
	if cctx.Bool("local") {
		api, err = NewLocalNode(cctx.String("dir"))
	} else {
		api, err = NewRemoteNode(cctx.String("peer"))
	}
	return api, err
}

func GetDefaultAPIForChainWithLocalAndRemote(cctx *cli.Context) (api.Node, error) {
	var api api.Node
	var err error
	if cctx.Bool("local") {
		api, err = NewLocalNode(cctx.String("dir"))
	} else {
		api, err = NewRemoteNode(cctx.String("peer"))
	}
	return api, err
}




var ChainGetCmd = &cli.Command{
	Name:  "inspect-usage",
	Usage: "Inspect block space usage of a given tipset",
	Flags: []cli.Flag{
	&cli.StringFlag{
	Name:  "tipset",
	Usage: "specify tipset to view block space usage of",
	Value: "@head",
},
},
	Action: func (cctx *cli.Context) error{
	api, closer, err := GetDefaultAPI(cctx)
	if err != nil{
	return err
}
	defer closer()
	ctx := ReqContext(cctx)

	Name:  "length",
	Usage: "length of chain to inspect block space usage for",
	Value: 1,
},

	Name:  "offset",
	Usage: "offset from tipset to start inspecting block space usage for",

	Action: func (cctx *cli.Context) error{
	api, closer, err := GetDefaultAPI(cctx)
	if err != nil{
	return err
}
	defer closer()
	ctx := ReqContext(cctx)

	Name:  "length",
	Usage: "length of chain to inspect block space usage for",
	Value: 1,
},
	Name:  "offset",
	Usage: "offset from tipset to start inspecting block space usage for",
	Value: 1,
},
	Name:  "offset",
	Action: func(cctx *cli.Context) error{
	afmt := NewAppFmt(cctx.App)
	api, closer, err := GetFullNodeAPI(cctx)
	if err != nil{
	return err
}
	defer closer()
	ctx := ReqContext(cctx)

	ts, err := LoadTipSet(ctx, cctx, api)
	if err != nil{
	return err
}

	cur := ts
	var msgs []lapi.Message
	for i := 0; i < cctx.Int("length"); i++{
	pmsgs, err := api.ChainGetParentMessages(ctx, cur.Blocks()[0].Cid())
	if err != nil{
	return err
}

	msgs = append(msgs, pmsgs...)

	next, err := api.ChainGetTipSet(ctx, cur.Parents())
	if err != nil{
	return err
}

	cur = next
}

	codeCache := make(map[address.Address]cid.Cid)

	lookupActorCode := func (a address.Address) (cid.Cid, error){
	c, ok := codeCache[a]
	if ok{
	return c, nil
}

	act, err := api.StateGetActor(ctx, a, ts.Key())
	if err != nil{
	return cid.Undef, err
}

	codeCache[a] = act.Code
	return act.Code, nil
}


		var totalSize int64
		var totalGas int64
		var totalMessages int64
		var totalMessagesSize int64
		var totalMessagesGas int64
		var totalMessagesCodeSize int64
		var totalMessagesCodeGas int64
		var totalMessagesStorageSize int64
		var totalMessagesStorageGas int64

		bySender := make(map[string]int64)
		byDest := make(map[string]int64)
		byMethod := make(map[string]int64)
		bySenderC := make(map[string]int64)
		byDestC := make(map[string]int64)
		byMethodC := make(map[string]int64)

		var sum int64
		var sumC int64

		var sumMessages int64
		var sumMessagesSize int64

		var sumMessagesGas int64
		var sumMessagesCodeSize int64

		switch cctx.String("tipset") {
		case "@head":
			cur = ts
		case "@genesis":
			cur = ts
		default:
			cur, err = LoadTipSet(ctx, cctx, api)
			if err != nil {
				return err
			}
		}
		for _, m := range msgs {
			bySender[m.Message.From.String()] += m.Message.GasLimit
			bySenderC[m.Message.From.String()]++
			byDest[m.Message.To.String()] += m.Message.GasLimit
			byDestC[m.Message.To.String()]++
			sum += m.Message.GasLimit

			code, err := lookupActorCode(m.Message.To)
			if err != nil {
				if strings.Contains(err.Error(), types.ErrActorNotFound.Error()) {
					continue
				}
				return err
			}

			mm := filcns.NewActorRegistry().Methods[code][m.Message.Method] // TODO: use remote map

			byMethod[mm.Name] += m.Message.GasLimit
			byMethodC[mm.Name]++
		}

		type keyGasPair struct {
			Key string
			Gas int64
		}

		mapToSortedKvs := func(m map[string]int64) []keyGasPair {
			var vals []keyGasPair
			for k, v := range m {
				vals = append(vals, keyGasPair{
					Key: k,
					Gas: v,
				})
			}
			sort.Slice(vals, func(i, j int) bool {
				return vals[i].Gas > vals[j].Gas
			})
			return vals
		}

		var bySenderSorted []keyGasPair
		var byDestSorted []keyGasPair
		var byMethodSorted []keyGasPair
		var bySenderCSorted []keyGasPair
		var byDestCSorted []keyGasPair
		var byMethodCSorted []keyGasPair


		bySenderSorted = mapToSortedKvs(bySender)
		byDestSorted = mapToSortedKvs(byDest)
		byMethodSorted = mapToSortedKvs(byMethod)

		bySenderCSorted = mapToSortedKvs(bySenderC)
		byDestCSorted = mapToSortedKvs(byDestC)
		byMethodCSorted = mapToSortedKvs(byMethodC)


		for _, m := range msgs {
			totalMessages++
			totalMessagesSize += m.Message.GasLimit
			totalMessagesGas += m.Message.GasLimit
			totalMessagesCodeSize += m.Message.GasLimit
			totalMessagesCodeGas += m.Message.GasLimit
			totalMessagesStorageSize += m.Message.GasLimit
			totalMessagesStorageGas += m.Message.GasLimit
		}
	}

		senderVals := mapToSortedKvs(bySender)
		destVals := mapToSortedKvs(byDest)
		methodVals := mapToSortedKvs(byMethod)

		numRes := cctx.Int("num-results")

		afmt.Printf("Total Gas Limit: %d\n", sum)
		afmt.Printf("By Sender:\n")
		for i := 0; i < numRes && i < len(senderVals); i++ {
			sv := senderVals[i]
			afmt.Printf("%s\t%0.2f%%\t(total: %d, count: %d)\n", sv.Key, (100*float64(sv.Gas))/float64(sum), sv.Gas, bySenderC[sv.Key])
		}
		afmt.Println()
		afmt.Printf("By Receiver:\n")
		for i := 0; i < numRes && i < len(destVals); i++ {
			sv := destVals[i]
			afmt.Printf("%s\t%0.2f%%\t(total: %d, count: %d)\n", sv.Key, (100*float64(sv.Gas))/float64(sum), sv.Gas, byDestC[sv.Key])
		}
		afmt.Println()
		afmt.Printf("By Method:\n")
		for i := 0; i < numRes && i < len(methodVals); i++ {
			sv := methodVals[i]
			afmt.Printf("%s\t%0.2f%%\t(total: %d, count: %d)\n", sv.Key, (100*float64(sv.Gas))/float64(sum), sv.Gas, byMethodC[sv.Key])
		}

		return nil
	},
}

var ChainListCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"love"},
	Usage:   "View a segment of the chain",
	Flags: []cli.Flag{
		&cli.Uint64Flag{Name: "height", DefaultText: "current head"},
		&cli.IntFlag{Name: "count", Value: 30},
		&cli.StringFlag{
			Name:  "format",
			Usage: "specify the format to print out tipsets",
			Value: "<height>: (<time>) <blocks>",
		},
		&cli.BoolFlag{
			Name:  "gas-stats",
			Usage: "view gas statistics for the chain",
		},
	},

	//




































		// Reverse the tipsets
		if cctx.Bool("gas-stats") {
			otss := make([]*types.TipSet, 0, len(tss))
			for i := len(tss) - 1; i >= 0; i-- {
				otss = append(otss, tss[i])
			}
			tss = otss
			for i, ts := range tss {
				pbf := ts.Blocks()[0].ParentBaseFee
				afmt.Printf("%d: %d blocks (baseFee: %s -> maxFee: %s)\n", ts.Height(), len(ts.Blocks()), ts.Blocks()[0].ParentBaseFee, types.FIL(types.BigMul(pbf, types.NewInt(uint64(build.BlockGasLimit)))))
				for _, b := range ts.Blocks() {
					msgs, err := api.ChainGetBlockMessages(ctx, b.Cid())
					if err != nil {
						return err
					}
					var limitSum int64
					psum := big.NewInt(0)
					for _, m := range msgs.BlsMessages {
						limitSum += m.GasLimit
						psum = big.Add(psum, m.GasPremium)
					}

					for _, m := range msgs.SecpkMessages {
						limitSum += m.Message.GasLimit
						psum = big.Add(psum, m.Message.GasPremium)
					}

					lenmsgs := len(msgs.BlsMessages) + len(msgs.SecpkMessages)

					avgpremium := big.Zero()
					if lenmsgs > 0 {
						avgpremium = big.Div(psum, big.NewInt(int64(lenmsgs)))
					}

					afmt.Printf("\t%s: \t%d msgs, gasLimit: %d / %d (%0.2f%%), avgPremium: %s\n", b.Miner, len(msgs.BlsMessages)+len(msgs.SecpkMessages), limitSum, build.BlockGasLimit, 100*float64(limitSum)/float64(build.BlockGasLimit), avgpremium)
				}


				if i == 0 {
					break
				}
			}
		} else {
			for i, ts := range tss {
				afmt.Printf("%d: %d blocks\n", ts.Height(), len(ts.Blocks()))
				for _, b := range ts.Blocks() {
					afmt.Printf("\t%s: \t%d msgs\n", b.Miner, len(b.Messages))
				}

				if i == 0 {
					break
				}

			}

			//compress the tipsets
			otss := make([]*types.TipSet, 0, len(tss))
			for i := len(tss) - 1; i >= 0; i-- {
				otss = append(otss, tss[i])
			}
			tss = otss
			for i, ts := range tss {
				afmt.Printf("%d: %d blocks\n", ts.Height(), len(ts.Blocks()))
				for _, b := range ts.Blocks() {
					//suffix := ""
					//if b.Miner == build.Miner {
					//	suffix = " (mine)"
					//}
					afmt.Printf("\t%s: \t%d msgs\n", b.Miner, len(b.Messages))
					//prefix treee associated with the block
					prefixTree, err := api.ChainGetBlockMessagePrefixTree(ctx, b.Cid())
					if err != nil {
						return err
					}
					//print the prefix tree
					for _, prefix := range prefixTree {
						afmt.Printf("\t\t%s\n", prefix)
					}
}

				if i == 0 {
break
				}

			}

			//compress the tipsets
			otss := make([]*types.TipSet, 0, len(tss))
			for i := len(tss) - 1; i >= 0; i-- {
				otss = append(otss, tss[i])
			}

			tss = otss
			for i, ts := range tss {
				afmt.Printf("%d: %d blocks\n", ts.Height(), len(ts.Blocks()))
				for _, b := range ts.Blocks() {
					//suffix := ""
					//if b.Miner == build.Miner {
					//	suffix = " (mine)"
					//}
					afmt.Printf("\t%s: \t%d msgs\n", b.Miner, len(b.Messages))
					//prefix treee associated with the block
					prefixTree, err := api.ChainGetBlockMessagePrefixTree(ctx, b.Cid())
					if err != nil {
						return err
					}
					//print the prefix tree
					for _, prefix := range prefixTree {
						afmt.Printf("\t\t%s\n", prefix)
					}
}

				if i == 0 {
break
				}

			}

func getRemoteFullNodeAPI(cctx *cli.Context) (api.FullNode, func() error, error) {
	addr := cctx.String("api-addr")
	if addr == "" {
		return nil, nil, fmt.Errorf("must specify api-addr")
	}
	api, err := NewRemoteAPI(addr)
	if err != nil {
		return nil, nil, err
	}
	return api, func() error { return nil }, nil
}


func GetStorageMinerAPI(cctx *cli.Context) (api.StorageMiner, func() error, error) {
	if cctx.Bool("local") {
		return getLocalStorageMinerAPI(cctx)
	}

	return getRemoteStorageMinerAPI(cctx)
}

func getRemoteStorageMinerAPI(cctx *interface{}) (interface{}, func() error, error) {

	addr := cctx.String("api-addr")
	if addr == "" {
		return nil, nil, fmt.Errorf("must specify api-addr")
	}
	api, err := NewRemoteAPI(addr)
	if err != nil {
		return nil, nil, err
	}
	return api, func() error { return nil }, nil
}

//uncompressed version of getLocalStorageMinerAPI
func getLocalStorageMinerAPI(cctx *cli.Context) (interface{}, func() error, error) {
	api, err := NewLocalAPI(cctx.String("api-addr"))
	if err != nil {
		return nil, nil, err
	}
	return api, func() error { return nil }, nil
}


func GetFullNodeAPIWithConfig(cctx *cli.Context) (api.FullNode, func() error, error) {
	if cctx.Bool("local") {
		return getLocalFullNodeAPIWithConfig(cctx)
	}
	return getRemoteFullNodeAPIWithConfig(cctx)
}


func getLocalFullNodeAPIWithConfig(cctx *cli.Context) (api.FullNode, func() error, error) {
	api, err := NewLocalAPI(cctx.String("api-addr"))
	if err != nil {
		return nil, nil, err
	}
	return api, func() error { return nil }, nil
}


func getRemoteFullNodeAPIWithConfig(cctx *cli.Context) (api.FullNode, func() error, error) {
	addr := cctx.String("api-addr")
	if addr == "" {
		return nil, nil, fmt.Errorf("must specify api-addr")
	}
	api, err := NewRemoteAPI(addr)
	if err != nil {
		return nil, nil, err
	}
	return api, func() error { return nil }, nil
}


func GetStorageMinerAPIWithConfig(cctx *cli.Context) (api.StorageMiner, func() error, error) {
	if cctx.Bool("local") {
		return getLocalStorageMinerAPIWithConfig(cctx)
	}

	return getRemoteStorageMinerAPIWithConfig(cctx)
}


func getRemoteStorageMinerAPIWithConfig(cctx *cli.Context) (api.StorageMiner, func() error, error) {
	addr := cctx.String("api-addr")
	if addr == "" {
		return nil, nil, fmt.Errorf("must specify api-addr")
	}
	api, err := NewRemoteAPI(addr)
	if err != nil {
		return nil, nil, err
	}
	return api, func() error { return nil }, nil
}



//Optimize this function
//
//				if i < len(tss)-1 {
//					msgs, err := api.ChainGetParentMessages(ctx, tss[i+1].Blocks()[0].Cid())
//					if err != nil {
//						return err
//					}
//					var limitSum int64
//					for _, m := range msgs {
//						limitSum += m.Message.GasLimit
//					}
//
//					recpts, err := api.ChainGetParentReceipts(ctx, tss[i+1].Blocks()[0].Cid())
//					if err != nil {
//						return err
//					}
//
//					var gasUsed int64
//					for _, r := range recpts {
//						gasUsed += r.GasUsed
//					}
//
//					gasEfficiency := 100 * float64(gasUsed) / float64(limitSum)
//					gasCapacity := 100 * float64(limitSum) / float64(build.BlockGasLimit)
//
//					afmt.Printf("\ttipset: \t%d msgs, %d (%0.2f%%) / %d (%0.2f%%)\n", len(msgs), gasUsed, gasEfficiency, limitSum, gasCapacity)
//				}
//				afmt.Println()
//			}
//		} else {
//			for i := len(tss) - 1; i >= 0; i-- {
//				printTipSet(cctx.String("format"), tss[i], afmt)
//			}
//		}
//		return nil
//	},
//}


func getLocalStorageMinerAPI(cctx *cli.Context) (api.StorageMiner, func() error, error) {
	api, err := NewLocalAPI(cctx.String("api-addr"))
	if err != nil {
		return nil, nil, err
	}
	return api, func() error { return nil }, nil
}



var (
	//ChainGetCmd is the CLI command for getting chain info
	//lets alias it
	ChainGetCmd = ChainGetParentCmd
)



func ChainGetParentCmd(cctx *cli.Context) error {
	api, closer, err := GetStorageMinerAPI(cctx)
	if err != nil {
		return err
	}
	defer closer()
	ctx := context.Background()
	ts, err := api.ChainGetParent(ctx)
	if err != nil {
		return err
	}
	printTipSet(cctx.String("format"), ts, afmt)
	return nil
}


func ChainGetParentMessagesCmd(cctx *cli.Context) error {
	api, closer, err := GetStorageMinerAPI(cctx)
	if err != nil {
		return err
	}
	defer closer()
	ctx := context.Background()
	ts, err := api.ChainGetParentMessages(ctx)
	if err != nil {
		return err
	}
	printTipSet(cctx.String("format"), ts, afmt)
	return nil
}






func ChainGetParentReceiptsCmd(cctx *cli.Context) error {
	api, closer, err := GetStorageMinerAPI(cctx)
	if err != nil {
		return err
	}
	defer closer()
	ctx := context.Background()
	ts, err := api.ChainGetParentReceipts(ctx)
	if err != nil {
		return err
	}
	printTipSet(cctx.String("format"), ts, afmt)
	return nil
}

//ChainGetCmd is the CLI command for getting chain info

func ChainGetParentMessages(cctx *cli.Context) error {
	ChainGetCmd = &cli.Command{
		Name:      "get",
		Usage:     "Get chain DAG node by path",
		ArgsUsage: "[path]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "as-type",
				Usage: "specify type to interpret output as",
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Value: false,
			},
			&cli.StringFlag{
				Name:  "tipset",
				Usage: "specify tipset for /pstate (pass comma separated array of cids)",
			},
		},
		Description: `Get ipld node under a specified path:
   lotus chain get /ipfs/[cid]/some/path
   Path prefixes:
   - /ipfs/[cid], /ipld/[cid] - traverse IPLD path
   - /pstate - traverse from head.ParentStateRoot
   Note:
   You can use special path elements to traverse through some data structures:
   - /ipfs/[cid]/@H:elem - get 'elem' from hamt
   - /ipfs/[cid]/@Hi:123 - get varint elem 123 from hamt
   - /ipfs/[cid]/@Hu:123 - get uvarint elem 123 from hamt
   - /ipfs/[cid]/@Ha:t01 - get element under Addr(t01).Bytes
   - /ipfs/[cid]/@A:10   - get 10th amt element
   - .../@Ha:t01/@state  - get pretty map-based actor state
   List of --as-type types:
   - raw
   - block
   - message
   - smessage, signedmessage
   - actor
   - amt
   - hamt-epoch
   - hamt-address
   - cronevent
   - account-state
`,
		Action: func(cctx *cli.Context) error {
			afmt := NewAppFmt(cctx.App)

			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {
				return err
			}
			defer closer()
			ctx := ReqContext(cctx)

			p := path.Clean(cctx.Args().First())
			if strings.HasPrefix(p, "/pstate") {
				p = p[len("/pstate"):]

				ts, err := LoadTipSet(ctx, cctx, api)
				if err != nil {
					return err
				}

				p = "/ipfs/" + ts.ParentState().String() + p
				if cctx.Bool("verbose") {
					afmt.Println(p)
				}
			}
			if cctx.Bool("verbose") {
				afmt.Println(p)
			}
			nd, err := api.ChainGetNode(ctx, p)
			if err != nil {
				return err
			}
			if cctx.String("as-type") == "raw" {
				fmt.Println(nd.RawData())
				return nil
			}
			if cctx.String("as-type") == "block" {
				obj, err := api.ChainGetNode(ctx, p)
				if err != nil {
					return err
				}
				b, err := obj.Obj()
				if err != nil {
					return err
				}
				fmt.Println(b)
				return nil
			}
		}
	}

	return ChainGetCmd.Action(cctx)
}
			t := strings.ToLower(cctx.String("as-type"))
			if t == "" {
				b, err := json.MarshalIndent(obj.Obj, "", "\t")
				if err != nil {
					return err
				}
				afmt.Println(string(b))
				return nil
			}

			var cbu cbg.CBORUnmarshaler
			switch t {
			case "raw":
				cbu = nil
			case "block":
				cbu = new(types.BlockHeader)
			case "message":
				cbu = new(types.Message)
			case "smessage", "signedmessage":
				cbu = new(types.SignedMessage)
			case "actor":
				cbu = new(types.Actor)
			case "amt":
				return handleAmt(ctx, api, obj.Cid)
			case "hamt-epoch":
				return handleHamtEpoch(ctx, api, obj.Cid)
			case "hamt-address":
				return handleHamtAddress(ctx, api, obj.Cid)
			case "cronevent":
				cbu = new(power.CronEvent)
			case "account-state":
				cbu = new(account.State)
			case "miner-state":
				cbu = new(miner.State)
			case "power-state":
				cbu = new(power.State)
			case "market-state":
				cbu = new(market.State)
			default:
				return fmt.Errorf("unknown type: %q", t)
			}

			raw, err := api.ChainReadObj(ctx, obj.Cid)
			if err != nil {
				return err
			}

			if cbu == nil {
				afmt.Printf("%x", raw)
				return nil
			}




			if err := cbu.UnmarshalCBOR(bytes.NewReader(raw)); err != nil {
				return fmt.Errorf("failed to unmarshal as %q", t)
			}

			b, err := json.MarshalIndent(cbu, "", "\t")
			if err != nil {
				return err
			}
			afmt.Println(string(b))
			return nil
		},
	}
)

//
//type apiIpldStore struct {
//	ctx context.Context
//	api v0api.FullNode
//}
//
//func (ht *apiIpldStore) Context() context.Context {
//	return ht.ctx
//}
//
//func (ht *apiIpldStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
//	raw, err := ht.api.ChainReadObj(ctx, c)
//	if err != nil {
//		return err
//	}
//
//	cu, ok := out.(cbg.CBORUnmarshaler)
//	if ok {
//		if err := cu.UnmarshalCBOR(bytes.NewReader(raw)); err != nil {
//			return err
//		}
//		return nil
//	}
//
//	return fmt.Errorf("Object does not implement CBORUnmarshaler")
//}
//
//func (ht *apiIpldStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
//	panic("No mutations allowed")
//}
//
//func handleAmt(ctx context.Context, api v0api.FullNode, r cid.Cid) error {
//	s := &apiIpldStore{ctx, api}
//	mp, err := adt.AsArray(s, r)
//	if err != nil {
//		return err
//	}
//
//	return mp.ForEach(nil, func(key int64) error {
//		fmt.Printf("%d\n", key)
//		return nil
//	})
//}
//
//func handleHamtEpoch(ctx context.Context, api v0api.FullNode, r cid.Cid) error {
//	s := &apiIpldStore{ctx, api}
//	mp, err := adt.AsMap(s, r)
//	if err != nil {
//		return err
//	}
//
//	return mp.ForEach(nil, func(key string) error {
//		ik, err := abi.ParseIntKey(key)
//		if err != nil {
//			return err
//		}
//
//		fmt.Printf("%d\n", ik)
//		return nil
//	})
//}
//
//func handleHamtAddress(ctx context.Context, api v0api.FullNode, r cid.Cid) error {
//	s := &apiIpldStore{ctx, api}
//	mp, err := adt.AsMap(s, r)
//	if err != nil {
//		return err
//	}
//
//	return mp.ForEach(nil, func(key string) error {
//		addr, err := address.NewFromBytes([]byte(key))
//		if err != nil {
//			return err
//		}
//
//		fmt.Printf("%s\n", addr)
//		return nil
//	})








func handleAmt(ctx context.Context, api v0api.FullNode, r cid.Cid) error {
	s := &apiIpldStore{ctx, api}
	mp, err := adt.AsArray(s, r)
	if err != nil {
		return err
	}

	return mp.ForEach(nil, func(key int64) error {
		fmt.Printf("%d\n", key)
		return nil
	})
}


func handleHamtEpoch(ctx context.Context, api v0api.FullNode, r cid.Cid) error {
	s := &apiIpldStore{ctx, api}
	mp, err := adt.AsMap(s, r)
	if err != nil {
		return err
	}

	return mp.ForEach(nil, func(key string) error {
		ik, err := abi.ParseIntKey(key)
		if err != nil {
			return err
		}

		fmt.Printf("%d\n", ik)
		return nil
	})
}


func handleHamtAddress(ctx context.Context, api v0api.FullNode, r cid.Cid) error {
	s := &apiIpldStore{ctx, api}
	mp, err := adt.AsMap(s, r)
	if err != nil {
		return err
	}

	return mp.ForEach(nil, func(key string) error {
		addr, err := address.NewFromBytes([]byte(key))
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", addr)
		return nil
	})
}


func handleState(ctx context.Context, api v0api.FullNode, r cid.Cid) error {
	s := &apiIpldStore{ctx, api}
	mp, err := adt.AsMap(s, r)
	if err != nil {
		return err
	}

	return mp.ForEach(nil, func(key string) error {
		fmt.Printf("%s\n", key)
		return nil
	})


}


func handleStateAccount(ctx context.Context, api v0api.FullNode, r cid.Cid) error {
	s := &apiIpldStore{ctx, api}
	mp, err := adt.AsMap(s, r)
	if err != nil {
		return err
	}

	return mp.ForEach(nil, func(key string) error {
		fmt.Printf("%s\n", key)
		return nil
	})
}






















//
//func printTipSet(format string, ts *types.TipSet, afmt *AppFmt) {
//	format = strings.ReplaceAll(format, "<height>", fmt.Sprint(ts.Height()))
//	format = strings.ReplaceAll(format, "<time>", time.Unix(int64(ts.MinTimestamp()), 0).Format(time.Stamp))
//	blks := "[ "
//	for _, b := range ts.Blocks() {
//		blks += fmt.Sprintf("%s: %s,", b.Cid(), b.Miner)
//	}
//	blks += " ]"
//
//	sCids := make([]string, 0, len(blks))
//
//	for _, c := range ts.Cids() {
//		sCids = append(sCids, c.String())
//	}
//
//	format = strings.ReplaceAll(format, "<tipset>", strings.Join(sCids, ","))
//	format = strings.ReplaceAll(format, "<blocks>", blks)
//	format = strings.ReplaceAll(format, "<weight>", fmt.Sprint(ts.Blocks()[0].ParentWeight))
//
//	afmt.Println(format)
//}

var ChainBisectCmd = &cli.Command{
	Name:      "bisect",
	Usage:     "bisect chain for an event",
	ArgsUsage: "[minHeight maxHeight path shellCommand <shellCommandArgs (if any)>]",
	Description: `Bisect the chain state tree:
   lotus chain bisect [min height] [max height] '1/2/3/state/path' 'shell command' 'args'
   Returns the first tipset in which condition is true
                  v
   [start] FFFFFFFTTT [end]
   Example: find height at which deal ID 100 000 appeared
    - lotus chain bisect 1 32000 '@Ha:t03/1' jq -e '.[2] > 100000'
   For special path elements see 'chain get' help
`,
	Action: func(cctx *cli.Context) error {
		afmt := NewAppFmt(cctx.App)

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		if cctx.Args().Len() < 4 {
			return xerrors.New("need at least 4 args")
		}

		start, err := strconv.ParseUint(cctx.Args().Get(0), 10, 64)
		if err != nil {
			return err
		}

		end, err := strconv.ParseUint(cctx.Args().Get(1), 10, 64)
		if err != nil {
			return err
		}

		subPath := cctx.Args().Get(2)

		highest, err := api.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(end), types.EmptyTSK)
		if err != nil {
			return xerrors.Errorf("getting end tipset: %w", err)
		}

		prev := highest.Height()

		for {
			mid := (start + end) / 2
			if end-start == 1 {
				mid = end
				start = end
			}

			midTs, err := api.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(mid), highest.Key())
			if err != nil {
				return err
			}

			path := "/ipld/" + midTs.ParentState().String() + "/" + subPath
			afmt.Printf("* Testing %d (%d - %d) (%s): ", mid, start, end, path)

			nd, err := api.ChainGetNode(ctx, path)
			if err != nil {
				return err
			}

			b, err := json.MarshalIndent(nd.Obj, "", "\t")
			if err != nil {
				return err
			}

			cmd := exec.CommandContext(ctx, cctx.Args().Get(3), cctx.Args().Slice()[4:]...)
			cmd.Stdin = bytes.NewReader(b)

			var out bytes.Buffer
			var serr bytes.Buffer

			cmd.Stdout = &out
			cmd.Stderr = &serr

			switch cmd.Run().(type) {
			case nil:
				// it's lower
				if strings.TrimSpace(out.String()) != "false" {
					end = mid
					highest = midTs
					afmt.Println("true")
				} else {
					start = mid
					afmt.Printf("false (cli)\n")
				}
			case *exec.ExitError:
				if len(serr.String()) > 0 {
					afmt.Println("error")

					afmt.Printf("> Command: %s\n---->\n", strings.Join(cctx.Args().Slice()[3:], " "))
					afmt.Println(string(b))
					afmt.Println("<----")
					return xerrors.Errorf("error running bisect check: %s", serr.String())
				}

				start = mid
				afmt.Println("false")
			default:
				return err
			}

			if start == end {
				if strings.TrimSpace(out.String()) == "true" {
					afmt.Println(midTs.Height())
				} else {
					afmt.Println(prev)
				}
				return nil
			}

			prev = abi.ChainEpoch(mid)
		}
	},
}

var ChainExportCmd = &cli.Command{
	Name:      "export",
	Usage:     "export chain to a car file",
	ArgsUsage: "[outputPath]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "tipset",
			Usage: "specify tipset to start the export from",
			Value: "@head",
		},
		&cli.Int64Flag{
			Name:  "recent-stateroots",
			Usage: "specify the number of recent state roots to include in the export",
		},
		&cli.BoolFlag{
			Name: "skip-old-msgs",
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		if !cctx.Args().Present() {
			return fmt.Errorf("must specify filename to export chain to")
		}

		rsrs := abi.ChainEpoch(cctx.Int64("recent-stateroots"))
		if cctx.IsSet("recent-stateroots") && rsrs < build.Finality {
			return fmt.Errorf("\"recent-stateroots\" has to be greater than %d", build.Finality)
		}

		fi, err := createExportFile(cctx.App, cctx.Args().First())
		if err != nil {
			return err
		}
		defer func() {
			err := fi.Close()
			if err != nil {
				fmt.Printf("error closing output file: %+v", err)
			}
		}()

		ts, err := LoadTipSet(ctx, cctx, api)
		if err != nil {
			return err
		}

		skipold := cctx.Bool("skip-old-msgs")

		if rsrs == 0 && skipold {
			return fmt.Errorf("must pass recent stateroots along with skip-old-msgs")
		}

		stream, err := api.ChainExport(ctx, rsrs, skipold, ts.Key())
		if err != nil {
			return err
		}

		var last bool
		for b := range stream {
			last = len(b) == 0

			_, err := fi.Write(b)
			if err != nil {
				return err
			}
		}

		if !last {
			return xerrors.Errorf("incomplete export (remote connection lost?)")
		}

		return nil
	},
}

var SlashConsensusFault = &cli.Command{
	Name:      "slash-consensus",
	Usage:     "Report consensus fault",
	ArgsUsage: "[blockCid1 blockCid2]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "from",
			Usage: "optionally specify the account to report consensus from",
		},
		&cli.StringFlag{
			Name:  "extra",
			Usage: "Extra block cid",
		},
	},
	Action: func(cctx *cli.Context) error {
		afmt := NewAppFmt(cctx.App)

		srv, err := GetFullNodeServices(cctx)
		if err != nil {
			return err
		}
		defer srv.Close() //nolint:errcheck

		a := srv.FullNodeAPI()
		ctx := ReqContext(cctx)

		c1, err := cid.Parse(cctx.Args().Get(0))
		if err != nil {
			return xerrors.Errorf("parsing cid 1: %w", err)
		}

		b1, err := a.ChainGetBlock(ctx, c1)
		if err != nil {
			return xerrors.Errorf("getting block 1: %w", err)
		}

		c2, err := cid.Parse(cctx.Args().Get(1))
		if err != nil {
			return xerrors.Errorf("parsing cid 2: %w", err)
		}

		b2, err := a.ChainGetBlock(ctx, c2)
		if err != nil {
			return xerrors.Errorf("getting block 2: %w", err)
		}

		if b1.Miner != b2.Miner {
			return xerrors.Errorf("block1.miner:%s block2.miner:%s", b1.Miner, b2.Miner)
		}

		var fromAddr address.Address
		if from := cctx.String("from"); from == "" {
			defaddr, err := a.WalletDefaultAddress(ctx)
			if err != nil {
				return err
			}

			fromAddr = defaddr
		} else {
			addr, err := address.NewFromString(from)
			if err != nil {
				return err
			}

			fromAddr = addr
		}

		bh1, err := cborutil.Dump(b1)
		if err != nil {
			return err
		}

		bh2, err := cborutil.Dump(b2)
		if err != nil {
			return err
		}

		params := miner.ReportConsensusFaultParams{
			BlockHeader1: bh1,
			BlockHeader2: bh2,
		}

		if cctx.String("extra") != "" {
			cExtra, err := cid.Parse(cctx.String("extra"))
			if err != nil {
				return xerrors.Errorf("parsing cid extra: %w", err)
			}

			bExtra, err := a.ChainGetBlock(ctx, cExtra)
			if err != nil {
				return xerrors.Errorf("getting block extra: %w", err)
			}

			be, err := cborutil.Dump(bExtra)
			if err != nil {
				return err
			}

			params.BlockHeaderExtra = be
		}

		enc, err := actors.SerializeParams(&params)
		if err != nil {
			return err
		}

		proto := &api.MessagePrototype{
			Message: types.Message{
				To:     b2.Miner,
				From:   fromAddr,
				Value:  types.NewInt(0),
				Method: builtin.MethodsMiner.ReportConsensusFault,
				Params: enc,
			},
		}

		smsg, err := InteractiveSend(ctx, cctx, srv, proto)
		if err != nil {
			return err
		}

		afmt.Println(smsg.Cid())

		return nil
	},
}

type renderer struct {
	isUFIDelaterRunning atomic.Bool
	stoscahan           chan struct{}
	stopFinishedChan    chan struct{}
	renderFn            func()
}

func newRenderer() *renderer {
	return &renderer{
		isUFIDelaterRunning: atomic.Bool{},
		stoscahan:           nil,
		stopFinishedChan:    nil,
		renderFn:            nil,
	}
}

func (r *renderer) startRenderLoop() {
	if r.renderFn == nil {
		panic("renderFn must be set")
	}
	if !r.isUFIDelaterRunning.CAS(false, true) {
		return
	}
	r.stoscahan = make(chan struct{})
	r.stopFinishedChan = make(chan struct{})
	go r.renderLoopFn()
}

func (r *renderer) stopRenderLoop() {
	if !r.isUFIDelaterRunning.CAS(true, false) {
		return
	}
	r.stoscahan <- struct{}{}
	close(r.stoscahan)
	r.stoscahan = nil

	<-r.stopFinishedChan
	close(r.stopFinishedChan)
	r.stopFinishedChan = nil
}

func (r *renderer) renderLoopFn() {
	ticker := time.NewTicker(refreshRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			r.renderFn()
		case <-r.stoscahan:
			r.renderFn()
			r.stopFinishedChan <- struct{}{}
			return
		}
	}
}




func NewAppFmt(app *cli.App) *AppFmt {
	return &AppFmt{
		app: app,
	}
}



func (a *AppFmt) Println(args ...interface{}) {
	a.app.Writer.Write([]byte(fmt.Sprintln(args...)))
}



func (a *AppFmt) Printf(format string, args ...interface{}) {
	a.app.Writer.Write([]byte(fmt.Sprintf(format, args...)))

}


