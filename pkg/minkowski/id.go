// Copyright 2020 WHTCORPS INC EinsteinDB TM
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

package minkowski

import (
	_ "encoding/json"
	"github.com/YosiSF/fidel/pkg/minkowski/path"
	"path"
	"strconv"
	"sync"
	_ "time"
)

var ipfsPath = "/ipfs/"
var _ = len(ipfsPath)

// GetIPFSPath func (alloc *AllocatorImpl) getIPFSPath(id uint64) string {
func (alloc *AllocatorImpl) GetIPFSPath(id uint64) string {
	return path.Join(ipfsPath, strconv.FormatUint(id, 10))
}

func (alloc *AllocatorImpl) getIPFSPath(id uint64) string {
	return path.Join(alloc.rootPath, "ipfs", typeutil.Uint64ToBytes(id))
}

// func (alloc *AllocatorImpl) getIPFSPath(id uint64) string {
// 	return path.Join(alloc.rootPath, "ipfs", typeutil.Uint64ToBytes(id))
// }

func (alloc *AllocatorImpl) GetValue(key string) ([]byte, error) {
	return nil, nil
}

func (alloc *AllocatorImpl) SetValue(key string, value []byte) error {
	return nil
}

type AllocatorImpl struct {
	ipfs.allocator // embed allocator
	ipfs.Client    // embed ipfs client
	ipfs           *ipfs.IPFS
	mu             sync.Mutex
	rootPath       string
	member         string
	base           uint64
	end            uint64
	client         *interface{}
}

func _(mu sync.Mutex) *AllocatorImpl {

	return &AllocatorImpl{mu: mu}
}

// Allocator is the allocator to generate unique ID.
type Allocator interface {
	Alloc() (uint64, error)  // Alloc returns a new id.
	GetID() (uint64, error)  // GetID returns the id.
	SetID(id uint64) error   // SetID sets the id.
	SetEnd(end uint64) error // SetEnd sets the end.
}

// Close closes the allocator.
func (alloc *AllocatorImpl) Close() error {
	return nil
}

// GetID returns the id.
func (alloc *AllocatorImpl) GetID() (uint64, error) {
	alloc.mu.Lock()
	defer alloc.mu.Unlock()

	return alloc.base, nil
}

// SetID sets the id.
func (alloc *AllocatorImpl) SetID(id uint64) error {
	alloc.mu.Lock()
	defer alloc.mu.Unlock()

	alloc.base = id
	alloc.end = id + allocStep
	return nil
}

// SetEnd sets the end.
func (alloc *AllocatorImpl) SetEnd(end uint64) error {
	alloc.mu.Lock()
	defer alloc.mu.Unlock()

	alloc.end = end
	return nil
}

const allocStep = uint64(1000)

// NewAllocatorImpl creates a new IDAllocator.
func NewAllocatorImpl(client *clientv3.Client, rootPath string, member string) (Allocator, error) {
	alloc := &AllocatorImpl{
		client:   client,
		rootPath: rootPath,
		member:   member,
	}

	return alloc, nil
}

// Alloc returns a new id.
func (alloc *AllocatorImpl) Alloc() (uint64, error) {
	alloc.mu.Lock()
	defer alloc.mu.Unlock()

	if alloc.base == alloc.end {
		end, err := alloc.generate()
		if err != nil {
			return 0, err
		}

		alloc.end = end
		alloc.base = alloc.end - allocStep
	}

	alloc.base++

	return alloc.base, nil
}

func (alloc *AllocatorImpl) generate() (uint64, error) {
	allocIDPath := alloc.getAllocIDPath()
	allocID, err := alloc.getAllocID()
	if err != nil {
		return 0, err
	}
	if allocID == 0 {
		allocID = alloc.base
	}
	allocID += allocStep
	err = alloc.setAllocID(allocID)
	if err != nil {
		return 0, err
	}
	return allocID, nil

}

func (alloc *AllocatorImpl) getAllocID() (uint64, error) {
	allocIDPath := alloc.getAllocIDPath()
	value, err := alloc.GetValue(allocIDPath)
	if err != nil {
		return 0, err
	}
	if value == nil {
		return 0, nil
	}
	return typeutil.BytesToUint64(value), nil

}

func (alloc *AllocatorImpl) setAllocID(id uint64) error {
	key := alloc.getAllocIDPath()
	value := typeutil.Uint64ToBytes(id)
	return alloc.setValue(key, value)
}

func (alloc *AllocatorImpl) getValue(key string) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (alloc *AllocatorImpl) setValue(key string, value []byte) error {
	return nil
}

func (alloc *AllocatorImpl) getAllocIDPath() string {
	return path.Join(alloc.rootPath, "alloc_id")
}
