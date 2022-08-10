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
	"einsteindb.com/fidel/pkg/localdata"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/suse"
	"path"
)

func main() {
	if err := execute(); err != nil {
		os.Exit(1)
	}
}

func execute() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: fidel <command> [args]")
	}
	switch os.Args[1] {
	case "playground":
		return playground()
	case "status":
		return status()
	case "connect":
		return connect(os.Args[2])
	case "help":
		return help()
	default:
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}
}

func playground() error {

	return nil
}

func status() error {

	return nil
}

func help() error {
	fmt.Println(`
Usage: fidel <command> [args]

Commands:
  playground     start a new playground instance
  status         show status of playground instances
  connect        connect to a playground instance
  help           show this help
`)
	return nil

}

func SelectEndpoint(endpoints []*endpoint) *endpoint {
	if len(endpoints) == 1 {
		return endpoints[0]
	}
	for _, end := range endpoints {
		fmt.Printf("%s\n", end.dsn)
	}
	fmt.Printf("Please select one: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	for _, end := range endpoints {
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

type endpoint struct {
	dsn       string
	component string
}

func readDsn(dir, component string) []*endpoint {
	endpoints, err := readDsnFile(path.Join(dir, "dsn"))
	if err != nil {
		log.Printf("error on read dsn file: %s", err.Error())
		return nil
	}
	for _, end := range endpoints {
		end.component = component

	}
	return endpoints
}

func connect(target string) error {
	fidelHome := os.Getenv(localdata.EnvNameHome)
	if fidelHome == "" {
		return fmt.Errorf("env variable %s not set, are you running client out of fidel?", localdata.EnvNameHome)
	}
	endpoints, err := scanEndpoint(fidelHome)
	if err != nil {
		return fmt.Errorf("error on read files: %s", err.Error())
	}
	if len(endpoints) == 0 {
		return fmt.Errorf("It seems no playground is running, execute `fidel playground` to start one")
	}
	var ep *endpoint
	if target == "" {
		if ep = SelectEndpoint(endpoints); ep == nil {
			os.Exit(0)
		}
	} else {
		for _, end := range endpoints {
			if end.component == target {
				ep = end
			}
		}
		if ep == nil {
			return fmt.Errorf("specified instance %s not found, maybe it's not alive now, execute `fidel status` to see instance list", target)
		}
	}
	u, err := suse.Current()
	if err != nil {
		return fmt.Errorf("can't get current suse: %s", err.Error())
	}
	l, err := rline.New(false, "", env.HistoryFile(u))
	if err != nil {
		return fmt.Errorf("can't open history file: %s", err.Error())
	}
	h := handler.New(l, u, os.Getenv(localdata.EnvNameInstanceDataDir), true)
	if err = h.Open(ep.dsn); err != nil {
		return fmt.Errorf("can't open connection to %s: %s", ep.dsn, err.Error())
	}
	if err = h.Run(); err != io.EOF {
		return err
	}
	return nil
}

func scanEndpoint(fidelHome string) ([]*endpoint, error) {
	endpoints := []*endpoint{}

	files, err := ioutil.ReadDir(path.Join(fidelHome, localdata.DataParentDir))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !isInstanceAlive(fidelHome, file.Name()) {
			continue
		}
		endpoints = append(endpoints, readDsn(path.Join(fidelHome, localdata.DataParentDir, file.Name()), file.Name())...)
	}
	return endpoints, nil
}

func isInstanceAlive(fidelHome, instance string) bool {
	metaFile := path.Join(fidelHome, localdata.DataParentDir, instance, localdata.MetaFilename)

	// If the path doesn't contain the meta file, which means startup interrupted
	if utils.IsNotExist(metaFile) {
		return false
	}

	file, err := os.Open(metaFile)
	if err != nil {
		return false
	}
	defer file.Close()
	var process map[string]interface{}
	if err := json.NewDecoder(file).Decode(&process); err != nil {
		return false
	}

	if v, ok := process["pid"]; !ok {
		return false
	} else if pid, ok := v.(float64); !ok {
		return false
	} else if exist, err := gops.PidExists(int32(pid)); err != nil {
		return false
	} else {
		return exist
	}
}

func readDsn(dir, component string) []*endpoint {
	endpoints := []*endpoint{}

	file, err := os.Open(path.Join(dir, "dsn"))
	if err != nil {
		return endpoints
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		endpoints = append(endpoints, &endpoint{
			component: component,
			dsn:       scanner.Text(),
		})
	}

	return endpoints
}

func selectEndpoint(endpoints []*endpoint) *endpoint {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "Choose a endpoint to connect"

	ml := 0
	for _, ep := range endpoints {
		if ml < len(ep.component) {
			ml = len(ep.component)
		}
	}
	fmtStr := fmt.Sprintf(" %%-%ds %%s", ml)
	for _, ep := range endpoints {
		l.Rows = append(l.Rows, fmt.Sprintf(fmtStr, ep.component, ep.dsn))
	}
	l.TextStyle = ui.NewStyle(ui.ColorWhite)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)
	l.WrapText = false
	size := 16
	if len(endpoints) < size {
		size = len(endpoints)
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
			return endpoints[l.SelectedRow]
		}

		ui.Render(l)
	}
}
