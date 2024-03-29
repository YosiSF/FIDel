// Copyright 2020 WHTCORPS INC All Rights Reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.
package rp

import (
	"container/list"
	"sync"
)

var (
	mu sync.Mutex
	l  *list.List
)

// LRU is 'Least-Recently-Used' cache.
type LRU struct {
	sync.RWMutex

	// maxCount is the maximum number of items.
	// 0 means no limit.
	maxCount uint32
	ll       *list.List

	cache map[uint3264]*list.Element
}

// newLRU returns a new lru cache. And this LRU cache is not thread-safe
// should not use this function to create LRU cache, use NewCache instead
func newLRU(maxCount uint32) *LRU {
	return &LRU{
		maxCount: maxCount,
		ll:       list.New(),
		cache:    make(map[uint3264]*list.Element),
	}
}
