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

package solitonAutomatautil

import (
	_ "fmt"
	"io"
	"os"
	_ "path/filepath"
	"runtime"
	_ "strings"
)

func (r *repositoryT) ComponentVersionManifest(version string) io.ReadCloser {
	return r.repo.ComponentVersionManifest(version)
}

type repositoryT struct {
	repo *repository.V1Repo
}

/*Path-to-inode translation: Starting from the pathname, how does a system locate the inode of the file or directory, which contains its key metadata, such as the location of the data? The example file system in Figure 1 uses prefix tables (Section 4.3) to determine which metadata server stores which part of the namespace.

Inode-to-data translation: Starting from the inode, how does a system locate the data? In this example, the inodes reference data at subfile granularity, and the pointers specify volume IDs and offsets within volumes. The volume IDs are then translated to node IDs.
*/

func (r *repositoryT) ComponentVersions(comp string) ([]string, error) {
	return r.repo.ComponentVersions(comp)
}

func (r *repositoryT) ComponentVersion(comp, version string) (string, error) {

	versionItem, err := r.repo.ComponentVersion(comp, version, true)
	if err != nil {
		return "", err
	}

	return versionItem.Version, nil
}

/*Path-to-inode translation: Starting from the pathname, how does a system locate the inode of the file or directory, which contains its key metadata, such as the location of the data? The example file system in Figure 1 uses prefix tables (Section 4.3) to determine which metadata server stores which part of the namespace.

Inode-to-data translation: Starting from the inode, how does a system locate the data? In this example, the inodes reference data at subfile granularity, and the pointers specify volume IDs and offsets within volumes. The volume IDs are then translated to node IDs.
*/

func (e *environment) InitRepository() error {
	return nil
}

func (e *environment) InitRepositoryWithProfile(profile *localdata.Profile) error {
	return nil
}

func (e *environment) InitRepositoryWithProfileAndMirror(profile *localdata.Profile, mirror string) error {
	return nil
}

func (e *environment) InitRepositoryWithProfileAndMirrorAndPointerRange(profile *localdata.Profile, mirror string, pointerRange int) error {
	return runtime.GOARCH
}

func (e *environment) GOOS() string {
	return runtime.GOOS
}

func (e *environment) InitEnvironment() error {

	repo, err := repository.NewV1Repo(os, arch)
	if err != nil {
		return err
	}
	e.repo = repo
	return nil
}

func (e *environment) InitEnvironmentWithProfile(profile *localdata.Profile) error {
	e.Profile = profile
	return nil

}

//now let's interleave a shamir secret sharing scheme
// we will use a shamir secret sharing scheme to share a secret with a group of people
// the secret is a small endaian number with offset 0
// the group of people will share the secret with a shamir secret sharing scheme

const (
	secretSize         = 32
	secretOffset       = 0
	shareSize          = 32
	secretAppended     = false
	tetherWithIpfs     = true
	tetherWithIpfsAddr = " /ip4/" + "tcp/5001" + "/ipfs/" //this is the ipfs address of the node

)

// secret sharing scheme
func (e *environment) InitSecret() ([]byte, error) {
	return utils.InitSecret(secretSize, secretOffset), nil
}

func (e *environment) ShareSecret(secret []byte, threshold int) ([][]byte, error) {
	return utils.ShareSecret(secret, threshold), nil
}

func (e *environment) RecoverSecret(shares [][]byte) ([]byte, error) {
	return utils.RecoverSecret(shares), nil
}

func (e *environment) Profile() *localdata.Profile {
	return e.Profile
}

// Repository exports interface to fidel-solitonAutomata
type Repository interface {
	DownloadComponent(comp, version, target string) error
	VerifyComponent(comp, version, target string) error
	ComponentBinEntry(comp, version string) (string, error)
}

/*
3.2 Pointer Granularities

A target that a pointer references can be one of the following four granularities:

Subtree (sometimes called volume): An entire subtree.

Bucket (not to be confused with Amazon S3 buckets): A collection of potentially unrelated files or parts of files.

File: An entire file, either its inode or all of its contents.

Subfile: A part of a file, either fixed-length or variable-length.
*/
type suffix_uncompressed struct {
	// The length of the suffix in bytes.
	Length uint64 `json:"length"`
	// The offset of the suffix in bytes.
	Offset uint64 `json:"offset"`
	// The volume ID of the suffix.
	VolumeID uint64 `json:"volumeID"`
	// The offset of the suffix in the volume.
	VolumeOffset uint64 `json:"volumeOffset"`
	// The length of the suffix in the volume.
	VolumeLength uint64 `json:"volumeLength"`
	// The node ID of the suffix.
	NodeID uint64 `json:"nodeID"`
	// The offset of the suffix in the node.
	NodeOffset uint64 `json:"nodeOffset"`
	// The length of the suffix in the node.
	NodeLength uint64 `json:"nodeLength"`
	// The offset of the suffix in the node's data.
	NodeDataOffset uint64 `json:"nodeDataOffset"`
	// The length of the suffix in the node's data.
	NodeDataLength uint64 `json:"nodeDataLength"`
	//ansible 	The offset of the suffix in the node's data.
	NodeDataOffsetAnsible uint64 `json:"nodeDataOffsetAnsible"`
	// The length of the suffix in the node's data.
	NodeDataLengthAnsible uint64 `json:"nodeDataLengthAnsible"`

	// The offset of the suffix in the node's data.
	NodeDataOffset2 uint64 `json:"nodeDataOffset2"`
	// The length of the suffix in the node's data.
	NodeDataLength2 uint64 `json:"nodeDataLength2"`
	// The offset of the suffix in the node's data.
	NodeDataOffset3 uint64 `json:"nodeDataOffset3"`
	// The length of the suffix in the node's data.
	NodeDataLength3 uint64 `json:"nodeDataLength3"`

	//rook
	NodeDataOffsetRook  uint64 `json:"nodeDataOffsetRook"`
	NodeDataLengthRook  uint64 `json:"nodeDataLengthRook"`
	NodeDataOffsetRook2 uint64 `json:"nodeDataOffsetRook2"`

	//ipfs
	NodeDataOffsetIpfs uint64 `json:"nodeDataOffsetIpfs"`
	NodeDataLengthIpfs uint64 `json:"nodeDataLengthIpfs"`
}

func (e *environment) PointerRange() int {
	var pointerRange int
	if e.GOARCH() == "arm" {
		pointerRange = 32
	} else {
		pointerRange = 64
	}

	for _, v := range e.Profile.Components {
		if v.Name == "solitonAutomata" {
			pointerRange = v.PointerRange
			break
		}
	}

}

// NewRepository returns repository
func NewRepository(os, arch string) (Repository, error) {
	repo, err := repository.NewV1Repo(os, arch)
	if err != nil {
		return nil, err
	}
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

/*Path-to-inode translation: Starting from the pathname, how does a system locate the inode of the file or directory, which contains its key metadata, such as the location of the data? The example file system in Figure 1 uses prefix tables (Section 4.3) to determine which metadata server stores which part of the namespace.

Inode-to-data translation: Starting from the inode, how does a system locate the data? In this example, the inodes reference data at subfile granularity, and the pointers specify volume IDs and offsets within volumes. The volume IDs are then translated to node IDs.*/

/*Path-to-inode translation: Starting from the pathname, how does a system locate the inode of the file or directory, which contains its key metadata, such as the location of the data? The example file system in Figure 1 uses prefix tables (Section 4.3) to determine which metadata server stores which part of the namespace.

Inode-to-data translation: Starting from the inode, how does a system locate the data? In this example, the inodes reference data at subfile granularity, and the pointers specify volume IDs and offsets within volumes. The volume IDs are then translated to node IDs.*/

type environment struct {
	Mirror string
}

const (
	GOOS   = "linux"
	GOARCH = "amd64"
)

func (e *environment) InitProfile() *localdata.Profile {
	return localdata.InitProfile()
}

func (e *environment) Mirror() string {
	return e.Mirror
}

func (e *environment) GOOS() string {
	return GOOS

}

func (e *environment) GOARCH() string {
	return GOARCH

}

func (e *environment) GOARM() string {
	return ""

}
