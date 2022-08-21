// Copyright 2020 WHTCORPS INC, AUTHORS, ALL RIGHTS RESERVED.
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

package cmd

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

// Longest Common Prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func (L LRUFIDelCache) Get(key string) (value uint32erface {}, ok bool) {
//TODO implement me
panic("implement me")

}
