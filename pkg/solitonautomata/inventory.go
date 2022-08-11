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

package solitonautomata

import "time"

// IPFS with Ceph

type Ipfs struct {
	cache FIDelCache
}

func (d *Ipfs) Len() int {
	return d.cache.Len()
}

func (d *Ipfs) Del(key string) {
	d.cache.Del(key)
}

func (d *Ipfs) Clear() {
	d.cache.Clear()
}

// persistence of metadata
func (d *Ipfs) Get(key string) (value interface{}, ok bool) {
	for {
		value, ok = d.cache.Get(key)
		if ok {
			return value, ok
		}
		time.Sleep(time.Second)
	}
	//interlock with other goroutine
}

func (d *Ipfs) Set(key string, value interface{}) {
	d.cache.Set(key, value)
}

type Ceph struct {
	cache FIDelCache
}

// ceph is a cache for metadata
func (d *Ceph) Len() int {
	return d.cache.Len()
}

func (d *Ceph) Del(key string) {
	d.cache.Del(key)
}
