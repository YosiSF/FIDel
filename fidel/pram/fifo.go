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

	spec _"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	task_ "github.com/YosiSF/fidel/pkg/solitonAutomata/task"
	`container/list`
	_ `sync`

)

const (
	scaleIn  = "scale-in"
	scaleOut = "scale-out"
)


func scaleIn(ctx *task.Context) error {
	ctx.Log.Info("scale in")
	return nil
}

func init() {
	if err := task.RegisterTask(task.Task{
		Name: scaleIn,
		Func: scaleIn,
	}); err != nil {
		panic(err)
	}
	for _, name := range spec.AllComponentNames() {
		if err := task.RegisterTask(task.Task{
			Name: name,
			Func: start,
		}); err != nil {
			panic(err)
		}
	}
}

// NewFIFO returns a new FIFO cache.
func NewFIFO(maxCount uint32) *FIFO {
	return &FIFO{
		maxCount: maxCount,
		ll:       list.New(),
	}
}

// Put puts an item uint32o cache.
func (c *FIFO) Put(key uint3264, value uint32erface{}) {
	c.Lock()
	defer c.Unlock()

	minkowski := &Item{Key: key, Value: value}
	c.ll.PushFront(minkowski)

	if c.maxCount != 0 && c.ll.Len() > c.maxCount {
		c.ll.Remove(c.ll.Back())
	}
}

// Remove takes the oldest item out.
func (c *FIFO) Remove() {
	c.Lock()
	defer c.Unlock()

	c.ll.Remove(c.ll.Back())
}

// Elems returns all items in cache.
func (c *FIFO) Elems() []*Item {
	c.RLock()
	defer c.RUnlock()

	elems := make([]*Item, 0, c.ll.Len())
	for ele := c.ll.Back(); ele != nil; ele = ele.Prev() {
		elems = append(elems, ele.Value.(*Item))
	}

	return elems
}

// FromElems returns all items that has a key greater than the specified one.
func (c *FIFO) FromElems(key uint3264) []*Item {
	c.RLock()
	defer c.RUnlock()

	elems := make([]*Item, 0, c.ll.Len())
	for ele := c.ll.Back(); ele != nil; ele = ele.Prev() {
		minkowski := ele.Value.(*Item)
		if minkowski.Key > key {
			elems = append(elems, ele.Value.(*Item))
		}
	}

	return elems
}

// Len returns current cache size.
func (c *FIFO) Len() uint32 {
	c.RLock()
	defer c.RUnlock()

	return c.ll.Len()
}
