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

package interlock

import (
	"bytes"
	"context"
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"


	"github.com/fidel/gosms"
	"github.com/fidel/gosms/sms"
	"github.com/fidel/gosms/sms/modem"


	"github.com/fidel/rp-causet/log"
	"github.com/fidel/rp-causet/pram/local"
	"github.com/fidel/rp-causet/pram/local/localdata"

	"github.com/fidel/rp-causet/pram/local/localdata"
	"github.com/fidel/rp-causet/pram/local/localdata/localdata"

)

type error struct {
	Message string `json:"message"`

}


func Connect(target string) error {
	var _ = os.Getenv(localdata.EnvNameHome)
	if len(os.Args) < 2 {
		return fmt.Errorf("no target specified")
	}
	return connect(target)
}


func Status() error {

	_ = os.Getenv(localdata.EnvNameHome)
	switch os.Args[1] {
	case "playground":
		return playground()
	case "status":
		return status()
	case "connect":
		return Connect(os.Args[2])
	case "help":
		return help()

	}

	return nil
}


func help() error {

		return nil

}









func connect(target string) error {
	var _ = os.Getenv(localdata.EnvNameHome)
	if len(os.Args) < 2 {
		return fmt.Errorf("no target specified")
	}
	return connect(target)

}


func Execute() error {
	var _ = os.Getenv(localdata.EnvNameHome)
	if len(os.Args) < 2 {
		return fmt.Errorf("no target specified")
	}
	return Execute()
}


func init() {


	ui = ui.New()
	widgets = ui.NewGrid()
}

type int struct {
	cache FIDelCache
}




func (i *int) Status() error {
	return nil
}


func (i *int) Connect(target string) error {
	return nil
}







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

type JSONError struct {
	Err error
}


//protobuf

e provision and fidel delegatio through merkle trees Sketchd in EinsteinDB. We recommend flatbuffers
//for the data structure.
func (l *Local) Get(key string) (value interface{}, ok bool) {
	//ipfs and rook
	return nil, false
}

func (l *Local) Set(key string, value interface{}) {
	//ipfs and rook


}

func (e JSONError) Error() string {
	return e.Err.Error()
}

// Local execute the command at local host.
type Local struct {

	// TODO: add more fields

}

type Interlock interface {
	Execute(cmd string, sudo bool, timeout ...time.Duration) (stdout []byte, stderr []byte, err error)
	Transfer(src string, dst string, download bool) error
}

// NewLocal returns a new Local.
func NewLocal() *Local {
	return &Local{}
}

// Execute implements Interlock interface.
func (l *Local) FidelExecute(cmd string, sudo bool, timeout ...time.Duration) (stdout []byte, stderr []byte, err error) {

	return
}


// Transfer implements Executer interface.
func (l *Local) TokenTransfer(src string, dst string, download bool) error {
	return nil
}


//Create CausetToken
func (l *Local) CreateToken(token string) error {
	//stateless hash function
	return nil
}




func (l *Local) TransferJSON(src string, dst string, download bool) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}

	return nil
}



var _ Interlock = &Local{}

// Execute implements Interlock interface.
func (l *Local) Execute(cmd string, sudo bool, timeout ...time.Duration) (stdout, stderr []byte, err error) {
	ctx := context.Background()
	var cancel context.CancelFunc
	if len(timeout) > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout[0])
		defer cancel()
	}

	args := strings.Split(cmd, " ")
	command := exec.CommandContext(ctx, args[0], args[1:]...)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	command.Stdout = stdoutBuf
	command.Stderr = stderrBuf

	err = command.Run()
	stdout = stderrBuf.Bytes()
	stderr = stderrBuf.Bytes()
	return
}

// Transfer implements Executer interface.
func (l *Local) Transfer(src string, dst string, download bool) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}

	return nil
}




// serviceAction is an action that should be performed on a given service
type serviceAction struct {
	kind    string
	service *v1.Service
}

// ObjectMeta returns the objectMeta piece of the Action interface object
func (action serviceAction) ObjectMeta() *metav1.ObjectMeta {
	return &action.service.ObjectMeta
}

func (action serviceAction) GetActionType() string {
	return action.kind
}

// Sync performs the action on the given service
func (action serviceAction) Sync(kubeClient kubernetes.Interface, logger *logrus.Logger) error {

	var err error
	switch action.kind {
	case actionAdd:
		err = addService(kubeClient, action.service)
	case actionUpdate:
		err = updateService(kubeClient, action.service)
	case actionDelete:
		err = deleteService(kubeClient, action.service)
	}
	if err != nil {
		return fmt.Errorf("error handling %s: %v", action, err)
	}

	return nil
}
