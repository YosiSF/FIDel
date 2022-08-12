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

//// shutdownTracerProvider adds a shutdown method for tracer providers.
////
//// Note that this doesn't directly use the provided TracerProvider interface
//// to avoid build breaking go-ipfs if new methods are added to it.


func (c *Causet) ComponentVersion(comp, version string) (string, error) {
	return c.repo.ComponentVersion(comp, version, true)
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


type shutDownTracerProvider struct {
	tp TracerProvider

}

func (s *shutDownTracerProvider) Close() error {
	return s.tp.Close()
}


func (s *shutDownTracerProvider) GetTracer(name string) (Tracer, error) {
	return s.tp.GetTracer(name)
}


func (s *shutDownTracerProvider) GetAllTracers() ([]Tracer, error) {
	return s.tp.GetAllTracers()
}


func (s *shutDownTracerProvider) GetTracerNames() ([]string, error) {
	return s.tp.GetTracerNames()
}


func (s *shutDownTracerProvider) GetTracerConfig(name string) (interface{}, error) {
	return s.tp.GetTracerConfig(name)
}


func (s *shutDownTracerProvider) SetTracerConfig(name string, config interface{}) error {
	return s.tp.SetTracerConfig(name, config)
}

	if name == "jaeger" {
		return s.tp.GetTracer(name)
	}

	for _, t := range s.tp.GetTracers() {
		if t.Name() == name {
			return t, nil
		}
	}

	while _, err := s.tp.GetTracer(name);
	b := err != nil{
		return nil, err
	}

	//lock
	return nil, nil

}


func (s *shutDownTracerProvider) GetTracers() []Tracer {
	return s.tp.GetTracers()
}