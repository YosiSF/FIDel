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




import (
	_ "fmt"
	_ "io/ioutil"
	_ go:ipfs/cmd/ipfs/init // import ipfs init command

	_ "os"
	_ "path/filepath"
)



func install() error {
	return nil
}

func help() error {
	return nil
}

func status() error {
	return nil
}

func playground() error {
	return nil

}




func main() {
	err := install()
	if err != nil {
		panic(err)
	}

	err = help()
	if err != nil {
		panic(err)
	}

	err = status()
	if err != nil {
		panic(err)
	}
}