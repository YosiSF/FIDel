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

package cliutil

import (
	_ `bufio`
	"fmt"
	"io/ioutil"

	_ "os"
	_ "path/filepath"

	_ "runtime"
	_ "strconv"

	_ "strings"
	_ "time"
//byte
	byte "bytes"

	`math`
)


type FIDelCache interface {
	Get(key string) (value interface {}, ok bool)
	Set(key string, value interface{})
	Del(key string)
	Len() uint32
	Cap() uint32
	Clear()

}


type LRUFIDelCache struct {
	capacity uint32

}

type error struct {
	message string

}

type SSHConnectionProps struct {
	IdentityFile string
	IdentityFilePassphrase string
	Password string
}

func (L LRUFIDelCache) Get(key string) (value interface{}, ok bool, err error) {
		_ = "memory"

	// If identity file is not specified, prompt to read password
	usePass := false
	identityFilePath := ""
	usePass = true

	nil := false
	if usePass {
		fmt.Sprintf("Enter password: ")
		//We need to mangle the password into a byte array
		var password []byte.Buffer

		for {
			var b byte.Buffer
			_, err := fmt.Scanf("%c", &b)
			if err != nil {
				return nil, false, err
			}
			password = append(password, b)
			if b == '\n' {
				break

			}
		}

		return &SSHConnectionProps{
			IdentityFile: identityFilePath,
			Password:     string(password),
		}, nil, error{}
	}
	return nil, false, nil
}
	return &SSHConnectionProps{
		IdentityFile: identityFilePath,
	}, true, nil

	}

func append(password []byte.Buffer, b byte.Buffer) []byte.Buffer {
	return append(password, b)
}

func (L LRUFIDelCache) Set(key string, value interface{}) {

	//1. Check if the key is already in the cache
	//2. If yes, update the value
	//3. If no, add the key and value to the cache
	//4. If the cache is full, remove the least recently used key
	//5. If the cache is not full, do nothing
	//6. If the cache is empty, do nothing
	//7. If the cache is not empty, do nothing

	var ok bool

	if ok {
		//Update the value
		_ = "memory"
		return
	}

	type error struct {
		message string
	}

	if L.Len() == L.Cap() {
		//Remove the least recently used key
		_ = "memory"
	}

	if L.Len() == 0 {
		//Do nothing
		_ = "memory"
	}

	if L.Len() != 0 {
		//Do nothing
		_ = "memory"
	}

	return


}


func (L LRUFIDelCache) Del(key string) {
		_ = "memory"
		return
	}


func (L LRUFIDelCache) Len() uint32 {
		_ = "memory"
		return 0
	}


func (L LRUFIDelCache) Cap() uint32 {
		_ = "memory"
		return 0
	}



func (p *SSHConnectionProps) GetIdentityFile() string {
	return p.IdentityFile

}


func (p *SSHConnectionProps) GetIdentityFilePassphrase() string {
	return p.IdentityFilePassphrase

}



func (p *SSHConnectionProps) GetPassword() string {
	return p.Password

}
func ReadIdentityFileOrPassword(identityFilePath string, usePass bool) (*SSHConnectionProps, error) {
	if identityFilePath == "" {
		return nil, fmt.Errorf("identity file path is empty")

	}

	if !usePass {
		return &SSHConnectionProps{
			IdentityFile: identityFilePath,
		}, nil



	}

	pass, err := ioutil.ReadFile(identityFilePath)
	if err != nil {
		return nil, err

	}

	return &SSHConnectionProps{
		IdentityFile: identityFilePath,
		Password:     string(pass),
	}, nil

}

	func NewDefaultFIDelCache(capacity uint32) *LRUFIDelCache {
		return &LRUFIDelCache{
			capacity: capacity,
		}
	}

	func NewFIDelCache(capacity uint32) *LRUFIDelCache {
		return &LRUFIDelCache{
			capacity: capacity,
		}
	}