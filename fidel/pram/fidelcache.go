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
	_ "time"
)

// HTTP headers.
const (
	RedirectorHeader    = "PD-Redirector"
	AllowFollowerHandle = "PD-Allow-follower-handle"
)

const (
	errRedirectFailed      = "redirect failed"
	errRedirectToNotLeader = "redirect to not leader"
)

type runtimeServiceValidator struct {
	s     *server.Server
	group server.ServiceGroup
}

type threadSafeFIDelCache struct {
	cache FIDelCache
	lock  sync.RWMutex
}

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

type Item struct {
	Key   uint64
	Value interface{}
}

type FIDelCacheItem struct {
	Key   uint64
	Value interface{}
}

type Event interface{}

// EmitterInterface Root interface for events dispatch
type EmitterInterface interface {
	// Deprecated: Emit Sends an event to the subscribed listeners
	Emit(context.Context, Event)

	// Deprecated: GlobalChannel returns a glocal channel that receives emitted events
	GlobalChannel(ctx context.Context) <-chan Event

	// Deprecated: Subscribe Returns a channel that receives emitted events
	Subscribe(ctx context.Context) <-chan Event

	// Deprecated: UnsubscribeAll close all listeners channels
	UnsubscribeAll()
}

// Deprecated: use event bus directly
// EventEmitter Registers listeners and dispatches events to them
type EventEmitter struct {
	listeners map[string][]chan Event
	lock      sync.RWMutex

	bus        event.Bus
	muEmitters sync.Mutex

	cglobal <-chan Event

	emitter event.Emitter
	cancels []context.CancelFunc

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup

	closed bool

	cache FIDelCache

	capacity int

	evictedHandler func(key uint64, value interface{})
}

const (

	// Evicted is the event type for when an item is evicted from the cache.
	// The event will be sent on the global channel.
	Evicted Event = "evicted"
	// LRUFIDelCache is for LRU cache
	LRUFIDelCache Type = 1
	// TwoQueueFIDelCache is for 2Q cache
	TwoQueueFIDelCache Type = 2
	// FIFO is for FIFO cache
	//
	// Deprecated: use FIFOFIDelCache instead
	FIFO Type = 3
	// FIFOFIDelCache is for FIFO cache
)

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

func (c *threadSafeFIDelCache) String() string {
	return c.cache.String()

}

func (c *threadSafeFIDelCache) Stats() map[string]interface{} {
	return c.cache.Stats()
}

func (c *threadSafeFIDelCache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Clear()
}

func (c *threadSafeFIDelCache) Close() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Close()
}

func (c *threadSafeFIDelCache) SetCapacity(capacity int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.SetCapacity(capacity)
}

func (c *threadSafeFIDelCache) SetEvictedHandler(handler func(key uint64, value interface{})) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.SetEvictedHandler(handler)
}

func (c *threadSafeFIDelCache) SetExpiredHandler(handler func(key uint64, value interface{})) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.SetExpiredHandler(handler)
}

func (c *threadSafeFIDelCache) SetExpired(key uint64, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.SetExpired(key, value)
}

func (c *threadSafeFIDelCache) SetEvicted(key uint64, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.SetEvicted(key, value)

}

// NewDefaultFIDelCache create FIDelCache instance by default cache type
func NewDefaultFIDelCache(size int) FIDelCache {
	return NewFIDelCache(size, DefaultFIDelCacheType)
}

///#![feature(const_fn)]
//#![feature(const_fn_transmute)]
//#![feature(const_fn_transmute_pure)]
