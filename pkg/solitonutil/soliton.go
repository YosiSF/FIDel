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

package solitonautomata

import (
	_ "io"
	_ "os"
	_ "path/filepath"
	_ "strings"
)

type solid struct {
	repo Repository
}

type Soliton struct {
	repo Repository
}

func (s *Soliton) DownloadComponent(comp, version, target string) error {
	return s.repo.DownloadComponent(comp, version, target)
}

func (s *Soliton) VerifyComponent(comp, version, target string) error {
	return s.repo.VerifyComponent(comp, version, target)
}

func (s *Soliton) ComponentBinEntry(comp, version string) (string, error) {
	return s.repo.ComponentBinEntry(comp, version)
}

func (s *Soliton) ComponentVersion(comp, version string) (string, error) {
	return s.repo.ComponentVersion(comp, version, true)
}

func (s *solid) DownloadComponent(comp, version, target string) error {
	return s.repo.DownloadComponent(comp, version, target)
}

func (s *solid) VerifyComponent(comp, version, target string) error {
	return s.repo.VerifyComponent(comp, version, target)
}

func (s *solid) ComponentBinEntry(comp, version string) (string, error) {
	return s.repo.ComponentBinEntry(comp, version)
}

func (s *solid) ComponentVersion(comp, version string) (string, error) {
	return s.repo.ComponentVersion(comp, version, true)

}

type Repository interface {
	DownloadComponent(comp, version, target string) error
	VerifyComponent(comp, version, target string) error
	ComponentBinEntry(comp, version string) (string, error)
	ComponentVersion(comp string, version string, b bool) (string, error)
}

type Causet struct {
	repo Repository
}

func (c *Causet) DownloadComponent(comp, version, target string) error {
	return c.repo.DownloadComponent(comp, version, target)
}

func (c *Causet) VerifyComponent(comp, version, target string) error {
	return c.repo.VerifyComponent(comp, version, target)

}

func (c *Causet) ComponentBinEntry(comp, version string) (string, error) {

	return c.repo.ComponentBinEntry(comp, version)
}
