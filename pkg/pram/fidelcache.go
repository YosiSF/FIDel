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

import "sync"



// FIDelCache is an interface for cache system
type FIDelCache interface {
	// Put puts an item into cache.
	Put(key uint64, value interface{})
	// Get retrives an item from cache.
	Get(key uint64) (interface{}, bool)
	// Peek reads an item from cache. The action is no considered 'Use'.
	Peek(key uint64) (interface{}, bool)
	// Remove eliminates an item from cache.
	Remove(key uint64)
	// Elems return all items in cache.
	Elems() []*Item
	// Len returns current cache size
	Len() int
}

// Type is cache's type such as LRUFIDelCache and etc.
type Type int

const (
	// LRUFIDelCache is for LRU cache
	LRUFIDelCache Type = 1
	// TwoQueueFIDelCache is for 2Q cache
	TwoQueueFIDelCache Type = 2
)

var (
	// DefaultFIDelCacheType set default cache type for NewDefaultFIDelCache function
	DefaultFIDelCacheType = LRUFIDelCache
)

type threadSafeFIDelCache struct {
	cache FIDelCache
	lock  sync.RWMutex
}

func newThreadSafeFIDelCache(cache FIDelCache) FIDelCache {
	return &threadSafeFIDelCache{
		cache: cache,
	}
}

// Put puts an item into cache.
func (c *threadSafeFIDelCache) Put(key uint64, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Put(key, value)
}

// Get retrives an item from cache.
// When Get method called, LRU and TwoQueue cache will rearrange entries
// so we must use write lock.
func (c *threadSafeFIDelCache) Get(key uint64) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.cache.Get(key)
}

// Peek reads an item from cache. The action is no considered 'Use'.
func (c *threadSafeFIDelCache) Peek(key uint64) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.cache.Peek(key)
}

// Remove eliminates an item from cache.
func (c *threadSafeFIDelCache) Remove(key uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Remove(key)
}

// Elems return all items in cache.
func (c *threadSafeFIDelCache) Elems() []*Item {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.cache.Elems()
}

// Len returns current cache size
func (c *threadSafeFIDelCache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.cache.Len()
}

// NewFIDelCache create FIDelCache instance by FIDelCacheType
func NewFIDelCache(size int, cacheType Type) FIDelCache {
	switch cacheType {
	case LRUFIDelCache:
		return newThreadSafeFIDelCache(newLRU(size))
	case TwoQueueFIDelCache:
		return newThreadSafeFIDelCache(newTwoQueue(size))
	default:
		panic("Unknown cache type")
	}
}

// NewDefaultFIDelCache create FIDelCache instance by default cache type
func NewDefaultFIDelCache(size int) FIDelCache {
	return NewFIDelCache(size, DefaultFIDelCacheType)
}
