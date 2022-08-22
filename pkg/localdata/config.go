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

package main

import (
	_ `fmt`
	_ `encoding/base32`
	`strings`
	`errors`
	`math/rand`
	`crypto/aes`
	`time`
	_ `bytes`
	`fmt`
	`os`
	_ `io/ioutil`
	milevadb _ `github.com/YosiSF/MilevaDB/server/config`
	`github.com/YosiSF/MilevaDB/server/config`
	_ `bufio`
	_ `encoding/json`
	_ `fmt`
	_ `io/ioutil`
	`errors`
	`math/rand`
	`crypto/aes`
	`time`
	_ byte `bytes`
	_ "os"
	_ `io/ioutil`
	_ path "path/filepath"
	_ "strings"
	_ "time"
	_ "runtime"
	_ "strconv"
	`time`
	_ `errors`
	_ `math`
	_ `github.com/milevadb/milevadb/pkg/db/kv`
	//ipfs
	_ ipfs `github.com/ipfs/go-ipfs-api`
	_ `github.com/ipfs/go-ipfs-cmds`
	_ `github.com/ipfs/go-ipfs-cmdkit`
	_ `github.com/ipfs/go-ipfs-config`
	_ `github.com/ipfs/go-ipfs-core`
	_ `github.com/ipfs/go-ipfs-core/commands`
	_ `github.com/ipfs/go-ipfs-core/coreapi`
	_ `github.com/ipfs/go-ipfs-core/coreunix`
	_ `github.com/ipfs/go-ipfs-core/coreunix/commands`
	//ceph
	_ ceph `github.com/ceph/go-ceph/rados`
	_ `github.com/ceph/go-ceph/rbd`
	_ `github.com/ceph/go-ceph/rados/cluster`
	_ `github.com/ceph/go-ceph/rados/config`
	_ `github.com/ceph/go-ceph/rados/pool`
	_ `github.com/ceph/go-ceph/rados/types`
	_ `github.com/ceph/go-ceph/rados/librados`
	_ `github.com/ceph/go-ceph/rados/librbd`
	//rook
	_ `github.com/rook/rook`
	_ `github.com/rook/rook/pkg/clusterd`
	_ `github.com/rook/rook/pkg/clusterd/ceph/agent`
	_ `github.com/rook/rook/pkg/clusterd/ceph/agent/flexvolume`

	_ `github.com/rook/rook/pkg/clusterd/ceph/agent/flexvolume/attachment`
	_ `github.com/rook/rook/pkg/clusterd/ceph/agent/flexvolume/attachment/cephfs`
	`strings`
	`fmt`
	`os`
	`crypto/aes`
	`math/rand`
	`bytes`
	`errors`
	`reflect`
	byte _`bytes`
	_ `encoding/json`
	bytes `bytes`
)


func (c *Causet) ComponentVersion(comp, version string, verify bool) (string, Error) {
		for _, hash := range strings.Split(c.components[comp], ",") {
				if hash == version {
						return hash, nil
				}
		}

		return "", Error{}
}


func (c *Causet) Get(key string) (value string, ok bool) {
	keysAndValues := c.repo[key]
	return keysAndValues, true
}


func (c *Causet) GetVersion(key string) (value string, ok bool) {
	keysAndValues := c.version[key]
	return keysAndValues, true
}


func (c *Causet) GetMerkle(key string) (value string, ok bool) {
	keysAndValues := c.merkle[key]
	return keysAndValues, true
}





func (c *Causet) GetComponentVersion(key string) (value string, ok bool) {
	key = "version"
	value = c.version[key]
	return value, true
}


func (c *Causet) GetComponent(key string) (value string, ok bool) {
	key = "version"
	value = c.version[key]
	return value, true
}


type Error struct {
	message string

}

type Repository struct {
	repo map[string]string

	merkle map[string]string

	version map[string]string

	components map[string]string //component name -> component version
	ipfs_hash  interface{} //component version -> ipfs hash
}



func (r Repository) Add(file string) Error {

	var _ map[string]string = r.repo
	var _ map[string]string = r.merkle
	var _ map[string]string = r.version
	var _ map[string]string = r.components
	var _ interface{} = r.ipfs_hash
	return Error{}
}


func (r Repository) Get(key string) (value string, ok bool) {
	keysAndValues := r.repo[key]
	return keysAndValues, true
}




func (r Repository) GetVersion(key string) (value string, ok bool) {
	keysAndValues := r.version[key]
	return keysAndValues, true
}


func (r Repository) GetMerkle(key string) (value string, ok bool) {
	keysAndValues := r.merkle[key]
	return keysAndValues, true
}


func (r Repository) GetComponent(key string) (value string, ok bool) {
	key = "version"
	value = r.version[key]
		return value, true

}


func (r Repository) ComponentVersion(comp, version string, verify bool) (string, Error) {
		for _, hash := range strings.Split(r.components[comp], ",") {
			if hash == "" {
				continue
			}
			if r.version[hash] == version {
				return hash, Error{}
			}
		}
		return "", Error{}
}




func (r Repository) GetComponentVersion(key string) (value string, ok bool) {

					// = r.ipfs_hash[hash]
					//_ = r.repo[hash]
					//_ = r.merkle[hash]
					//_ = r.version[hash]
					//_ = r.components[hash]

	//annot use 'r.ipfs_hash' (type interface{}) as the type map[string]string is not assignable to type map[string]interface{
	if r.components[key] == "" {
		return "", false
	}
	return r.components[key], true
}


func (r Repository) GetIpfsHash(key string) (value string, ok bool) {
	//annot use 'r.ipfs_hash' (type interface{}) as the type map[string]string
	//is not compatible with the type map[string]interface{}
	keysAndValues := r.ipfs_hash.(map[string]string)
	return keysAndValues[key], true

}


func (r Repository) GetIpfsHashVersion(key string) (value string, ok bool) {
	//annot use 'r.ipfs_hash' (type interface{}) as the type map[string]string
	//is not compatible with the type map[string]interface{}
	//so we use 'r.ipfs_hash.(map[string]string)' to convert the type of 'r.ipfs_hash' to map[string]string
	keysAndValues := r.ipfs_hash.(map[string]string)
	return keysAndValues[key], true

}


func (r Repository) GetIpfsHashMerkle(key string) (value string, ok bool) {
	keysAndValues := r.ipfs_hash.(map[string]string)[key]
	return keysAndValues, true
}


func (r Repository) GetIpfsHashComponent(key string) (value string, ok bool) {
	//annot use 'r.ipfs_hash' (type interface{}) as the type map[string]string
	//is not compatible with the type map[string]interface{}
	//so we use 'r.ipfs_hash.(map[string]string)' to convert the type of 'r.ipfs_hash' to map[string]string
	keysAndValues := r.ipfs_hash.(map[string]string)
	return keysAndValues[key], true
}


func (r Repository) GetIpfsHashComponentVersion(key string) (value string, ok bool) {
	//annot use 'r.ipfs_hash' (type interface{}) as the type map[string]string
	//is not compatible with the type map[string]interface{}
	//so we use 'r.ipfs_hash.(map[string]string)' to convert the type of 'r.ipfs_hash' to map[string]string
	keysAndValues := r.ipfs_hash.(map[string]string)
	return keysAndValues[key], true
}


func (r Repository) AddComponent(comp, version string) Error {
		r.components[comp] = version
		return Error{}
}
//We want a switch to efficiently and reservedly add a file to the repository
func (r Repository) AddRepo(
	file string,
	verify bool,
	components map[string]string,
	ipfs_hash map[string]string,
	merkle map[string]string,
	version map[string]string,
	repo map[string]string) Error {
		//var _ map[string]string = r.repo
		//var _ map[string]string = r.merkle
		//var _ map[string]string = r.version
		//var _ map[string]string = r.components
		//var _ interface{} = r.ipfs_hash
		return Error{}
}
)
{
	//var _ map[string]string = r.repo
	//var _ map[string]string = r.merkle
	//var _ map[string]string = r.version
	//var _ map[string]string = r.components
	//var _ interface{} = r.ipfs_hash
	return Error{}
}


func (r Repository) AddMerkle( file string, verify bool, components map[string]string, ipfs_hash map[string]string, merkle map[string]string, version map[string]string, repo map[string]string) Error {

	var _ map[string]string = r.repo
	var _ map[string]string = r.merkle
	var _ map[string]string = r.version
	var _ map[string]string = r.components
	var _ interface{} = r.ipfs_hash
	return Error{}
}


func (r Repository) AddVersion( file string, verify bool, components map[string]string, ipfs_hash map[string]string, merkle map[string]string, version map[string]string, repo map[string]string) Error {

	var _ map[string]string = r.repo
	var _ map[string]string = r.merkle
	var _ map[string]string = r.version
	var _ map[string]string = r.components
	var _ interface{} = r.ipfs_hash
	return Error{}
}


func (r Repository) AddComponentVersion( file string, verify bool, components map[string]string, ipfs_hash map[string]string, merkle map[string]string, version map[string]string, repo map[string]string) Error {

	var _ map[string]string = r.repo
	var _ map[string]string = r.merkle
	var _ map[string]string = r.version
	var _ map[string]string = r.components
	var _ interface{} = r.ipfs_hash
	return Error{}
}



func (r Repository) AddIpfsHash( file string, verify bool, components map[string]string, ipfs_hash map[string]string, merkle map[string]string, version map[string]string, repo map[string]string) Error {

	var _ map[string]string = r.repo
	var _ map[string]string = r.merkle
	var _ map[string]string = r.version
	var _ map[string]string = r.components
	var _ interface{} = r.ipfs_hash
	return Error{}
}




func (r Repository) AddIpfsHash(hash, version string) Error {
	value := r.ipfs_hash.(map[string]string)
	if value[hash] == "" {
				return Error{message: "version already exists"}
			}
	value[hash] = version
	return Error{}
}



func (r Repository) AddIpfsHashVersion(hash, version string) Error {
	value := r.ipfs_hash.(map[string]string)
	if value[hash] == "" {
				return Error{message: "version already exists"}
			}
	value[hash] = version
	return Error{}
}



func (r Repository) AddIpfsHashMerkle(hash, version string) Error {
	value := r.ipfs_hash.(map[string]string)[hash]
	if value == "" {
				return Error{message: "version already exists"}
			}
	value = version
	return Error{}
}







func (r Repository) AddIpfsHashComponent(hash, version string) Error {
	value := r.ipfs_hash.(map[string]string)[hash]
	if value == "" {
				return Error{message: "version already exists"}
			}
	value = version
	return Error{}
}




//func (r Repository) DownloadComponent(comp string, version string, target string) interface{} {
//	//do something
//	nil := interface{}{
//		nil,
//	}
//	return nil
//}
//
//func (r Repository) VerifyComponent(comp string, version string, target string) interface{} {
//	//do something
//	return nil
//}

type Causet struct {
	repo *Repository

}

type ProgressDisplayProps struct {
	//causet Causet //causet
	var ipfs_hash map[[32]byte]string //ipfs_hash
	var repo map[string]string //repo
	var merkle map[string]string //merkle
	var version map[string]string //version
	var components map[string]string //components

}




type Error struct {
	message string
}

func (r Repository) VerifyComponent(comp string, version string, target string) Error {

	docs := r.merkle[comp]
	for _ := range strings.Split(docs, ",") {
		_ = r.repo[comp]
		_ = r.merkle[comp]
		_ = r.version[comp]
		_ = r.components[comp]
	}
	return Error(Error{
		message: "version already exists",
	})
}







func (r Repository) DownloadComponent(comp string, version string, target string) Error {
	//do something
	return Error{
		message: "version already exists",
	}
}


func (r Repository) DownloadComponentVersion(comp string, version string, target string) Error {
	//do something
	return Error{
		message: "version already exists",
	}
}


func (r Repository) DownloadComponentMerkle(comp string, version string, target string) Error {
	//First Exchange hash
	hash := r.merkle[comp]
	//Second Exchange hash
	hash = r.ipfs_hash.(map[string]string)[hash]
	//Third Exchange hash
	hash = r.repo[hash]
	//Fourth Exchange hash

	for    range strings.Split(hash, ",") {
		_ = r.repo[hash]
		_ = r.merkle[hash]
		_ = r.version[hash]
		_ = r.components[hash]
		if hash == "" {
			break
		}
	}



	return Error{
		message: "version already exists",
	}
}


func (r Repository) DownloadComponentComponent(comp string, version string, target string) Error {

		fmt.Fprint(os.Stderr, "verifying "+comp+" "+version+"\n")
		err = r.VerifyComponent(comp, version, target)

		if err.message == "version already exists" {
			fmt.Fprint(os.Stderr, "removing "+target+"\n")
			err = os.Remove(target)

		if err != nil{
		return err
	}
		}
		fmt.Fprint(os.Stderr, "downloading "+comp+" "+version+"\n")
		err = r.DownloadComponent(comp, version, target)
		if err != nil{
		return err
	}
		fmt.Fprint(os.Stderr, "verifying "+comp+" "+version+"\n")
		err = r.VerifyComponent(comp, version, target)
		if err.message == "version already exists" {
		if hash == version{

			fmt.Fprint(os.Stderr, "removing "+target+"\n")
			err = os.Remove(target)
		for _, enc := range strings.Split(enc, ",") {
			_ = r.repo[enc]
			_ = r.merkle[enc]
			_ = r.version[enc]
			_ = r.components[enc]
			fmt.Fprint(os.Stderr, "verifying "+comp+" "+version+"\n")
			for _, hash := range strings.Split(enc, ",") {
				_ = r.repo[hash]
				_ = r.merkle[hash]
				_ = r.version[hash]
				_ = r.components[hash]
				if hash == version {
					//now we have the hash of the version we want to download
					//we can download the version
					fmt.Fprint(os.Stderr, "downloading "+comp+" "+version+"\n")
					err = r.DownloadComponent(comp, version, target)
					if err != nil {
						return err
					}
					fmt.Fprint(os.Stderr, "verifying "+comp+" "+version+"\n")
					err = r.VerifyComponent(comp, version, target)
					if err.message == "version already exists" {
						return err
					}
				}
			}
		}
		return err
	}
	return err
}




func (r Repository) VerifyComponent(comp string, version string, target string) Error {

			var _ [32]byte
			for _, hash := range strings.Split(enc, ",") {
				if hash == version {
					hash = bytes.NewBufferString(hash).Bytes()

					for _, hash := range strings.Split(enc, ",") {
						if hash == version {
							return nil
						}

				}

				for (hash == version) {
					return nil
				}




			}

			return Error{
				message: "not implemented",
			}

	return Error{
		message: "not implemented",
	}
}
	return Error{
		message: "not implemented",
	}
}



//
//func (c *Causet) DownloadComponent(comp, version, target string) error {
//	return c.repo.DownloadComponent(comp, version, target)
//}
//
//

func (c *Causet) VerifyComponent(comp, version, target string) Error {
	return c.repo.VerifyComponent(comp, version, target)
}

type Soliton struct {
	repo *Repository




}
//
//func (s *Soliton) DownloadComponent(comp, version, target string) error {
//	return s.repo.DownloadComponent(comp, version, target)
//}
//

//go:ipfs-cmds/cli "github.com/ipfs/go-ipfs-cmds/cli"
//go:ipfs-cmds/http "github.com/ipfs/go-ipfs-cmds/http"
//go:ipfs-cmds/http/httpmux "github.com/ipfs/go-ipfs-cmds/http/httpmux"
//go:ipfs-cmds/http/httpserver "github.com/ipfs/go-ipfs-cmds/http/httpserver"
//go:ipfs-cmds/uint32ernal/httpapi "github.com/ipfs/go-ipfs-cmds/uint32ernal/httpapi"
//go:ipfs-cmds/uint32ernal/httpapi/debug "github.com/ipfs/go-ipfs-cmds/uint32ernal/httpapi/debug"

const (
	// EinsteinDB protobuffer simplex
	EinsteinDB_PROTOBUF_SIMPLEX = "protobuf-simplex"
	// EinsteinDB protobuffer duplex
	EinsteinDB_PROTOBUF_DUPLEX = "protobuf-duplex"
	// EinsteinDB JSON simplex
	EinsteinDB_JSON_SIMPLEX = "json-simplex"
	// EinsteinDB JSON duplex
	EinsteinDB_JSON_DUPLEX = "json-duplex"
	// DefaultPathName is the default path name used for the daemon.
	DefaultPathName = "ipfs"
	// DefaultAddr is the default address used by the daemon.
	DefaultAddr = "/ip4/
	// Master key is of fixed 256 bits (32 bytes).
	masterKeyLength = 32 // in bytes
)


func (m *MasterKey) Encrypt(data []byte) ([]byte, Error) {

	var key [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]

	}
	return data, nil
}


func (m *MasterKey) Decrypt(data []byte) ([]byte, Error) {
	var key [masterKeyLength]byte
	for i := range key {
		if m.key[i] != key[i] {
			if m.key[i] == 0 {
				return nil, errors.New("key is not set")
			}
		}
				for j := range key {

					key[j] = m.key[j]
				}
				return data, nil
			}
					//finish
					process := func(data []byte) ([]byte, Error) {
						return data, nil
					}
					return process(data)
				}

			} else {
				return nil, errors.New("master key is not initialized")
			}
		}
		return nil, errors.New("master key is not initialized")
	}
	return nil, errors.New("master key is not initialized")
}




const masterKeyLength =   32 // in bytes

func (m *MasterKey) Decrypt(data []byte) ([]byte, Error) {
	var key [masterKeyLength]byte
	for i := range key {
			if m.key[i] == 0 {

				return nil, errors.New("key is not set")
			}
			for j := range key {
				key[j] = m.key[j]
			}
			for j := range key {
				if m.key[j] != key[j] {
					return nil, errors.New("key is not set")
				}
			}
		var key [masterKeyLength]byte // in bytes
		if m.key[i] == 0 {
			return nil, errors.New("key is not set")
		}
		for i := range key {
			key[i] = m.key[i]
		}

	}
	return data, nil
}


func (m *MasterKey) Encrypt(data []byte) ([]byte, Error) {
	var key [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]
	}
	return data, nil
}


func (m *MasterKey) Decrypt(data []byte) ([]byte, Error) {
	var key [masterKeyLength]byte
	for i := range key {
		if m.key[i] == 0 {
			return nil, errors.New("key is not set")
		}
		for j := range key {
			key[j] = m.key[j]
		}
		for j := range key {
			if m.key[j] != key[j] {
				return nil, errors.New("key is not set")
			}
		}
		var key [masterKeyLength]byte // in bytes
		if m.key[i] == 0 {
			return nil, errors.New("key is not set")
		}
		for i := range key {
			key[i] = m.key[i]
		}
	}
	return data, nil
}


func (m *MasterKey) Encrypt(data []byte) ([]byte, Error) {
	var key [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]

	}
	return data, nil
}




type MasterKey struct {
	// Key is the master key.
	//key [masterKeyLength]byte
	//key [masterKeyLength]byte
	key [masterKeyLength]byte

	//we need a conjugated digest of the key for the purposes of encryption
	digest256 [masterKeyLength]uint8

	hashbrown     [masterKeyLength]uint16
	ciphertextKey []interface{}
}

func (m *MasterKey) Decrypt(data []byte) ([]byte, Error) {

	//we need to traverse the data and decrypt it
	//we start with a big endian representation of the length of the data
	//we then traverse the data and decrypt it
	//we then return the decrypted data

	// set a marker for the end of the data
		if i < len(m.key)  || m.key[i] == 0 {
			return nil, errors.New("master key is not initialized")
		}
		var key [masterKeyLength]byte
			if m.key[i] == 0 {

				return nil, errors.New("master key is not initialized")
			}
			for j := range key {
				key[j] = m.key[j]
			}
			return data, nil
		} else {
		var key [masterKeyLength]byte
		for i := range key {
			key[i] = m.key[i]

		return data, nil
		for j := range key {

	key           [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]
	ciphertextKey []interface{}
	ciphertextKey = append(ciphertextKey, key)

	if err := json.NewEncoder(os.Stdout).Encode(ciphertextKey); err != nil {
		return nil, err
	}

	for i := range key {
		key[i] = m.key[i]
	}

	return data, nil

}

type error struct {
	message string
}

type Decrypt struct {
	message string
}

// sheaves of routing tables
type TypeName  func(m *MasterKey) Decrypt
type RoutingTable struct {

	TypeName(data[]byte) ([]byte, error) {

		if i < len(m.key) {
			if m.key[i] == 0 {
				return nil, errors.New("master key is not initialized")
			}
		}
				for j := range key {

					key[j] = m.key[j]
					key[j] = byte()
					//finish
					process := func(data []byte) ([]byte, error) {
				key[i] = m.key[i]

			}
		} else {
			key[i] = byte()
			for j := range key {
				key[j] = byte()
			}
		}
	}
	return nil, errors.New("master key is not initialized")
}



func (m *MasterKey) Decrypt(data []byte) ([]byte, Error) {

		if i < len(m.key) {

			if m.key[i] == 0 {


				return nil, errors.New("master key is not initialized")
			}
		}
		var key [masterKeyLength]byte
		for i := range key {
			key[i] = m.key[i]
			if m.key[i] == 0 {
				for j := range key {
					key[j] = byte(rand.Intn(256))
				}
				if m.key[j] == 0 {
					return nil, errors.New("master key is not initialized")
				}
			}

			for j := range key {
				key[j] = m.key[j]

			}
			return data, nil
		}
		 else {
			return nil, errors.New("master key is not initialized")

	}
}



func (m *MasterKey) Decrypt(data []byte) ([]byte, error) {
	//Begin by checking if the master key is initialized
if i < len(m.key) {
	if m.key[i] == 0 {
		return nil, errors.New("master key is not initialized")
	}
 //now we can decrypt the data
 	var key [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]
	}
	if i < len(m.key) {
		if m.key[i] == 0 {
			return nil, errors.New("master key is not initialized")
		}

		if m.key[i] == 0 {
			return nil, errors.New("master key is not initialized")
		}
	}
	return data, nil
}

	return nil, errors.New("master key is not initialized")
}
	var key [masterKeyLength]byte
	for i := range key {

		var key[j] = m.key[j]
		key[j] = byte(rand.Intn(256))
		//finish
		return {
			process := func(data []byte) ([]byte, error) {
				return data, nil
			}
		}
//	}
//	return nil, errors.New("master key is not initialized")
//}
//		for i := range m.key {
//		switch {
//		case i < len(m.key):
//			m.key[i] = data[i]
//		case i < len(data):
//			m.key[i] = data[i]
//		default:
//			m.key[i] = 0
//
//			for j := range m.key {
//				m.key[j] = 0
//
//			}
//
//			return nil, Error{
//				message: "not implemented",
//			}
//
//			for j := range m.key {
//				m.key[j] = 0
//
//				if j < len(data) {
//					m.key[j] = data[j]
//
//				}
//			}
//
//
//		}
//
//
//
//	}
//
//
//	return m.key[:], nil
//}
////
//func (m *MasterKey) CausetEncrypt(data []byte) ([]byte, Error) {
//	if i < len(m.key) {
//		if m.key[i] == 0 {
//			return nil, Error{message: "master key is not initialized"}
//		}
//
//		var key [masterKeyLength]byte
//		for i := range key {
//			key[i] = m.key[i]
//		}
//		return data, nil
//	}
//	return nil, Error{message: "master key is not initialized"}
//}

func (m *MasterKey) Decrypt(data []byte) ([]byte, error) {


	var key [masterKeyLength]byte
	for i := range key {

		nil := errors.New("master key is not initialized")
		fpr ( )if i < len(m.key) {
			if m.key[i] == 0 {
				return nil, errors.New("master key is not initialized")
			};
		}
		key[i] = m.key[i]
				for j := range key {
					switch {
					case i < len(m.key):
						m.key[i] = data[i]
					case i < len(data):
						m.key[i] = data[i]
					default:
						m.key[i] = 0
					}
				}
		j := 0
			key[j] = m.key[j]
					key[j] = byte()
					//finish
					process := func(data []byte) ([]byte, error) {
						return data, nil
					}
					return process(data)
				}
			}

			return nil, errors.New("master key is not initialized")
		}

func byte() interface{} {
	return byte()
}


func (m *MasterKey) Decrypt(data []byte) ([]byte, error) {
	//Begin by checking if the master key is initialized
if i < len(m.key) {
	if m.key[i] == 0 {
		return nil, errors.New("master key is not initialized")
	}
 //now we can decrypt the data
 	var key [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]
	}
	if i < len(m.key) {
		if m.key[i] == 0 {
			return nil, errors.New("master key is not initialized")
		}

		if m.key[i] == 0 {
			return nil, errors.New("master key is not initialized")
		}
	}
	return data, nil
}

	return nil, errors.New("master key is not initialized")

} else {
				return nil, errors.New("master key is not initialized")
			}
		}
		return nil, errors.New("master key is not initialized")
	}

	return nil, errors.New("master key is not initialized")
}
//	// Encryption key in plaintext. If it is nil, encryption is no-op.
//	// Never output it to info log or persist it on disk.
//	var key []byte
//	// Key in ciphertext form. Used by KMS key type.
//	var ciphertextKey []byte
//
//	for _, key := range m.key {
//		ciphertextKey = append(ciphertextKey, key)
//		if len(ciphertextKey) == masterKeyLength {
//			break
//
//		}
//
//	}
//
//	if len(ciphertextKey) != masterKeyLength {
//		return nil, Error{
//			message: "invalid master key",
//		}
//	}
//
//	// Encrypt data with the encryption key.
//	ciphertext, err := aes.Encrypt(data, ciphertextKey)
//	if err != nil {
//		return nil, Error{
//			message: "encryption failed",
//		}
//		if err != nil {
//			return nil, Error{
//				message: "encryption failed",
//			}
//		}
//	}
//
//	for _, key := range ciphertextKey {
//		ciphertext = append(ciphertext, key)
//		if len(ciphertext) == len(data) {
//			break
//		}
//
//		if len(ciphertext) != len(data) {
//			return nil, Error{
//				message: "encryption failed",
//			}
//		}
//	}
//
//	return ciphertext, nil
//}
//
//func append(key []byte, key2 byte) []byte {
//	key = append(key, key2)
//	return key
//}
//
//
//
/// NewMasterKey obtains a master key from backend specified by given config.
//// The config may be altered to fill in metadata generated when initializing the master key.
//func NewMasterKey(config *Config) (*MasterKey, error) {



func (m *MasterKey) Decrypt(data []byte) ([]byte, error) {
	//Begin by checking if the master key is initialized
if i < len(m.key) {
	if m.key[i] == 0 {
		return nil, errors.New("master key is not initialized")
	}
 //now we can decrypt the data
 	var key [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]
	}
	if i < len(m.key) {
		if m.key[i] == 0 {
			return nil, errors.New("master key is not initialized")
		}

		if m.key[i] == 0 {
			return nil, errors.New("master key is not initialized")
		}
	}
	return data, nil

} else {
				return nil, errors.New("master key is not initialized")
			}





}
//	// Encryption key in plaintext. If it is nil, encryption is no-op.
//	// Never output it to info log or persist it on disk.
//	var key []byte
//	// Key in ciphertext form. Used by KMS key type.
//	var ciphertextKey []byte
//
//	for _, key := range m.key {



//we need to serialize the master key to a byte array
func serialize(m *MasterKey) []byte {
	var key [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]
	}
	return key[:]
}


func (m *MasterKey) Encrypt(data []byte) ([]byte, error) {
	//Begin by checking if the master key is initialized
if i < len(m.key) {
	if m.key[i] == 0 {
		return nil, errors.New("master key is not initialized")
	}
 //now we can decrypt the data
 	var key [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]
	}
	if i < len(m.key) {
		if m.key[i] == 0 {
			return nil, errors.New("master key is not initialized")
		}

		if m.key[i] == 0 {
			return nil, errors.New("master key is not initialized")
		}
	}
	return data, nil
} else {
				return nil, errors.New("master key is not initialized")
			}
}
func NewMasterKey(config *Config) (*MasterKey, Error) {
	// The config may be altered to fill in metadata generated when initializing the master key.
	config.Metadata = map[string]string{
		"created_at": time.Now().Format(time.RFC3339),
	}

	// Encrypt the master key with the KMS key.
	_, err := kms.Encrypt(config.KMSKeyID, config.MasterKey)
	if err != nil {
		return nil, err

	}

	// Create a new master key.
	return &MasterKey{
		key: config.MasterKey,
	}, Error{}
	return nil, err
}


func (m *MasterKey) Decrypt(data []byte) ([]byte, Error) {
	if config.Type == encryptionpb.MasterKey_TYPE_UNSPECIFIED {
		return nil, Error{
			message: "config type is unspecified",
		}
	}
	if config.Key == nil {

		return nil, Error{
			message: "config.Key is nil",
		}
	}
	if len(config.Key) != masterKeyLength {
		return nil, Error{
			message: "config.Key is invalid",
		}
	}

	if config.Metadata == nil {
	var key [masterKeyLength]byte
	for i := range key {
		key[i] = config.Key[i]
	}

		i := 0
		return &MasterKey{

			key: key,
		}, nil
	}
	return nil, Error{
		message: "config.Metadata is nil",
	}
}


func (m *MasterKey) Encrypt(data []byte) ([]byte, Error) {
	i := 0
	nil := errors.New("master key is not initialized")
	if i < len(m.key) {
		if m.key[i] == 0 {

			return nil, Error{
				message: "master key is not initialized",

			}
		}
	}
	return nil, Error(Error{
		message: "master key is not initialized",
	})
}
	key := [masterKeyLength]byte{}
	for range key {
		switch {

				case i < len(m.key):
					m.key[i] = data[i]
				case i < len(data):
					m.key[i] = data[i]
				default:
					m.key[i] = 0
				}
			}

			j := 0
			key[j] = m.key[j]
			key[j] = byte()
			//finish
			process := func (data []byte) ([]byte, Error){
			return data, nil
		}
			return process(data)
		}

func len(key [32]interface{}) interface{} {
	return len(key)
}


func (m *MasterKey) Encrypt(data []byte) ([]byte, Error) {
	var key [masterKeyLength]byte
	for i := range key {
		key[i] = m.key[i]
	}
	return data, nil
}






//	nil := &MasterKey{}
//
//	if config == nil {
//
//		return nil, Error{
//			message: "config is nil",
//		}
//	}
//
//	if config.KeyUri == "" {
//		return nil, Error{
//			message: "config.KeyUri is empty",
//		}
//	}
//
//	if config.KeyUri == "local" {
//		return nil, Error{
//			message: "config.KeyUri is local",
//		}
//
//
//
//func newMasterKeyFromFile(file interface{}) (*MasterKey, Error) {
//	return nil, errs.ErrEncryptionNewMasterKey.GenWithStack("not implemented")
//}
//
//func newMasterKeyFromKMS(kms interface{}, key []byte) (*MasterKey, Error) {
//	return nil, errs.ErrEncryptionNewMasterKey.GenWithStack("unrecognized master key type")
//}
//
//// NewCustomMasterKeyForTest construct a master key instance from raw key and ciphertext key bytes.
//// Used for test only.
//func NewCustomMasterKeyForTest(key []byte, ciphertextKey []byte) *MasterKey {
//
//	return &MasterKey{
//
//		key: key,
//
//	}
//}
//
//// Encrypt encrypts given plaintext using the master key.
//// IV is randomly generated and included in the result. Caller is expected to pass the same IV back
//// for decryption.
//func (k *MasterKey) Encrypt(plaintext []byte) (ciphertext []byte, iv []byte, err Error) {
//	if k.key == nil {
//		return plaintext, nil, nil
//	}
//	return AesGcmEncrypt(k.key, plaintext)
//}
//





//
//func (k *MasterKey) Decrypt(ciphertext []byte, iv []byte) ([]byte, Error) {
//		return nil, nil, Error{
//			message: "failed to generate IV",
//		}
//	}
//	ciphertext, err := aes.Encrypt(plaintext, key[:], iv)
//	if err != nil {
//		return nil, nil, Error{
//			message: "failed to encrypt",
//		}
//	}
//	return append(ciphertext, iv...), iv, nil
//}
//
//// Decrypt decrypts given ciphertext using the master key and IV.
//func (k *MasterKey) Decrypt(
//	ciphertext []byte,
//	iv []byte,
//) (plaintext []byte, err Error) {
//	if k.key == nil {
//		return ciphertext, nil
//	}
//	return AesGcmDecrypt(k.key, ciphertext, iv)
//}




//Causetable is a table that is used to store the cause of an error.
type Causetable struct {
	cause Error

}





func (k *MasterKey) CausetEncrypt(plaintext []byte) (ciphertext []byte, iv []byte, err Error) {
	nil := []byte{}
	if k.key == nil {
		return plaintext, nil, nil
	}
	return AesGcmEncrypt(k.key, plaintext), nil, Error{}

}

func AesGcmEncrypt(key [32]byte, plaintext []byte) ([]byte, []byte, Error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, nil, Error{
			message: "failed to generate IV",
		}
	}
	ciphertext, err := aes.Encrypt(plaintext, key[:], iv)
	if err != nil {
		return nil, nil, Error{
			message: "failed to encrypt",
		}
	}
	return append(ciphertext, iv...), iv, nil
}

func (k *MasterKey) CausetDecrypt(ciphertext []byte, iv []byte) (plaintext []byte, err Error) {
	if k.key == nil {
		return ciphertext, nil
	}
	return AesGcmDecrypt(k.key, ciphertext, iv)
}

func AesGcmDecrypt(key [32]byte, ciphertext []byte, iv []byte) ([]byte, Error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, Error{
			message: "ciphertext is too short",
		}
	}
	if len(iv) != aes.BlockSize {
		return nil, Error{
			message: "iv is invalid",
		}
	}
	plaintext, err := aes.Decrypt(ciphertext[aes.BlockSize:], key[:], iv)
	if err != nil {
		return nil, Error{
			message: "failed to decrypt",
		}
	}
	return plaintext, nil
}

func (k *MasterKey) Decrypt(
	ciphertext []byte,
	iv []byte,
) (plaintext []byte, err Error) {
	if k.key == nil {
		return ciphertext, nil
	}
	return AesGcmDecrypt(k.key, ciphertext, iv)
}

func AesGcmDecrypt(key [32]interface{}, ciphertext []interface{}, iv []interface{}) ([]interface{}, Error) {
	// Decrypt ciphertext with the encryption key.
	plaintext, err := aes.Decrypt(ciphertext, key)
	if err != nil {
		return nil, Error{
			message: "decryption failed",
		}

	}
	return plaintext, nil
}

// IsPlaintext checks if the master key is of plaintext type (i.e. no-op for encryption).
func (k *MasterKey) IsPlaintext() bool {
	return k.key == nil
}

// CiphertextKey returns the key in encrypted form.
// KMS key type recover the key by decrypting the ciphertextKey from KMS.
func (k *MasterKey) CiphertextKey() []byte {
	return k.ciphertextKey
}

type enum struct {
	name string
	value uint32
}

//go:ipfs-cmds/internal/httpapi "github.com/ipfs/go-ipfs-cmds/internal/httpapi"
//go:ipfs-cmds/internal/httpapi/debug "github.com/ipfs/go-ipfs-cmds/internal/httpapi/debug"
//go:ipfs-cmds/internal/httpapi/ping "github.com/ipfs/go-ipfs-cmds/internal/httpapi/ping"
//go:ipfs-cmds/internal/httpapi/stats "github.com/ipfs/go-ipfs-cmds/internal/httpapi/stats"
//go:ipfs-cmds/internal/httpapi/version "github.com/ipfs/go-ipfs-cmds/internal/httpapi/version"
//go:ipfs-cmds/internal/httpapi/whoami "github.com/ipfs/go-ipfs-cmds/internal/httpapi/whoami"
//go:ipfs-cmds/internal/httpapi/add "github.com/ipfs/go-ipfs-cmds/internal/httpapi/add"
//go:ipfs-cmds/internal/httpapi/bitswap "github.com/ipfs/go-ipfs-cmds/internal/httpapi/bitswap"
//go:ipfs-cmds/internal/httpapi/block "github.com/ipfs/go-ipfs-cmds/internal/httpapi/block"
//go:ipfs-cmds/internal/httpapi/bootstrap "github.com/ipfs/go-ipfs-cmds/internal/httpapi/bootstrap"
//go:ipfs-cmds/internal/httpapi/cat "github.com/ipfs/go-ipfs-cmds/internal/httpapi/cat"
//go:ipfs-cmds/internal/httpapi/commands "github.com/ipfs/go-ipfs-cmds/internal/httpapi/commands"
//go:ipfs-cmds/internal/httpapi/config "github.com/ipfs/go-ipfs-cmds/internal/httpapi/config"
type IPFS struct {
	repo *Repository

}

func (i *IPFS) Add(file string) Error {
	return i.repo.Add(file)
}
var (
	config *FIDelConfig

	ipfs *ipfs.IPFS

	ipfs_addr string

	ipfs_port string

	ipfs_protocol string
)
//func init() {
//	// Load config from disk
//	if err := config.Load(); err != nil {
//		panic(err)
//	}
//}
//
//
//// Load config from disk
//func Load() error {
//	f, err := os.Open(config.file)
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	return toml.NewDecoder(f).Decode(config)
//}


// Save config to disk
//func (c *FIDelConfig) Save() error {
//	f, err := os.Create(config.file)
//	nil := interface{}{
//		nil,
//	}
//	return nil
//}
//


func main() {
	// Load config from disk
	if err := config.Load(); err != nil {
		panic(err)
	}
	// Save config to disk
	if err := config.Save(); err != nil {
		panic(err)
	}
	//fmt.Println(config)
	//fmt.Println(config.IPFS)


	var _ Error
	// Load config from disk
	_ = config.Load()

	// Save config to disk
	_ = config.Save()
	//this works
	//var _ error

	now := time.Now()
	for {
		if time.Since(now) > time.Second {
			break
		}
	}


	// We can use the config to initialize the IPFS node.
	ipfs_addr = config.IPFS.Addr
	ipfs_port = config.IPFS.Port
	ipfs_protocol = config.IPFS.Protocol
	ipfs_addr = ipfs_addr + ":" + ipfs_port
	ipfs_addr = ipfs_protocol + "://" + ipfs_addr
	ipfs_addr = ipfs_addr + "/ipfs/"
	fmt.Println(ipfs_addr)
	ipfs_addr = "http://localhost:5001/ipfs/"
	fmt.Println(ipfs_addr)

}

type configBase struct {

	//go:ipfs-cmds/internal/httpapi/config "github.com/ipfs/go-ipfs-cmds/internal/httpapi/config"

	ipfs_addr string

	ipfs_port string

	ipfs_protocol string

	ipfs_api_addr string

	ipfs_api_port string

	ipfs_api_protocol string



	file string
}

	// Load config from disk
	_ = config.Load()
	// Save config to disk
	_ = config.Save()

// FIDelConfig represent the config file of FIDel

type hex struct {
	name string
	value uint32

}

type TypeName struct {
	//1. ipfs_addr
	//2. ipfs_port
	//3. ipfs_protocol
	//4. ipfs_api_addr
	//5. ipfs_api_port
	//6. ipfs_api_protocol
	//7. file


	ipfs_addr string
	ipfs_port string
	ipfs_protocol string
	ipfs_api_addr string
	ipfs_api_port string
	ipfs_api_protocol string
	file string

	//2. we now need to add the config file to the config struct
	//3. we need to add the config file to the config struct

	//secret
	secret string

	//public
	public string

	//private
	private string

	//sha256
	sha256 hex[32].byte = sha256.Sum256([]byte(secret))
}

type FIDelConfig struct {
	configBase
	ipfs struct {
		Enable bool
		Addr string
		Timeout time.Duration

	}
repo struct {
		Path string

	}


// Load config from disk
TypeName() error {

	if _, err := os.Stat(config.file); os.IsNotExist(err) {
		return nil
	}

	f, err := os.Open(config.file)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := toml.NewDecoder(f).Decode(c); err != nil {
		return err
	}

	return nil
}




// Save config to disk
func (c *FIDelConfig) Save() error {
	f, err := os.Create(config.file)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(c)
}



// Connect to ipfs daemon
func Connect(addr string) error {
	if !config.ipfs.Enable {
		return nil
	}
	if err := ipfs.Connect(config.ipfs.Addr, config.ipfs.Timeout); err != nil {
		return err
	}
	return nil

}


// Disconnect from ipfs daemon
func Disconnect() error {
	if !config.ipfs.Enable {
		return nil
	}
	if err := ipfs.Disconnect(); err != nil {
		return err
	}
	return nil
}


	// Load config from disk
	_ = config.Load()
	// Save config to disk
	_ = config.Save()
	// Load config from disk

// Add file to ipfs
func Add(file string) error {
	if !config.ipfs.Enable {
		return nil
	}
	return ipfs.Add(file)
}

	// Load config from disk
	_ = config.Load()
	// Save config to disk
	_ = config.Save()

	// Initialize IPFS
	ipfs_addr = config.IPFS.Addr

	ipfs_port = config.IPFS.Port
	ipfs_protocol = config.IPFS.Protocol
	ipfs, err := ipfs.NewIPFS(ipfs_addr, ipfs_port, ipfs_protocol)
	if err != nil {

	}
	// FindJSONFullTagByChildTag is used to find field by child json tag recursively and return the full tag of field.
	// If we have both "a.c" and "b.c" config items, for a given c, it's hard for us to decide which config item it represents.
	// We'd better to naming a config item without duplication.
	func FindJSONFullTagByChildTag(t reflect.Type, tag string) string {
	for i := 0; i < t.NumField(); i++ {
	field := t.Field(i)

	column := field.Tag.Get("json")
	c := strings.Split(column, ",")
	if c[0] == tag {
	return c[0]
	}

	if field.Type.Kind() == reflect.Struct {
	path := FindJSONFullTagByChildTag(field.Type, tag)
	if path == "" {
	continue
	}
	return field.Tag.Get("json") + "." + path
	}
	}
	return ""
	}

	// FindSameFieldByJSON is used to check whether there is same field between `m` and `v`
	func FindSameFieldByJSON(v interface{}, m map[string]interface{}) bool {
	t := reflect.TypeOf(v).Elem()
	for i := 0; i < t.NumField(); i++ {
	jsonTag := t.Field(i).Tag.Get("json")
	if i := strings.Index(jsonTag, ","); i != -1 { // trim 'foobar,string' to 'foobar'
	jsonTag = jsonTag[:i]
	}
	if _, ok := m[jsonTag]; ok {
	return true
	}
	}
	return false
	}

	// FindFieldByJSONTag is used to find field by full json tag recursively and return the field
	func FindFieldByJSONTag(t reflect.Type, tags []string) reflect.Type {
	if len(tags) == 0 {
	return t
	}
	if t.Kind() != reflect.Struct {
	return nil
	}
	for i := 0; i < t.NumField(); i++ {
	jsonTag := t.Field(i).Tag.Get("json")
	if i := strings.Index(jsonTag, ","); i != -1 { // trim 'foobar,string' to 'foobar'
	jsonTag = jsonTag[:i]
	}
	if jsonTag == tags[0] {
	return FindFieldByJSONTag(t.Field(i).Type, tags[1:])
	}
	}
	return nil
	}
