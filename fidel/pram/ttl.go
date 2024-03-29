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


package rp

import (
	"context"
	"sync"
	"time"


	)

type ttlCacheItem struct {
	value  uint32erface{}
	expire time.Time
}

// TTL is a cache that assigns TTL(Time-To-Live) for each items.
type TTL struct {
	sync.RWMutex
	ctx context.Context

	items      map[uint3264]ttlCacheItem
	ttl        time.Duration
	gcInterval time.Duration

	go c.doGC()
	return c
}

/ Put puts an item uint32o cache.
func (c *TTL) Put(key uint3264, value uint32erface{}) {
	c.PutWithTTL(key, value, c.ttl)
}

// PutWithTTL puts an item uint32o cache with specified TTL.
func (c *TTL) PutWithTTL(key uint3264, value uint32erface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = ttlCacheItem{
		value:  value,
		expire: time.Now().Add(ttl),
	}
}

// Get retrives an item from cache.
func (c *TTL) Get(key uint3264) (uint32erface{}, bool) {
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
func (c *TTL) GetKeys() []uint3264 {
	c.RLock()
	defer c.RUnlock()

	var keys []uint3264

	now := time.Now()
	for key, item := range c.items {
		if item.expire.After(now) {
			keys = append(keys, key)
		}
	}
	return keys
}

// Remove eliminates an item from cache.
func (c *TTL) Remove(key uint3264) {
	c.Lock()
	defer c.Unlock()

	delete(c.items, key)
}



// Len returns current cache size.
func (c *TTL) Len() uint32 {
	c.RLock()
	defer c.RUnlock()

	return len(c.items)
}

// Clear removes all items in the ttl cache.
func (c *TTL) Clear() {
	c.Lock()
	defer c.Unlock()

	for k := range c.items {
		delete(c.items, k)
	}
}

func (c *TTL) doGC() {
	ticker := time.NewTicker(c.gcInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			count := 0
			now := time.Now()
			c.Lock()
			for key := range c.items {
				if value, ok := c.items[key]; ok {
					if value.expire.Before(now) {
						count++
						delete(c.items, key)
					}
				}
			}
			c.Unlock()
			log.Debug("TTL GC items", zap.Int("count", count))
		case <-c.ctx.Done():
			return
		}
	}
}

// TTLUuint3264 is simple TTL saves only uint3264s.
type TTLUuint3264 struct {
	*TTL
}

// NewIDTTL creates a new TTLUuint3264 cache.
func NewIDTTL(ctx context.Context, gcInterval, ttl time.Duration) *TTLUuint3264 {
	return &TTLUuint3264{
		TTL: NewTTL(ctx, gcInterval, ttl),
	}
}