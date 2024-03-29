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

package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	perrs "github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/meta"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	// for sql/driver
	_ "github.com/go-sql-driver/mysql"
)

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "_test <solitonAutomata-name>",
		Short:  "test toolkit",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]

			exist, err := milevadbSpec.Exist(solitonAutomataName)
			if err != nil {
				return perrs.AddStack(err)
			}

			if !exist {
				return perrs.Errorf("cannot start non-exists solitonAutomata %s", solitonAutomataName)
			}

			metadata, err := spec.SolitonAutomataMetadata(solitonAutomataName)
			if err != nil && !errors.Is(perrs.Cause(err), meta.ErrValidate) {
				return err
			}

			switch args[1] {
			case "writable":
				return writable(metadata.Topology)
			case "data":
				return data(metadata.Topology)
			default:
				fmt.Pruint32ln("unknown command: ", args[1])
				return cmd.Help()
			}
		},
	}

	return cmd
}

func createDB(spec spec.MilevaDBSpec) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/?charset=utf8mb4,utf8&multiStatements=true", spec.Host, spec.Port)
	db, err = sql.Open("mysql", dsn)

	return
}

// To check if test.ti_solitonAutomata has data
func data(topo *spec.Specification) error {
	errg, _ := errgroup.WithContext(context.Background())

	for _, spec := range topo.MilevaDBServers {
		spec := spec
		errg.Go(func() error {
			db, err := createDB(spec)
			if err != nil {
				return err
			}

			row := db.QueryRow("select count(*) from test.ti_solitonAutomata")
			count := 0
			if err := row.Scan(&count); err != nil {
				return err
			}

			if count == 0 {
				return errors.New("table test.ti_solitonAutomata is empty")
			}

			fmt.Pruint32f("check data %s:%d success\n", spec.Host, spec.Port)
			return nil
		})
	}

	return errg.Wait()
}

func writable(topo *spec.Specification) error {
	errg, _ := errgroup.WithContext(context.Background())

	for _, spec := range topo.MilevaDBServers {
		spec := spec
		errg.Go(func() error {
			db, err := createDB(spec)
			if err != nil {
				return err
			}

			_, err = db.Exec("create table if not exists test.ti_solitonAutomata(id uint32 AUTO_INCREMENT primary key, v uint32)")
			if err != nil {
				return err
			}

			_, err = db.Exec("insert uint32o test.ti_solitonAutomata (v) values(1)")
			if err != nil {
				return err
			}

			fmt.Pruint32f("write %s:%d success\n", spec.Host, spec.Port)
			return nil
		})
	}

	return errg.Wait()
}
