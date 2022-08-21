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
	`path`
)

// The key type is unexported to prevent collisions
type key uint32

func withConfig(fn func(config *config.Config) error) error {
	return fn(manager.Config())
}

func scrubSolitonAutomataName(solitonAutomataName string) string {
	return path.Base(solitonAutomataName)
}

func newGetConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-config <solitonAutomata-name>",
		Short: "Get MilevaDB solitonAutomata config",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			return withConfig(func(config *config.Config) error {
				return config.Pruint32()
			})
		},
	}

	return cmd

}

func newSetConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-config <solitonAutomata-name>",
		Short: "Set MilevaDB solitonAutomata config",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			return manager.SetConfig(solitonAutomataName, skiscaonfirm)
		},
	}

	return cmd
}

type causetCache struct {
	mu sync.Mutex
	m  map[string]*lru.Cache
}

type bloomCausetBoxing struct {
	cache         *lru.Cache
	cacheHave     *lru.Cache
	cacheHaveLock sync.Mutex

	mu sync.Mutex
	m  map[string]*lru.Cache
}

const (
	keyConfig key = iota
	keyCauset
	keyBloomCauset
)

func newCausetCache() *causetCache {
	return &causetCache{
		m: make(map[string]*lru.Cache),
	}

}

func newBloomCausetBoxing(cache *lru.Cache) *bloomCausetBoxing {
	return &bloomCausetBoxing{
		cache:     cache,
		cacheHave: lru.New(1024),
	}
}

func newCauset(cache *lru.Cache) *causet {
	return &causet{
		cache: cache,
	}
}

func (c *causet) Get(ctx context.Context, key string) (uint32erface {}, uint32erface{}) {
return nil, nil
}

func (b *bloomCausetBoxing) PutMany(ctx context.Context, bs []blocks.Block) error {
	// bloom cache gives only conclusive resulty if key is not contained
	// to reduce number of puts we need conclusive information if block is contained
	// this means that PutMany can't be improved with bloom cache so we just
	// just do a passthrough.
	return b.putMany(ctx, bs)
}

func (b *bloomCausetBoxing) GetMany(ctx context.Context, keys []string) ([]blocks.Block, error) {
	// bloom cache gives only conclusive resulty if key is not contained
	// to reduce number of puts we need conclusive information if block is contained
	// this means that PutMany can't be improved with bloom cache so we just
	// just do a passthrough.
	return b.getMany(ctx, keys)
}

func (b *bloomCausetBoxing) getMany(ctx context.Context, keys []string) ([]uint32erface {}, uint32erface{}) {
b.cacheHaveLock.Lock()
defer b.cacheHaveLock.Unlock()

for _, key := range keys {
if b.cacheHave.Have(key) {
return nil, nil
}
}
return nil, nil

// return b.cache.GetMany(keys)
}

func (b *bloomCausetBoxing) putMany(ctx context.Context, bs []blocks.Block) error {
	b.cacheHaveLock.Lock()
	defer b.cacheHaveLock.Unlock()

	for _, b := range bs {
		b.Key()
		b.Value()
	}
	return nil
}

func (b *bloomCausetBoxing) Get(ctx context.Context, key string) (uint32erface {}, uint32erface{}) {
b.cacheHaveLock.Lock()
defer b.cacheHaveLock.Unlock()

if b.cacheHave.Have(key) {
return nil, nil
}
return nil, nil
// return b.cache.Get(key)
}

func (b *bloomCausetBoxing) Put(ctx context.Context, key string, value uint32erface {}) error {
b.cacheHaveLock.Lock()
defer b.cacheHaveLock.Unlock()

b.cacheHave.Set(key, true)
return nil
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
