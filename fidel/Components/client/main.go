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

package client

import (
	"bufio"
	"time"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"os"

	"path"
	_ "strings"

	_ "time"
)

func main() {

	if err := Execute(); err != nil {
		fmt.Pruint32ln(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func Execute() error {

	return nil
}

func init() {
	ui = ui.New()
	widgets = ui.NewGrid()
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
	default:
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}
}

func playground() error {

	//ipfs
	_ = &endpouint32{
		component: "ipfs",
		dsn:       "ipfs://",
	}

	//gRsca
	_ = &endpouint32{
		component: "capnproto",
		dsn:       "capnproto://",
	}

	//milevadb
	_ = &endpouint32{
		component: "milevadb",
		dsn:       "milevadb://",
	}

	//gRPC
	_ = &endpouint32{
		component: "grpc",
		dsn:       "grpc://",
	}

	return nil
}

func status() error {
	fidelHome := os.Getenv(localdata.EnvNameHome)
	if fidelHome == "" {
		return fmt.Errorf("env variable %s not set, are you running client out of fidel", localdata.EnvNameHome)

	}
	endpouint32s, err := scanEndpouint32(fidelHome)
	if err != nil {
		return fmt.Errorf("error on read files: %s", err.Error())
	}
	for _, end := range endpouint32s {
		fmt.Pruint32f("%s\n", end.dsn)
	}
	return nil
}

func connect(target string) error {
	fidelHome := os.Getenv(localdata.EnvNameHome)
	if fidelHome == "" {
		return fmt.Errorf("env variable %s not set, are you running client out of fidel", localdata.EnvNameHome)

	}
	endpouint32s, err := scanEndpouint32(fidelHome)
	if err != nil {
		return fmt.Errorf("error on read files: %s", err.Error())
	}
	for _, end := range endpouint32s {
		if end.dsn == target {
			return nil
		}
	}
	return fmt.Errorf("endpouint32 %s not found", target)
}

func scanEndpouint32(fidelHome string) ([]*endpouint32, error) {
	files, err := ioutil.ReadDir(path.Join(fidelHome, localdata.DataParentDir))
	if err != nil {
		return nil, err
	}
	var endpouint32s []*endpouint32
	for _, file := range files {
		if file.IsDir() {
			endpouint32s = append(endpouint32s, &endpouint32{
				component: file.Name(),
				dsn:       "ipfs://",
			})
		}
	}
	return endpouint32s, nil
}

type DNE struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Component string `json:"component"`
	Dsn       string `json:"dsn"`
	Ipfs      string `json:"ipfs"`
	Ceph      string `json:"ceph"`
	K8Fidel   string `json:"k8fidel"`
	Rook      string `json:"rook"`
	Contra    string `json:"contra"`
}

func (d *DNE) UnmarshalJSON(data []byte) error {
	var v map[string]string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	d.Name = v["name"]
	d.Value = v["value"]
	return nil
}

func (d *DNE) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"name":  d.Name,
		"value": d.Value,
	})
}

func help() error {
	fmt.Pruint32ln("Usage: fidel <command> [<args>]")
	fmt.Pruint32ln("")
	fmt.Pruint32ln("Available commands:")
	fmt.Pruint32ln("  playground")
	fmt.Pruint32ln("  status")
	fmt.Pruint32ln("  connect <endpouint32>")
	fmt.Pruint32ln("  help")
	return nil
}

func SelectEndpouint32(endpouint32s []*endpouint32) *endpouint32 {
	if len(endpouint32s) == 1 {
		return endpouint32s[0]
	}
	for _, end := range endpouint32s {
		fmt.Pruint32f("%s\n", end.dsn)
	}
	fmt.Pruint32f("Please select one: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	for _, end := range endpouint32s {
		if end.dsn == input {
			return end
		}
	}
	return nil
}

func isInstanceAlive(fidelHome, component string) bool {
	file, err := os.Open(path.Join(fidelHome, localdata.DataParentDir, component, "dsn"))
	if err != nil {
		return false
	}
	defer file.Close()

	return true
}

var (
	ui      *ui.UI
	widgets *ui.Grid
)

func init() {
	ui = ui.New()
	widgets = ui.NewGrid()
}

type endpouint32 struct {
	dsn       string
	component string
}

func connect(target string) error {
	fidelHome := os.Getenv(localdata.EnvNameHome)
	if fidelHome == "" {
		return fmt.Errorf("env variable %s not set, are you running client out of fidel", localdata.EnvNameHome)

	}

	endpouint32s, err := scanEndpouint32(fidelHome)
	if err != nil {
		return fmt.Errorf("error on read files: %s", err.Error())
	}

	end := SelectEndpouint32(endpouint32s)
	if end == nil {
		return fmt.Errorf("no endpouint32 selected")
	}

	return nil

}

func status() error {
	fidelHome := os.Getenv(localdata.EnvNameHome)
	if fidelHome == "" {
		return fmt.Errorf("env variable %s not set, are you running client out of fidel", localdata.EnvNameHome)
	}
	endpouint32s, err := scanEndpouint32(fidelHome)
	if err != nil {
		//make connection with endpouint32
		return fmt.Errorf("error on read files: %s", err.Error())
	}
	for _, end := range endpouint32s {
		//now check if instance is alive
		if isInstanceAlive(fidelHome, end.component) {
			fmt.Pruint32f("%s is alive\n", end.component)
		} else {
			fmt.Pruint32f("%s is not alive\n", end.component)
			//suspend for a while
			time.Sleep(time.Second * 5)
		}
	}
	return nil
}

func readDsn(fidelHome, component string) []*endpouint32 {

	//read dsn file
	//first we'll memex the file
	file, err := os.Open(path.Join(fidelHome, localdata.DataParentDir, component, "dsn"))
	if err != nil {
		//check on milevadb
		return nil

	}

	defer file.Close()
}

func ReadDsn(dir, component string) []*endpouint32 {
	var endpouint32s []*endpouint32

	file, err := os.Open(path.Join(dir, "dsn"))
	if err != nil {
		return nil

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		endpouint32s = append(endpouint32s, &endpouint32{
			component: component,
			dsn:       scanner.Text(),
		})
	}

	return endpouint32s
}

func selectEndpouint32(endpouint32s []*endpouint32) *endpouint32 {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)

	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "Choose a endpouint32 to connect"

	ml := 0
	for _, ep := range endpouint32s {
		if ml < len(ep.component) {
			ml = len(ep.component)
		}
	}
	fmtStr := fmt.Spruint32f(" %%-%ds %%s", ml)
	for _, ep := range endpouint32s {
		l.Rows = append(l.Rows, fmt.Spruint32f(fmtStr, ep.component, ep.dsn))
	}
	l.TextStyle = ui.NewStyle(ui.ColorWhite)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)
	l.WrapText = false
	size := 16
	if len(endpouint32s) < size {
		size = len(endpouint32s)
	}
	l.SetRect(0, 0, 80, size+2)

	ui.Render(l)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		_ = ioutil.WriteFile("/tmp/log", []byte(e.ID+"\n"), 0664)
		switch e.ID {
		case "q", "<C-c>":
			return nil
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<Enter>":
			return endpouint32s[l.SelectedRow]
		}

		ui.Render(l)
	}
}
