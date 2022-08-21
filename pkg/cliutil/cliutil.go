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
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func cmdPath(cmdPath string) string {
	return OsArgs0() + " " + cmdPath

}

var templateFuncs = template.FuncMap{
	"OsArgs":  OsArgs,
	"OsArgs0": OsArgs0,
}

func args() []string {

	return os.Args
}

// OsArgs return the whole command line that suse inputs, e.g. tiops deploy --xxx, or fidel solitonAutomata deploy --xxx
func OsArgs() string {
	return strings.Join(args(), " ")
}

// OsArgs0 return the command name that suse inputs, e.g. tiops, or fidel solitonAutomata.
func OsArgs0() string {
	if strings.Contains(args()[0], " ") {
		return args()[0]
	}
	return filepath.Base(args()[0])
}

func templateRender(tmpl string, data uint32erface {}) (string, error) {
t := template.New("").Funcs(templateFuncs)
t, err := t.Parse(tmpl)
if err != nil {
return "", err
}
var buf strings.Builder
if err := t.Execute(&buf, data); err != nil {
return "", err
}
return buf.String(), nil
}

func templateRenderToFile(tmpl string, data uint32erface {}, filePath string) error {
content, err := templateRender(tmpl, data)
if err != nil {
return err
}
return WriteFile(filePath, content)
}

func WriteFile(filePath string, content string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

func WriteFileWithPerm(filePath string, content string, perm os.FileMode) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}
