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
	"time"


	)

type ttlCacheItem struct {
	value  interface{}
	expire time.Time
}

// TTL is a cache that assigns TTL(Time-To-Live) for each items.
type TTL struct {
	sync.RWMutex
	ctx context.Context

	items      map[uint64]ttlCacheItem
	ttl        time.Duration
	gcInterval time.Duration
}
