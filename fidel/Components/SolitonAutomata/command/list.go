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
	address "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	cid "github.com/ipfs/go-cid"
)

/*
Block rewards

Block rewards are large sums that are given to the storage provider credited for a new block. Unlike storage fees, these rewards do not come from an associated client; rather, the network “prints” new FIL as both an inflationary measure and an incentive to providers advancing the chain. All active storage providers on the network have a chance at receiving a block reward, their chance at such being directly proportional to the amount of storage space currently being contributed to the network.

The mechanism to earn the right to provide a new block is called WinningPoSt. In the Filecoin network, time is discretized into a series of epochs – the blockchain’s height corresponds to the number of elapsed epochs. At the beginning of each epoch, a small number of storage providers are elected to provide new blocks. Additionally to the block reward, each storage provider can collect the fees associated to each message included in the block.

The number of blocks on every tipset is based on a Poisson distribution of a random variable with λ = 5. Provider implementations may use several strategies to choose which messages to include in every block to minimize overlap. Only the “first execution” of each message will collect the associated fees, with executions ordered per the hash of the VRF (Verifiable Random Function) ticket associated to the block.

Verified clients

To further incentivize the storage of “useful” data over simple capacity commitments, storage provider have the additional opportunity to compete for special deals offered by verified clients. Such clients are certified with respect to their intent to offer deals involving the storage of meaningful data, and the power a storage provider earns for these deals is augmented by a multiplier. The total amount of power a given storage provider has, after accounting for this multiplier, is known as quality-adjusted power.


*/

// /////////////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////////////////////////////////////////////////////////////////////////////

const (
	// SolitonAutomataListCmd is the command to list all solitonAutomatas.
	SolitonAutomataListCmd = "list"
)

type FilCmdState struct {
	cid                 cid.Cid
	addr                address.Address
	balance             big.Int
	Signer              []address.Address // signer address
	NumApproved         uint64
	NumRejected         uint64
	NumPending          uint64
	NumSettled          uint64
	NumTotal            uint64
	UnlockHeight        uint64
	Status              string
	PendingSolitonTxns  []cid.Cid
	SettledSolitonTxns  []cid.Cid
	RejectedSolitonTxns []cid.Cid
}

func (s *FilCmdState) String() string {
	return s.cid.String()
}

func (s *FilCmdState) Cid() cid.Cid {
	return s.cid
}

func (s *FilCmdState) Address() address.Address {
	return s.addr
}

func (s *FilCmdState) Balance() big.Int {
	return s.balance
}

func (s *FilCmdState) Signer() []address.Address {
	return s.Signer
}

type IpfsWithCeph struct {
	Ipfs string
	Ceph string
}

type Topology struct {
	GlobalOptions    spec.GlobalOptions
	MonitoredOptions spec.MonitoredOptions
	InterlockOptions spec.InterlockOptions
	FoundationDBs    []spec.FoundationDBSpec
}

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all solitonAutomatas",
		RunE: func(cmd *cobra.Command, args []string) error {
			return manager.ListSolitonAutomata()
		},
	}
	return cmd
}

type ManagerInterface interface {
	//WithEinsteinDBIpfs
	InitiateSolitonAutomata(cid cid.Cid) error
	ListSolitonAutomata() error
	//WithoutIpfs
	InitiateSolitonAutomataWithoutIpfs(cid cid.Cid) error
	//Ceph
	InitiateSolitonAutomataWithCeph(cid cid.Cid, ceph string) error
	//Ipfs
	InitiateSolitonAutomataWithIpfs(cid cid.Cid, ipfs string) error
	//Ipfs and Ceph
	InitiateSolitonAutomataWithIpfsAndCeph(cid cid.Cid, ipfs string, ceph string) error
	//rook v1
	InitiateSolitonAutomataWithRookV1(cid cid.Cid, ipfs string, ceph string) error
	//rook v2
	InitiateSolitonAutomataWithRookV2(cid cid.Cid, ipfs string, ceph string) error
	//rook v3
	InitiateSolitonAutomataWithRookV3(cid cid.Cid, ipfs string, ceph string) error
	//isovalent
	InitiateSolitonAutomataWithIsovalent(cid cid.Cid, ipfs string, ceph string) error
	//isovalent v2

}

type Manager struct {
	Topology *Topology
	//IpfsWithCeph *IpfsWithCeph
	//Ipfs string
	//Ceph string

}

func (m *Manager) InitiateSolitonAutomata(cid cid.Cid) error {
	// 1. get ipfs and ceph
	// 2. initiate solitonAutomata
	// 3. return error
	ipfs := m.Topology.GlobalOptions.Ipfs
	ceph := m.Topology.GlobalOptions.Ceph
	return m.InitiateSolitonAutomataWithIpfsAndCeph(cid, ipfs, ceph)
}

func (m *Manager) InitiateSolitonAutomataWithoutIpfs(cid cid.Cid) error {
	return nil
}

func (m *Manager) InitiateSolitonAutomataWithCeph(cid cid.Cid, ceph string) error {
	return nil
}

// Here we use ipfs and ceph to initiate solitonAutomata
func (m *Manager) InitiateSolitonAutomataWithIpfsAndCeph(cid cid.Cid, ipfs string, ceph string) error {

	// Provision
	// 1. get ipfs and ceph
	// 2. initiate solitonAutomata
	// 3. return error

}
