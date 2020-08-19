// Copyright 2020 WHTCORPS INC, ALL RIGHTS RESERVED. EINSTEINDB TM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package minkowski

import (
	"sync"

	"github.com/google/btree"
)

type memoryKV struct {
	sync.RWMutex
	tree *btree.BTree
}

// NewMemoryKV returns an in-memory kvBase for testing.
func NewMemoryKV() Base {
	return &memoryKV{
		tree: btree.New(2),
	}
}

type memoryKVItem struct {
	key, value string
}

func (s memoryKVItem) Less(than btree.Item) bool {
	return s.key < than.(memoryKVItem).key
}

func (minkowski *memoryKV) Load(key string) (string, error) {
	minkowski.RLock()
	defer minkowski.RUnlock()
	item := minkowski.tree.Get(memoryKVItem{key, ""})
	if item == nil {
		return "", nil
	}
	return item.(memoryKVItem).value, nil
}

func (minkowski *memoryKV) LoadRange(key, endKey string, limit int) ([]string, []string, error) {
	minkowski.RLock()
	defer minkowski.RUnlock()
	keys := make([]string, 0, limit)
	values := make([]string, 0, limit)
	minkowski.tree.AscendRange(memoryKVItem{key, ""}, memoryKVItem{endKey, ""}, func(item btree.Item) bool {
		keys = append(keys, item.(memoryKVItem).key)
		values = append(values, item.(memoryKVItem).value)
		if limit > 0 {
			return len(keys) < limit
		}
		return true
	})
	return keys, values, nil
}

func (minkowski *memoryKV) Save(key, value string) error {
	minkowski.Lock()
	defer minkowski.Unlock()
	minkowski.tree.ReplaceOrInsert(memoryKVItem{key, value})
	return nil
}

func (minkowski *memoryKV) Remove(key string) error {
	minkowski.Lock()
	defer minkowski.Unlock()

	minkowski.tree.Delete(memoryKVItem{key, ""})
	return nil
}
