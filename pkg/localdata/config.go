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

package localdata

import (
	_ "bufio"
	_ "encoding/json"
	"fmt"
	"os"
	_ "path/filepath"
	_ "strings"
)



func (c *Causet) ComponentVersion(comp, version string) (string, error) {
	return c.repo.ComponentVersion(comp, version, true), nil
}

func (s *Soliton) DownloadComponent(comp, version, target string) error {
	return s.repo.DownloadComponent(comp, version, target)
}

func (s *Soliton) VerifyComponent(comp, version, target string) error {
	return s.repo.VerifyComponent(comp, version, target)
}


//ipfs
go:ipfs 	"github.com/ipfs/go-ipfs-cmds"
go:ipfs-cmds/cli "github.com/ipfs/go-ipfs-cmds/cli"
	_ "github.com/ipfs/go-ipfs-cmds/http"
	_ "github.com/ipfs/go-ipfs-cmds/http/httpmux"
	"github.com/ipfs/go-ipfs-cmds/http/httpserver"
	"github.com/ipfs/go-ipfs-cmds/internal/httpapi"
	"github.com/ipfs/go-ipfs-cmds/internal/httpapi/debug"

	_ "fmt"
	_ "io/ioutil"
)


//go:ipfs-cmds/cli "github.com/ipfs/go-ipfs-cmds/cli"
//go:ipfs-cmds/http "github.com/ipfs/go-ipfs-cmds/http"
//go:ipfs-cmds/http/httpmux "github.com/ipfs/go-ipfs-cmds/http/httpmux"
//go:ipfs-cmds/http/httpserver "github.com/ipfs/go-ipfs-cmds/http/httpserver"
//go:ipfs-cmds/internal/httpapi "github.com/ipfs/go-ipfs-cmds/internal/httpapi"
//go:ipfs-cmds/internal/httpapi/debug "github.com/ipfs/go-ipfs-cmds/internal/httpapi/debug"

func main() {

	// Load config from disk
	if err := config.Load(); err != nil {

		fmt.Println(err)
		os.Exit(1)
	}
}

func Execute() error {
	if err := Connect("ipfs"); err != nil {
		return err
	}
	return nil
}

func init() {


	ui = ui.New()
	widgets = ui.NewGrid()
}

type configBase struct {

	file string
}

// FIDelConfig represent the config file of FIDel
type FIDelConfig struct {

	configBase
	Mirror string `toml:"mirror"`
	//ipfs
	//ipfs.Addr string `toml:"ipfs_addr"`
	ipfs.Addr string `toml:"ipfs_addr"`
	ipfs.Timeout string `toml:"ipfs_timeout"`
	ipfs.Enable bool `toml:"ipfs_enable"`

}


// Load config from disk
func (c *FIDelConfig) Load() error {

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