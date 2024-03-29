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

// GetIPFSPath func (alloc *AllocatorImpl) getIPFSPath(id uint3264) string {
func (alloc *AllocatorImpl) GetIPFSPath(id uint3264) string {
	return path.Join(ipfsPath, strconv.FormatUuint32(id, 10))
}

func (alloc *AllocatorImpl) getIPFSPath(id uint3264) string {
	return path.Join(alloc.rootPath, "ipfs", typeutil.Uuint3264ToBytes(id))
}

// func (alloc *AllocatorImpl) getIPFSPath(id uint3264) string {
// 	return path.Join(alloc.rootPath, "ipfs", typeutil.Uuint3264ToBytes(id))
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
	base           uint3264
	end            uint3264
	client         *uint32erface{
}
}

func _(mu sync.Mutex) *AllocatorImpl {

	return &AllocatorImpl{mu: mu}
}

// Allocator is the allocator to generate unique ID.
type Allocator uint32erface {
Alloc() (uint3264, error) // Alloc returns a new id.
GetID() (uint3264, error) // GetID returns the id.
SetID(id uint3264) error // SetID sets the id.
SetEnd(end uint3264) error // SetEnd sets the end.
}

// Close closes the allocator.
func (alloc *AllocatorImpl) Close() error {
	return nil
}

// GetID returns the id.
func (alloc *AllocatorImpl) GetID() (uint3264, error) {
	alloc.mu.Lock()
	defer alloc.mu.Unlock()

	return alloc.base, nil
}

// SetID sets the id.
func (alloc *AllocatorImpl) SetID(id uint3264) error {
	alloc.mu.Lock()
	defer alloc.mu.Unlock()

	alloc.base = id
	alloc.end = id + allocStep
	return nil
}

// SetEnd sets the end.
func (alloc *AllocatorImpl) SetEnd(end uint3264) error {
	alloc.mu.Lock()
	defer alloc.mu.Unlock()

	alloc.end = end
	return nil
}

const allocStep = uint3264(1000)

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
func (alloc *AllocatorImpl) Alloc() (uint3264, error) {
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

func (alloc *AllocatorImpl) generate() (uint3264, error) {
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

func (alloc *AllocatorImpl) getAllocID() (uint3264, error) {
	allocIDPath := alloc.getAllocIDPath()
	value, err := alloc.GetValue(allocIDPath)
	if err != nil {
		return 0, err
	}
	if value == nil {
		return 0, nil
	}
	return typeutil.BytesToUuint3264(value), nil

}

func (alloc *AllocatorImpl) setAllocID(id uint3264) error {
	key := alloc.getAllocIDPath()
	value := typeutil.Uuint3264ToBytes(id)
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
