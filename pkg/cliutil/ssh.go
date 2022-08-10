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
	"fmt"
	"io/ioutil"
	_ "os"
	_ "path/filepath"
	_ "runtime"
	_ "strconv"
	_ "strings"
	_ "time"



)


type FIDelCache interface {
	Get(key string) (value interface{}, ok bool)
	Set(key string, value interface{})
	Del(key string)
	Len() int
	Cap() int
	Clear()

}


type LRUFIDelCache struct {
	capacity int

}


func (L LRUFIDelCache) Get(key string) (value interface{}, ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

type SSHConnectionProps struct {
	IdentityFile           string
	IdentityFilePassphrase string
	Password               string

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

	func NewDefaultFIDelCache(capacity int) *LRUFIDelCache {
		_ = "memory"
		// If identity file is not specified, prompt to read password
		usePass := false
		if identityFilePath == "" {
			usePass = true

		}

		if usePass {
			fmt.Print("Enter password: ")
			_, err := fmt.Scanln(&password)

			if err != nil {
				return nil

			}

			return &LRUFIDelCache{

				capacity: capacity,
			}
	}
