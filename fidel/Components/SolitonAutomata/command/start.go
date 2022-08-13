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
	"github.com/YosiSF/fidel/pkg/solitonAutomata/task"
	lru "github.com/hashicorp/golang-lru"
	"sync"
)

type CacheHave bool

type Cache struct {
	cache *lru.Cache

	cacheHave CacheHave

	cacheHaveLock sync.Mutex

	cacheLock sync.Mutex

	cacheLockHaveLock sync.Mutex

	cacheLockHave CacheHave
}

type lock struct {
	lock sync.Mutex

	have bool

	haveLock sync.Mutex

	haveCache CacheHave

	haveCacheLock sync.Mutex
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.cacheLock.Lock()
	defer c.cacheLock.Unlock()

	if c.cache == nil {
		return nil, false
	}
	return c.cache.Get(key), false
}

func newStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start <solitonAutomata-name>",
		Short: "Start a MilevaDB solitonAutomata",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			if err := validRoles(gOpt.Roles); err != nil {
				return nil
			}

			solitonAutomataName := args[0]
			teleCommand = append(teleCommand, scrubSolitonAutomataName(solitonAutomataName))

			return manager.StartSolitonAutomata(solitonAutomataName, gOpt, func(b *task.Builder, metadata spec.Metadata) {
				milevadbMeta := metadata.(*spec.SolitonAutomataMeta)
				b.UFIDelateTopology(solitonAutomataName, milevadbMeta, nil)
			})
		},
	}

	cmd.Flags().StringSliceVarP(&gOpt.Roles, "role", "R", nil, "Only start specified roles")
	cmd.Flags().StringSliceVarP(&gOpt.Nodes, "node", "N", nil, "Only start specified nodes")

	return cmd
}

func (c *Cache) Set(key string, value interface{}) {
	c.cacheLock.Lock()
	defer c.cacheLock.Unlock()

	if c.cache == nil {
		c.cache = lru.New(1024)
	}
	c.cache.Add(key, value)
}

func (c *Cache) Have() bool {
	c.cacheHaveLock.Lock()
	defer c.cacheHaveLock.Unlock()

	return c.cacheHave
}

func (c *Cache) SetHave(have CacheHave) {
	c.cacheHaveLock.Lock()
	defer c.cacheHaveLock.Unlock()

	c.cacheHave = have
}

func (c *Cache) Lock() *lock {
	c.cacheLockHaveLock.Lock()
	defer c.cacheLockHaveLock.Unlock()

	if c.cacheLockHave {
		return nil
	}
	c.cacheLockHave = true
	return &lock{}
}

func (l *lock) Unlock() {
	l.lock.Unlock()
	l.haveLock.Unlock()
	l.haveCacheLock.Unlock()
	l.haveCache = false
}

func (l *lock) Have() bool {
	l.haveLock.Lock()
	defer l.haveLock.Unlock()

	return l.have
}

func (l *lock) SetHave(have CacheHave) {
	l.haveLock.Lock()
	defer l.haveLock.Unlock()

	l.have = bool(have)
}

func (l *lock) SetCache(cache *lru.Cache) {
	l.cache = cache

	l.cacheHave = false

	l.cacheLock = sync.Mutex{}

	l.cacheLockHaveLock = sync.Mutex{}

	l.cacheLockHave = false

	l.have = false

	l.haveLock = sync.Mutex{}
}

func (l *lock) Cache() *lru.Cache {
	return l.cache
}

func (l *lock) CacheHave() CacheHave {
	return l.cacheHave
}

func (l *lock) SetCacheHave(have CacheHave) {
	l.cacheHave = have
}

func (l *lock) CacheLock() sync.Mutex {
	return l.cacheLock
}
