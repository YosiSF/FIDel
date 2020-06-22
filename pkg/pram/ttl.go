//Copyright 2020 WHTCORPS INC All Rights Reserved
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


package pram

import (
	"context"
	"sync"
	"time"


	)

type ttlCacheItem struct {
	value  interface{}
	expire time.Time
}

// TTL is a cache that assigns TTL(Time-To-Live) for each items.
type TTL struct {
	sync.RWMutex
	ctx context.Context

	items      map[uint64]ttlCacheItem
	ttl        time.Duration
	gcInterval time.Duration

	go c.doGC()
	return c
}

/ Put puts an item into cache.
func (c *TTL) Put(key uint64, value interface{}) {
	c.PutWithTTL(key, value, c.ttl)
}

// PutWithTTL puts an item into cache with specified TTL.
func (c *TTL) PutWithTTL(key uint64, value interface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = ttlCacheItem{
		value:  value,
		expire: time.Now().Add(ttl),
	}
}

// Get retrives an item from cache.
func (c *TTL) Get(key uint64) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if item.expire.Before(time.Now()) {
		return nil, false
	}

	return item.value, true
}

// GetKeys returns all keys that are not expired.
func (c *TTL) GetKeys() []uint64 {
	c.RLock()
	defer c.RUnlock()

	var keys []uint64

	now := time.Now()
	for key, item := range c.items {
		if item.expire.After(now) {
			keys = append(keys, key)
		}
	}
	return keys
}

// Remove eliminates an item from cache.
func (c *TTL) Remove(key uint64) {
	c.Lock()
	defer c.Unlock()

	delete(c.items, key)
}