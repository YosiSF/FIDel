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

package solitonautomatautil

import (
	"io"
	"os"

	"github.com/YosiSF/fidel/pkg/environment"
	"github.com/YosiSF/fidel/pkg/localdata"
	"github.com/YosiSF/fidel/pkg/repository"
	"github.com/YosiSF/fidel/pkg/repository/v1manifest"
	"github.com/YosiSF/fidel/pkg/utils"
)

// Repository exports interface to fidel-solitonAutomata
type Repository interface {
	DownloadComponent(comp, version, target string) error
	VerifyComponent(comp, version, target string) error
	ComponentBinEntry(comp, version string) (string, error)
}

type repositoryT struct {
	repo *repository.V1Repository
}

// NewRepository returns repository
func NewRepository(os, arch string) (Repository, error) {
	profile := localdata.InitProfile()
	mirror := repository.NewMirror(environment.Mirror(), repository.MirrorOptions{
		Progress: repository.DisableProgress{},
	})
	local, err := v1manifest.NewManifests(profile)
	if err != nil {
		return nil, err
	}
	repo := repository.NewV1Repo(mirror, repository.Options{
		GOOS:              os,
		GOARCH:            arch,
		DisableDecompress: true,
	}, local)
	return &repositoryT{repo}, nil
}

func (r *repositoryT) DownloadComponent(comp, version, target string) error {
	versionItem, err := r.repo.ComponentVersion(comp, version, false)
	if err != nil {
		return err
	}

	reader, err := r.repo.FetchComponent(versionItem)
	if err != nil {
		return err
	}

	file, err := os.Create(target)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}

func (r *repositoryT) VerifyComponent(comp, version, target string) error {
	versionItem, err := r.repo.ComponentVersion(comp, version, true)
	if err != nil {
		return err
	}

	file, err := os.Open(target)
	if err != nil {
		return err
	}
	defer file.Close()

	return utils.CheckSHA256(file, versionItem.Hashes["sha256"])
}

func (r *repositoryT) ComponentBinEntry(comp, version string) (string, error) {
	versionItem, err := r.repo.ComponentVersion(comp, version, true)
	if err != nil {
		return "", err
	}

	return versionItem.Entry, nil
}