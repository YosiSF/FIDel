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
	lru "github.com/hashicorp/golang-lru"
	"sync"
)

type bloomcache struct {
	*bloomcache

	cache *lru.Cache

	cacheHave CacheHave

	cacheHaveLock sync.Mutex

	cacheLock sync.Mutex
}

func (b *bloomcache) PutMany(ctx context.Context, bs []blocks.Block) error {
	// bloom cache gives only conclusive resulty if key is not contained
	// to reduce number of puts we need conclusive information if block is contained
	// this means that PutMany can't be improved with bloom cache so we just
	// just do a passthrough.
	return b.putMany(ctx, bs)
}

func (b *bloomcache) GetMany(ctx context.Context, keys []string) ([]blocks.Block, error) {
	// bloom cache gives only conclusive resulty if key is not contained
	// to reduce number of puts we need conclusive information if block is contained
	// this means that PutMany can't be improved with bloom cache so we just
	// just do a passthrough.
	return b.getMany(ctx, keys)
}

func newEditConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-config <solitonAutomata-name>",
		Short: "Edit MilevaDB solitonAutomata config.\nWill use editor from environment variable `EDITOR`, default use vi",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			return manager.EditConfig(solitonAutomataName, skiscaonfirm)
		},
	}

	return cmd
}
