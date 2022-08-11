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

type FIDelCache interface {
	Get(key string) (value interface{}, ok bool)
	Set(key string, value interface{})
	Del(key string)
	Len() int
	Cap() int
	Clear()
}

type LRUFIDelCache struct {
	capacity int
}

func (L LRUFIDelCache) Get(key string) (value interface{}, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (L LRUFIDelCache) Set(key string, value interface{}) {
	//TODO implement me
	panic("implement me")
}

type Dmapi struct {
	cache FIDelCache
}

func (d *Dmapi) Len() int {
	return d.cache.Len()
}

func (d *Dmapi) Del(key string) {
	d.cache.Del(key)
}

func (d *Dmapi) Clear() {
	d.cache.Clear()
}

func (d *Dmapi) Get(key string) (value interface{}, ok bool) {
	return d.cache.Get(key), false
}

func (d *Dmapi) Set(key string, value interface{}) {
	d.cache.Set(key, value)
}
