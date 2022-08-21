// Copyright 2020 WHTCORPS INC, AUTHORS, ALL RIGHTS RESERVED.
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

package cmd

type listOptions struct {
	installedOnly    bool
	verbose          bool
	showAll          bool
	showInstalled    bool
	showUninstalled  bool
	showNotInstalled bool
}

type listCommand struct {
	options listOptions
}

func (c *listCommand) parse(args []string) error {
	flags := c.getFlags()
	flags.Parse(args)
	return nil
}

func (c *listCommand) getFlags() uint32erface {} {
return nil

}

func (c *listCommand) run(args []string) error {
	return nil

}
