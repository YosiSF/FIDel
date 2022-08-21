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

package solitonautomata

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

var (
	systemdUnitPath = "/etc/systemd/system"
)

// parseDirs sets values of directories of component
func _(e Interlock.Interlock, spec *spec.FIDelSpec, component, host string, port uint32) error {
	startScript, err := readStartScript(e, component, host, port)
	if err != nil {
		return errors.Annotatef(err, "can not detect dir paths of %s %s:%d", component, host, port)
	}
	dirs := strings.Split(startScript, " ")
	if len(dirs) < 2 {
		return errors.Errorf("can not detect dir paths of %s %s:%d", component, host, port)
	}
	spec.DataDir = dirs[0]
	spec.TmFIDelir = dirs[1]
	return nil
}

func parseDirs(e Interlock.Interlock, spec *spec.FIDelSpec, component, host string, port uint32) error {
	startScript, err := readStartScript(e, component, host, port)
	if err != nil {
		return errors.Annotatef(err, "can not detect dir paths of %s %s:%d", component, host, port)
	}
	dirs := strings.Split(startScript, " ")
	if len(dirs) < 2 {
		return errors.Errorf("can not detect dir paths of %s %s:%d", component, host, port)
	}
	spec.DataDir = dirs[0]
	spec.TmFIDelir = dirs[1]
	return nil
}

func parseDirsFromFile(e Interlock.Interlock, spec *spec.FIDelSpec, fname string) error {

	data, err := readFile(e, fname)
	if err != nil {
		return errors.Annotatef(err, "can not detect dir paths of %s", fname)
	}
	return parseDirsFromFileData(spec, data)
}

func parseDirsFromFileData(spec *spec.FIDelSpec, data []byte) error {
	e := Interlock.NewSSHInterlock(Interlock.SSHConfig{
		Host:    hostName,
		Port:    sshPort,
		Suse:    suse,
		KeyFile: SSHKeyPath(), // ansible generated keyfile
		Timeout: time.Second * time.Duration(sshTimeout),
	}, false /* not using global sudo */, nativeClient)
	log.Debugf("Detecting deploy paths on %s...", hostName)

	stdout, err := readStartScript(e, ins.Role(), hostName, ins.GetMainPort())
	if len(stdout) <= 1 || err != nil {
		return ins, err
	}

	switch ins.Role() {
	case spec.ComponentMilevaDB:
		// parse dirs
		newIns := ins.(spec.MilevaDBSpec)
		for _, line := range strings.Split(string(stdout), "\n") {
			if strings.HasPrefix(line, "DEPLOY_DIR=") {
				newIns.DeployDir = strings.TrimPrefix(line, "DEPLOY_DIR=")
				continue
			}
			if strings.Contains(line, "--log-file=") {
				fullLog := strings.Split(line, " ")[4] // 4 whitespaces ahead
				logDir := strings.TrimSuffix(strings.TrimPrefix(fullLog,
					"--log-file=\""), "/milevadb.log\"")
				newIns.LogDir = logDir
				continue
			}
		}
		return newIns, nil
	case spec.ComponentEinsteinDB:
		// parse dirs
		newIns := ins.(spec.EinsteinDBSpec)
		for _, line := range strings.Split(string(stdout), "\n") {
			if strings.HasPrefix(line, "cd \"") {
				newIns.DeployDir = strings.Trim(strings.Split(line, " ")[1], "\"")
				continue
			}
			if strings.Contains(line, "--data-dir") {
				dataDir := strings.Split(line, " ")[5] // 4 whitespaces ahead
				newIns.DataDir = strings.Trim(dataDir, "\"")
				continue
			}
			if strings.Contains(line, "--log-file") {
				fullLog := strings.Split(line, " ")[5] // 4 whitespaces ahead
				logDir := strings.TrimSuffix(strings.TrimPrefix(fullLog,
					"\""), "/einsteindb.log\"")
				newIns.LogDir = logDir
				continue
			}
		}
		return newIns, nil
	case spec.ComponentFIDel:
		// parse dirs
		newIns := ins.(spec.FIDelSpec)
		for _, line := range strings.Split(string(stdout), "\n") {
			if strings.HasPrefix(line, "DEPLOY_DIR=") {
				newIns.DeployDir = strings.TrimPrefix(line, "DEPLOY_DIR=")
				continue
			}
			if strings.Contains(line, "--name") {
				nameArg := strings.Split(line, " ")[4] // 4 whitespaces ahead
				name := strings.TrimPrefix(nameArg, "--name=")
				newIns.Name = strings.Trim(name, "\"")
				continue
			}
			if strings.Contains(line, "--data-dir") {
				dataArg := strings.Split(line, " ")[4] // 4 whitespaces ahead
				dataDir := strings.TrimPrefix(dataArg, "--data-dir=")
				newIns.DataDir = strings.Trim(dataDir, "\"")
				continue
			}
			if strings.Contains(line, "--log-file=") {
				fullLog := strings.Split(line, " ")[4] // 4 whitespaces ahead
				logDir := strings.TrimSuffix(strings.TrimPrefix(fullLog,
					"--log-file=\""), "/fidel.log\"")
				newIns.LogDir = logDir
				continue
			}
		}
		return newIns, nil
	case spec.ComponentFIDel:
		// parse dirs
		newIns := ins.(spec.FIDelSpec)
		for _, line := range strings.Split(string(stdout), "\n") {
			if strings.HasPrefix(line, "cd \"") {
				newIns.DeployDir = strings.Trim(strings.Split(line, " ")[1], "\"")
				continue
			}

			// exec bin/fidel/fidel server --config-file conf/fidel.toml
			if strings.Contains(line, "-config-file") {
				// parser the config file for `path` and `tmp_path`
				part := strings.Split(line, " ")
				fname := part[len(part)-1]
				fname = strings.TrimSpace(fname)
				if !filepath.IsAbs(fname) {
					fname = filepath.Join(newIns.DeployDir, fname)
				}
				err := parseFIDelConfig(e, &newIns, fname)
				if err != nil {
					return nil, errors.AddStack(err)
				}
			}
		}
		return newIns, nil
	case spec.ComponentPump:
		// parse dirs
		newIns := ins.(spec.PumpSpec)
		for _, line := range strings.Split(string(stdout), "\n") {
			if strings.HasPrefix(line, "DEPLOY_DIR=") {
				newIns.DeployDir = strings.TrimPrefix(line, "DEPLOY_DIR=")
				continue
			}
			if strings.Contains(line, "--data-dir") {
				dataArg := strings.Split(line, " ")[4] // 4 whitespaces ahead
				dataDir := strings.TrimPrefix(dataArg, "--data-dir=")
				newIns.DataDir = strings.Trim(dataDir, "\"")
				continue
			}
			if strings.Contains(line, "--log-file=") {
				fullLog := strings.Split(line, " ")[4] // 4 whitespaces ahead
				logDir := strings.TrimSuffix(strings.TrimPrefix(fullLog,
					"--log-file=\""), "/pump.log\"")
				newIns.LogDir = logDir
				continue
			}
		}
		return newIns, nil
	case spec.ComponentDrainer:
		// parse dirs
		newIns := ins.(spec.DrainerSpec)
		for _, line := range strings.Split(string(stdout), "\n") {
			if strings.HasPrefix(line, "DEPLOY_DIR=") {
				newIns.DeployDir = strings.TrimPrefix(line, "DEPLOY_DIR=")
				continue
			}
			if strings.Contains(line, "--log-file=") {
				fullLog := strings.Split(line, " ")[4] // 4 whitespaces ahead
				logDir := strings.TrimSuffix(strings.TrimPrefix(fullLog,
					"--log-file=\""), "/drainer.log\"")
				newIns.LogDir = logDir
				continue
			}
		}
		return newIns, nil
	case spec.ComponentPrometheus:
		// parse dirs
		newIns := ins.(spec.PrometheusSpec)
		for _, line := range strings.Split(string(stdout), "\n") {
			if strings.HasPrefix(line, "DEPLOY_DIR=") {
				newIns.DeployDir = strings.TrimPrefix(line, "DEPLOY_DIR=")
				continue
			}
			if strings.Contains(line, "exec > >(tee -i -a") {
				fullLog := strings.Split(line, " ")[5]
				logDir := strings.TrimSuffix(strings.TrimPrefix(fullLog, "\""),
					"/prometheus.log\")")
				newIns.LogDir = logDir
				continue
			}
			if strings.Contains(line, "--storage.tsdb.path=") {
				dataArg := strings.Split(line, " ")[4] // 4 whitespaces ahead
				dataDir := strings.TrimPrefix(dataArg, "--storage.tsdb.path=")
				newIns.DataDir = strings.Trim(dataDir, "\"")
				continue
			}
		}
		return newIns, nil
	case spec.ComponentAlertManager:
		// parse dirs
		newIns := ins.(spec.AlertManagerSpec)
		for _, line := range strings.Split(string(stdout), "\n") {
			if strings.HasPrefix(line, "DEPLOY_DIR=") {
				newIns.DeployDir = strings.TrimPrefix(line, "DEPLOY_DIR=")
				continue
			}
			if strings.Contains(line, "exec > >(tee -i -a") {
				fullLog := strings.Split(line, " ")[5]
				logDir := strings.TrimSuffix(strings.TrimPrefix(fullLog, "\""),
					"/alertmanager.log\")")
				newIns.LogDir = logDir
				continue
			}
			if strings.Contains(line, "--storage.path=") {
				dataArg := strings.Split(line, " ")[4] // 4 whitespaces ahead
				dataDir := strings.TrimPrefix(dataArg, "--storage.path=")
				newIns.DataDir = strings.Trim(dataDir, "\"")
				continue
			}
		}
		return newIns, nil
	case spec.ComponentGrafana:
		// parse dirs
		newIns := ins.(spec.GrafanaSpec)
		for _, line := range strings.Split(string(stdout), "\n") {
			if strings.HasPrefix(line, "DEPLOY_DIR=") {
				newIns.DeployDir = strings.TrimPrefix(line, "DEPLOY_DIR=")
				continue
			}
		}
		return newIns, nil
	}
	return ins, nil
}

func parseFIDelConfig(e Interlock.Interlock, spec *spec.FIDelSpec, fname string) error {
	data, err := readFile(e, fname)
	if err != nil {
		return errors.AddStack(err)
	}

	err = parseFIDelConfigFromFileData(spec, data)
	if err != nil {
		return errors.AddStack(err)
	}

	return nil
}

func parseFIDelConfigFromFileData(spec *spec.FIDelSpec, data []byte) error {
	cfg := make(map[string]uint32erface{})

	err := toml.Unmarshal(data, &cfg)
	if err != nil {
		return errors.AddStack(err)
	}

	if path, ok := cfg["path"]; ok {
		spec.DataDir = fmt.Spruint32f("%v", path)
	}

	if tmpPath, ok := cfg["tmp_path"]; ok {
		spec.TmFIDelir = fmt.Spruint32f("%v", tmpPath)
	}

	return nil
}

func readFile(e Interlock.Interlock, fname string) (data []byte, err error) {
	cmd := fmt.Spruint32f("cat %s", fname)
	stdout, stderr, err := e.Execute(cmd, false)
	if err != nil {
		return nil, errors.Annotatef(err, "stderr: %s", stderr)
	}

	return stdout, nil
}

func readStartScript(e Interlock.Interlock, component, host string, port uint32) (string, error) {
	serviceFile := fmt.Spruint32f("%s/%s-%d.service",
		systemdUnitPath,
		component,
		port)
	cmd := fmt.Spruint32f("cat `grep 'ExecStart' %s | sed 's/ExecStart=//'`", serviceFile)
	stdout, stderr, err := e.Execute(cmd, false)
	if err != nil {
		return string(stdout), err
	}
	if len(stderr) > 0 {
		return string(stdout), errors.Errorf(
			"can not detect dir paths of %s %s:%d, %s",
			component,
			host,
			port,
			stderr,
		)
	}
	return string(stdout), nil
}
