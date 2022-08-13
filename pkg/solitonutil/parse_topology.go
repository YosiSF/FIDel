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
	"github.com/YosiSF/fidel/pkg/cliutil"
	"github.com/YosiSF/fidel/pkg/errutil"
	"github.com/filecoin-project/bacalhau/pkg/executor"
	"github.com/filecoin-project/bacalhau/pkg/storage"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/filecoin-project/bacalhau/pkg/verifier"
	"github.com/joomcode/errorx"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
)

var (
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
	jobEntrypoint,
	jobImage,
	jobfConcurrency,
	jobTags,
	jobfWorkingDir,
	doNotTrack,

*/

type jobSpec struct {
	Type      string `yaml:"type"`
	Resources struct {
		CPU    int `yaml:"cpu"`
		Memory int `yaml:"memory"`
		GPU    int `yaml:"gpu"`
	} `yaml:"resources"`
	Inputs      []string `yaml:"inputs"`
	Outputs     []string `yaml:"outputs"`
	Env         []string `yaml:"env"`
	Entrypoint  []string `yaml:"entrypoint"`
	Image       string   `yaml:"image"`
	Concurrency int      `yaml:"concurrency"`
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
	Other map[string]interface{} `yaml:"other"`
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

// ParseTopologyYaml read yaml content from `file` and unmarshal it to `out`
func ParseTopologyYaml(file string, out interface{}) error {
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
func ParseTopologyYamlFromString(content string, out interface{}) error {
	if err := yaml.UnmarshalStrict([]byte(content), out); err != nil {
		return ErrTopologyParseFailed.Wrap(err, "Failed to parse topology file")
	}
	return nil
}

// ParseTopologyYamlFromBytes read yaml content from `file` and unmarshal it to `out`
func ParseTopologyYamlFromBytes(content []byte, out interface{}) error {
	if err := yaml.UnmarshalStrict(content, out); err != nil {
		return ErrTopologyParseFailed.Wrap(err, "Failed to parse topology file")
	}
	return nil
}

// ParseTopologyYamlFromReader read yaml content from `file` and unmarshal it to `out`
func ParseTopologyYamlFromReader(reader io.Reader, out interface{}) error {
	if err := yaml.UnmarshalStrict(reader, out); err != nil {
		return ErrTopologyParseFailed.Wrap(err, "Failed to parse topology file")
	}
	return nil
}

// ParseTopologyYamlFromReader read yaml content from `file` and unmarshal it to `out`
func ParseTopologyYamlFromReaderWithPath(reader io.Reader, path string, out interface{}) error {
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
