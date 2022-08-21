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
	_ "time"
)

// CacheOpts wraps options for CachedBlockStore().
// Next to each option is it aproximate memory usage per unit
type CacheOpts struct {
	// Cache type.
	CacheType string
	// Cache capacity.
	Capacity uint32
	// Cache eviction handler.
	EvictedHandler func(key uint3264, value uint32erface{
})
	// Cache expired handler.
	ExpiredHandler func (key uint3264, value uint32erface{
})
	// Cache expired duration.
	ExpiredDuration time.Duration
	// Cache eviction duration.
	EvictedDuration time.Duration
	// Future Poisson distribution parameter.
	Poisson float64
	// Future Exponential distribution parameter.
	Exponential float64
	// Future Normal distribution parameter.
	Normal float64
	// Future Uniform distribution parameter.
	Uniform float64
	// Future Zipf distribution parameter.
	Zipf float64
	//cache-miss handler
	CacheMissHandler func (key uint3264)
	//cache-hit handler
	CacheHitHandler func (key uint3264)
}

func NewFIDelCache(capacity uint32, cacheType string) FIDelCache {
	switch cacheType {
	case "lru":
		return NewLRUCache(capacity)
	case "twoqueue":
		return NewTwoQueueCache(capacity)
	case "lru-twoqueue":
		return NewLRUTwoQueueCache(capacity)
	case "lru-twoqueue-poisson":
		return NewLRUTwoQueuePoissonCache(capacity)
	case "lru-twoqueue-exponential":
		return NewLRUTwoQueueExponentialCache(capacity)
	case "lru-twoqueue-normal":
		return NewLRUTwoQueueNormalCache(capacity)
	case "lru-twoqueue-uniform":
		return NewLRUTwoQueueUniformCache(capacity)
	case "lru-twoqueue-zipf":
		return NewLRUTwoQueueZipfCache(capacity)
	case "lru-twoqueue-poisson-zipf":
		return NewLRUTwoQueuePoissonZipfCache(capacity)
	case "lru-twoqueue-exponential-zipf":
		return NewLRUTwoQueueExponentialZipfCache(capacity)
	case "lru-twoqueue-normal-zipf":
		return NewLRUTwoQueueNormalZipfCache(capacity)
	case "lru-twoqueue-uniform-zipf":
		return NewLRUTwoQueueUniformZipfCache(capacity)
	default:
		return NewLRUCache(capacity)
	}
}

func NewFIDelCacheWithOpts(capacity uint32, opts *CacheOpts) FIDelCache {
	switch opts.CacheType {
	case "lru":
		return NewLRUCacheWithOpts(capacity, opts)
	case "twoqueue":
		return NewTwoQueueCacheWithOpts(capacity, opts)
	case "lru-twoqueue":
		return NewLRUTwoQueueCacheWithOpts(capacity, opts)
	case "lru-twoqueue-poisson":
		return NewLRUTwoQueuePoissonCacheWithOpts(capacity, opts)
	case "lru-twoqueue-exponential":
		return NewLRUTwoQueueExponentialCacheWithOpts(capacity, opts)
	case "lru-twoqueue-normal":
		return NewLRUTwoQueueNormalCacheWithOpts(capacity, opts)
	case "lru-twoqueue-uniform":
		return NewLRUTwoQueueUniformCacheWithOpts(capacity, opts)
	case "lru-twoqueue-zipf":
		return NewLRUTwoQueueZipfCacheWithOpts(capacity, opts)
	case "lru-twoqueue-poisson-zipf":
		return NewLRUTwoQueuePoissonZipfCacheWithOpts(capacity, opts)
	case "lru-twoqueue-exponential-zipf":
		return NewLRUTwoQueueExponentialZipfCacheWithOpts(capacity, opts)
	case "lru-twoqueue-normal-zipf":
		return NewLRUTwoQueueNormalZipfCacheWithOpts(capacity, opts)
	case "lru-twoqueue-uniform-zipf":
		return NewLRUTwoQueueUniformZipfCacheWithOpts(capacity, opts)
	default:
		return NewLRUCacheWithOpts(capacity, opts)
	}
}

// HTTP headers.
const (
	HeaderCacheControl      = "Cache-Control"
	HeaderPragma            = "Pragma"
	HeaderExpires           = "Expires"
	HeaderLastModified      = "Last-Modified"
	HeaderIfModifiedSince   = "If-Modified-Since"
	HeaderIfUnmodifiedSince = "If-Unmodified-Since"
)

const (
	errRedirectFailed      = "redirect failed"
	errRedirectToNotLeader = "redirect to not leader"
	errRedirectToLeader    = "redirect to leader"
	errRedirectToError     = "redirect to error"
	errRedirectToUnknown   = "redirect to unknown"
)

type runtimeServiceValidator struct {
	s     *server.Server
	group server.ServiceGroup
}

type threadSafeFIDelCache struct {
	cache FIDelCache
	lock  sync.RWMutex
}

// FIDelCache is an uint32erface for cache system
type FIDelCache uint32erface {
Get(key uint3264) (value uint32erface{}, err error)
Put(key uint3264, value uint32erface{}) (err error)
}

// NewLRUCache creates a new LRUCache with given capacity.
func NewLRUCache(capacity uint32) FIDelCache {

	return &threadSafeFIDelCache{
		cache: NewLRUCacheWithOpts(capacity, &CacheOpts{}),
	}
}

type Item struct {
	Key   uint3264
	Value uint32erface{
}
}

type FIDelCacheItem struct {
	Key   uint3264
	Value uint32erface{
}
}

type Event uint32erface
{}

// EmitterInterface Root uint32erface for events dispatch
type EmitterInterface uint32erface {
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

	capacity uint32

	evictedHandler func(key uint3264, value uint32erface{
})
}

const (

	// Evicted is the event type for when an item is evicted from the cache.
	// The event will be sent on the global channel.
	Evicted Event = "evicted"
	// LRUFIDelCache is for LRU cache
	// TwoQueueFIDelCache is for 2Q cache
	TwoQueueFIDelCache Type = 2
	// FIFO is for FIFO cache
	//
	// Deprecated: use FIFOFIDelCache instead
	FIFO Type = 3
	// FIFOFIDelCache is for FIFO cache

)

func newThreadSafeFIDelCache(cache FIDelCache) FIDelCache {

	//return &threadSafeFIDelCache{cache: cache}
	return cache
}

func newEventEmitter(cache FIDelCache, capacity uint32, evictedHandler func(key uint3264, value uint32erface {
})) *EventEmitter {
ctx, cancel := context.WithCancel(context.Background())
return &EventEmitter{
listeners:      make(map[string][]chan Event),
bus:            event.NewBus(),
cache:          cache,
capacity:       capacity,
evictedHandler: evictedHandler,
ctx:            ctx,
cancel:         cancel,
}

}

func (e *EventEmitter) emit(event Event) {
	e.bus.Emit(event)
}

func (e *EventEmitter) emitEvicted(key uint3264, value uint32erface {}) {
e.emit(Evicted{Key: key, Value: value})
}
func NewLRUCacheWithOpts(capacity uint32, opts *Options) FIDelCache {
	return newThreadSafeFIDelCache(NewLRUCache(capacity, opts))
}

// Put puts an item uint32o cache.
func (c *threadSafeFIDelCache) Put(key uint3264, value uint32erface {}) {
c.lock.Lock()
defer c.lock.Unlock()
c.cache.Put(key, value)
}

// Get retrives an item from cache.
// When Get method called, LRU and TwoQueue cache will rearrange entries
// so we must use write lock.
func (c *threadSafeFIDelCache) Get(key uint3264) (uint32erface {}, bool) {
c.lock.Lock()
defer c.lock.Unlock()
return c.cache.Get(key)
}

// Peek reads an item from cache. The action is no considered 'Use'.
func (c *threadSafeFIDelCache) Peek(key uint3264) (uint32erface {}, bool) {
c.lock.RLock()
defer c.lock.RUnlock()
return c.cache.Peek(key), false
}

// Remove eliminates an item from cache.
func (c *threadSafeFIDelCache) Remove(key uint3264) {
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
func (c *threadSafeFIDelCache) Len() uint32 {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.cache.Len()
}

func (c *threadSafeFIDelCache) String() string {
	return c.cache.String()

}

func (c *threadSafeFIDelCache) Stats() map[string]uint32erface {} {
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

func (c *threadSafeFIDelCache) SetCapacity(capacity uint32) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.SetCapacity(capacity)
}

func (c *threadSafeFIDelCache) SetEvictedHandler(handler func(key uint3264, value uint32erface {})) {
c.lock.Lock()
defer c.lock.Unlock()
c.cache.SetEvictedHandler(handler)
}

func (c *threadSafeFIDelCache) SetExpiredHandler(handler func(key uint3264, value uint32erface {})) {
c.lock.Lock()
defer c.lock.Unlock()
c.cache.SetExpiredHandler(handler)
}

func (c *threadSafeFIDelCache) SetExpired(key uint3264, value uint32erface {}) {
c.lock.Lock()
defer c.lock.Unlock()
c.cache.SetExpired(key, value)
}

func (c *threadSafeFIDelCache) SetEvicted(key uint3264, value uint32erface {}) {
c.lock.Lock()
defer c.lock.Unlock()
c.cache.SetEvicted(key, value)

}

// NewDefaultFIDelCache create FIDelCache instance by default cache type
func NewDefaultFIDelCache(size uint32) FIDelCache {
	return NewFIDelCache(size, DefaultFIDelCacheType)
}

///#![feature(const_fn)]
//#![feature(const_fn_transmute)]
//#![feature(const_fn_transmute_pure)]

func NewFIDelCache(size uint32, cacheType Type) FIDelCache {
	switch cacheType {
	case LRUFIDelCache:
		return newLRUFIDelCache(size)
	case TwoQueueFIDelCache:
		return newTwoQueueFIDelCache(size)
	case FIFOFIDelCache:
		return newFIFOFIDelCache(size)
	default:
		return newLRUFIDelCache(size)
	}
}

func newLRUFIDelCache(size uint32) *LRUFIDelCache {
	return &LRUFIDelCache{
		cache: NewLRUCache(size, nil),
	}
}
