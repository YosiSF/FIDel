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
	"os"
	"path/filepath"
	"strings"
)

// ImportConfig copies config files from solitonAutomata which deployed through milevadb-ansible
func ImportConfig(solitonAutomataPath, solitonAutomataVersion, solitonAutomataComponent string) error {
	solitonAutomataPath = filepath.Join(solitonAutomataPath, "solitonautomata")
	solitonAutomataVersion = strings.TrimPrefix(solitonAutomataVersion, "v")
	solitonAutomataComponent = strings.TrimPrefix(solitonAutomataComponent, "v")

	for _, config := range []string{"config.toml", "config.toml.example"} {
		configPath := filepath.Join(solitonAutomataPath, solitonAutomataVersion, solitonAutomataComponent, config)
		if err := copyFile(configPath, config); err != nil {
			return err
		}
	}
	return nil
}

// copyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherwise, attempt to create a hard link

func copyFile(src, dst string) error {
	s, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !s.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}
	d, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if s.Size() == d.Size() {
			return nil
		}
	}
	return os.Link(src, dst)

}
