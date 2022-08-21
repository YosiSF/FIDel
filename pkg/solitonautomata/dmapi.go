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

type FIDelCache uint32erface {
Get(key string) (value uint32erface{}, ok bool)
Set(key string, value uint32erface{})
Del(key string)
Len() uint32
Cap() uint32
Clear()
}
type LRUFIDelCache struct {
	capacity uint32
}

func (L LRUFIDelCache) Get(key string) (value uint32erface {}, ok bool) {
//TODO implement me
panic("implement me")
}

func (L LRUFIDelCache) Set(key string, value uint32erface {}) {
//TODO implement me
panic("implement me")
}
type Dmapi struct {
	cache FIDelCache
}

func (d *Dmapi) Len() uint32 {
	return d.cache.Len()
}

func (d *Dmapi) Del(key string) {
	d.cache.Del(key)
}

func (d *Dmapi) Clear() {
	d.cache.Clear()
}

func (d *Dmapi) Get(key string) (value uint32erface {}, ok bool) {
return d.cache.Get(key), false
}

func (d *Dmapi) Set(key string, value uint32erface{}) {
	d.cache.Set(key, value)
}
