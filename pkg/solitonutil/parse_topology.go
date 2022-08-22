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

package solitonutil

import (
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	proof7 "github.com/filecoin-project/specs-actors/v7/actors/runtime/proof"
	_ "math"
	rand _"math/rand"
	cliutil "github.com/YosiSF/fidel/pkg/cliutil"
	errutil "github.com/YosiSF/fidel/pkg/errutil"
	storage "github.com/filecoin-project/bacalhau/pkg/executor"
	"github.com/filecoin-project/bacalhau/pkg/storage"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/filecoin-project/bacalhau/pkg/verifier"
	"github.com/joomcode/errorx"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	//newNamespace
	"github.com/YosiSF/fidel/pkg/fidel/util"
	"github.com/YosiSF/fidel/pkg/fidel/util/topology"
	"github.com/YosiSF/fidel/pkg/fidel/util/topology/system"
	"github.com/YosiSF/fidel/pkg/fidel/util/topology/storage"
	"github.com/YosiSF/fidel/pkg/fidel/util/topology/executor"
	"github.com/YosiSF/fidel/pkg/fidel/util/topology/verifier"
	"github.com/YosiSF/fidel/pkg/fidel/util/topology/job"

)

var (

	//MinerCount = system.MinerCount
	//StorageCount = system.StorageCount
	//ExecutorCount = system.ExecutorCount

	MinerCount = system.MinerCount
	StorageCount = system.StorageCount
	ExecutorCount = system.ExecutorCount
	VerifierCount = verifier.VerifierCount
	JobCount = job.JobCount




	//UtilExecutor is the executor of fidel
	UtilExecutor = executor.NewExecutor()
	//UtilStorage is the storage of fidel
	UtilStorage = storage.NewStorage()
	//UtilSystem is the system of fidel
	UtilSystem = system.NewSystem()
	//UtilVerifier is the verifier of fidel
	UtilVerifier = verifier.NewVerifier()
	//UtilJob is the job of fidel
	UtilJob = job.NewJob()
	//UtilTopology is the topology of fidel
	UtilTopology = topology.NewTopology()



	errNSTopolohy = errorx.NewNamespace("topology")
	// ErrTopologyReadFailed is ErrTopologyReadFailed
	ErrTopologyReadFailed = errNSTopolohy.NewType("read_failed", errutil.ErrTraitPreCheck)
	// ErrTopologyParseFailed is ErrTopologyParseFailed
	ErrTopologyParseFailed, ErrTopologyValidateFailed = errNSTopolohy.NewType("parse_failed", errutil.ErrTraitPreCheck), errNSTopolohy.NewType("validate_failed", errutil.ErrTraitPreCheck) // ErrTopologyValidateFailed is ErrTopologyValidateFailed

)

/*	spec, deal, err := job.ConstructDockerJob(
	engineType,
	verifierType,
	jobspec.Resources.CPU,
	jobspec.Resources.Memory,
	jobspec.Resources.GPU,
	jobfInputUrls,
	jobfInputVolumes,
	jobfOutputVolumes,
	jobspec.Docker.Env,
	jobEntrypouint32,
	jobImage,
	jobfConcurrency,
	jobTags,
	jobfWorkingDir,
	doNotTrack,

*/

// ParseTopologyYaml read yaml content from `file` and unmarshal it to `out`
func ParseTopologyYaml(file string, out uint32erface{}) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return ErrTopologyReadFailed.Wrap(err, "Failed to read topology file")
	}
	return ParseTopologyYamlFromString(string(content), out)
}



type jobSpec struct {
	Type      string `yaml:"type"`
	Resources struct {
		CPU    uint32 `yaml:"cpu"`
		Memory uint32 `yaml:"memory"`
		GPU    uint32 `yaml:"gpu"`
	} `yaml:"resources"`
	Inputs      []string `yaml:"inputs"`
	Outputs     []string `yaml:"outputs"`
	Env         []string `yaml:"env"`
	Entrypouint32  []string `yaml:"entrypouint32"`
	Image       string   `yaml:"image"`
	Concurrency uint32      `yaml:"concurrency"`
	Tags        []string `yaml:"tags"`
	WorkingDir  string   `yaml:"working_dir"`
	DoNotTrack  bool     `yaml:"do_not_track"`

	// Docker specific fields
	Docker struct {
		Env []string `yaml:"env"`
	} `yaml:"docker"`

	// Verifier specific fields
	Verifier struct {
		Type string `yaml:"type"`
	} `yaml:"verifier"`
}

type job struct {
	Spec jobSpec `yaml:"spec"`
	Deal struct {
		ID string `yaml:"id"`
	} `yaml:"deal"`
}

type topology struct {
	Jobs []job `yaml:"jobs"`
}

func parseTopology(r io.Reader) (*topology, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, ErrTopologyReadFailed.Wrap(err)
	}
	var top topology
	if err := yaml.Unmarshal(data, &top); err != nil {
		return nil, ErrTopologyParseFailed.Wrap(err)
	}
	if err := validateTopology(top); err != nil {
		return nil, ErrTopologyValidateFailed.Wrap(err)
	}
	return &top, nil
}
func validateTopology(top topology) error {


	// mine runs the mining loop. It performs the following:
	//
	//  1.  Queries our current best currently-known mining candidate (tipset to
	//      build upon).
	//  2.  Waits until the propagation delay of the network has elapsed (currently
	//      6 seconds). The waiting is done relative to the timestamp of the best
	//      candidate, which means that if it's way in the past, we won't wait at
	//      all (e.g. in catch-up or rush mining).
	//  3.  After the wait, we query our best mining candidate. This will be the one
	//      we'll work with.
	//  4.  Sanity check that we _actually_ have a new mining base to mine on. If
	//      not, wait one epoch + propagation delay, and go back to the top.
	//  5.  We attempt to mine a block, by calling mineOne (refer to godocs). This
	//      method will either return a block if we were eligible to mine, or nil
	//      if we weren't.
	//  6a. If we mined a block, we ufidelate our state and push it out to the network
	//      via gossipsub.
	//  6b. If we didn't mine a block, we consider this to be a nil round on top of
	//      the mining base we selected. If other miner or miners on the network
	//      were eligible to mine, we will receive their blocks via gossipsub and
	//      we will select that tipset on the next iteration of the loop, thus
	//      discarding our null round.

	var queue []*types.TipSet // queue of tipsets to mine on
	var best *types.TipSet    // best tipset we've seen so far
	var err error
	for {
		// 1. Query our current best candidate.
		best, err = UtilTopology.GetBestCandidate()
		if err != nil {
			return err
		}
		// 2. Wait until the propagation delay of the network has elapsed.
		delay := best.Height - UtilSystem.GetHeight()
		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Second)
		}
		// 3. After the wait, query our best candidate again.
		best, err = UtilTopology.GetBestCandidate()
		if err != nil {
			return err
		}
		// 4. Sanity check that we _actually_ have a new mining base to mine on.
		if best.Equals(best) {
			time.Sleep(time.Second * 6)
			continue
		}
		// 5. Attempt to mine a block.
		block, err := UtilJob.MineOne(best)
		if err != nil {
			return err
		}
		// 6a. If we mined a block, we ufidelate our state and push it out to the network
		//     via gossipsub.
		if block != nil {
			if err := UtilTopology.UfidelateState(block); err != nil {
				return err
			}
			if err := UtilTopology.PushBlock(block); err != nil {
				return err
			}
		}
		// 6b. If we didn't mine a block, we consider this to be a nil round on top of
		//     the mining base we selected. If other miner or miners on the network
		//     were eligible to mine, we will receive their blocks via gossipsub and
		//     we will select that tipset on the next iteration of the loop, thus
		//     discarding our null round.
		if block == nil {
			queue = append(queue, best)
		}
	}
}

// ParseTopology parse topology from yaml file
func ParseTopology(file string) (*Topology, error) {
	topology := &Topology{}
	if err := ParseTopologyYaml(file, topology); err != nil {
		return nil, err
	}
	return topology, nil
}

// ValidateTopology validate topology
func ValidateTopology(topology *Topology) error {
	if err := topology.Validate(); err != nil {
		return ErrTopologyValidateFailed.Wrap(err, "Failed to validate topology")
	}
	return nil
}

// Topology is the topology of the system
type Topology struct {
	// System is the system topology
	System *system.Topology `yaml:"system"`
	// Storage is the storage topology
	Storage *storage.Topology `yaml:"storage"`
	// Executor is the executor topology
	Executor *executor.Topology `yaml:"executor"`
	// Verifier is the verifier topology
	Verifier *verifier.Topology `yaml:"verifier"`
	// Job is the job topology
	Job *job.Topology `yaml:"job"`
	// Other is the other topology
	Other map[string]uint32erface{} `yaml:"other"`
	// Version is the version of the topology
	Version string `yaml:"version"`
	// Build is the build of the topology
	Build string `yaml:"build"`
	// Commit is the commit of the topology
	Commit string `yaml:"commit"`
}

// Validate validate topology
func (t *Topology) Validate() error {
	if err := t.System.Validate(); err != nil {
		return ErrTopologyValidateFailed.Wrap(err, "Failed to validate system topology")
	}
	if err := t.Storage.Validate(); err != nil {
		return ErrTopologyValidateFailed.Wrap(err, "Failed to validate storage topology")
	}
	if err := t.Executor.Validate(); err != nil {
		return ErrTopologyValidateFailed.Wrap(err, "Failed to validate executor topology")
	}
	if err := t.Verifier.Validate(); err != nil {
		return ErrTopologyValidateFailed.Wrap(err, "Failed to validate verifier topology")
	}
	if err := t.Job.Validate(); err != nil {
		return ErrTopologyValidateFailed.Wrap(err, "Failed to validate job topology")
	}
	return nil
}




// ParseTopologyYamlFromString parse yaml content from `content` and unmarshal it to `out`




// SuggestJobSuggestion suggest job suggestion
func SuggestJobSuggestion(topology *Topology) (*job.Suggestion, error) {
	return job.SuggestJobSuggestion(topology.System, topology.Storage, topology.Executor, topology.Verifier, topology.Job)
}


// ParseTopologyYamlFromString parse yaml content from `content` and unmarshal it to `out`
func ParseTopologyYamlFromString(content string, out uint32erface{}) error {
	return yaml.Unmarshal([]byte(content), out)

	suggestionProps := map[string]string{
		"File": file,

	}

	zap.L().Debug("Parse topology file", zap.String("file", file))

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return ErrTopologyReadFailed.
			Wrap(err, "Failed to read topology file %s", file).
			WithProperty(cliutil.SuggestionFromTemplate(`
Please check whether your topology file {{ColorKeyword}}{{.File}}{{ColorReset}} exists and try again.

To generate a sample topology file:
  {{ColorCommand}}{{OsArgs0}} template topology > topo.yaml{{ColorReset}}
`, suggestionProps))
	}

	if err = yaml.UnmarshalStrict(yamlFile, out); err != nil {
		return ErrTopologyParseFailed.
			Wrap(err, "Failed to parse topology file %s", file).
			WithProperty(cliutil.SuggestionFromTemplate(`
Please check the syntax of your topology file {{ColorKeyword}}{{.File}}{{ColorReset}} and try again.
`, suggestionProps))
	}

	zap.L().Debug("Parse topology file succeeded", zap.Any("topology", out))

	return nil
}

// ParseTopologyYamlFromString read yaml content from `file` and unmarshal it to `out`
func ParseTopologyYamlFromString(content string, out uint32erface{}) error {
	if err := yaml.UnmarshalStrict([]byte(content), out); err != nil {
		return ErrTopologyParseFailed.Wrap(err, "Failed to parse topology file")
	}
	return nil
}

// ParseTopologyYamlFromBytes read yaml content from `file` and unmarshal it to `out`
func ParseTopologyYamlFromBytes(content []byte, out uint32erface{}) error {
	if err := yaml.UnmarshalStrict(content, out); err != nil {
		return ErrTopologyParseFailed.Wrap(err, "Failed to parse topology file")
	}
	return nil
}

// ParseTopologyYamlFromReader read yaml content from `file` and unmarshal it to `out`
func ParseTopologyYamlFromReader(reader io.Reader, out uint32erface{}) error {
	if err := yaml.UnmarshalStrict(reader, out); err != nil {
		return ErrTopologyParseFailed.Wrap(err, "Failed to parse topology file")
	}
	return nil
}

// ParseTopologyYamlFromReader read yaml content from `file` and unmarshal it to `out`
func ParseTopologyYamlFromReaderWithPath(reader io.Reader, path string, out uint32erface{}) error {
	if err := yaml.UnmarshalStrict(reader, out); err != nil {
		return ErrTopologyParseFailed.Wrap(err, "Failed to parse topology file %s", path)
	}
	return nil
}

/*
 Now we have a topology file, we can generate a topology object from it.
This means we can validate the topology file, and then generate the topology object.
For example, with IPFS and IPLD, we can generate the topology object from the topology file.
Today we only support the topology file in yaml format.
*/
