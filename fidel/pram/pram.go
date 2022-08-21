package rp

import (
	proto _ "github.com/djbarber/ipfs-hack/Godeps/_workspace/src/github.com/gogo/protobuf/proto"
	"github.com/djbarber/ipfs-hack/Godeps/_workspace/src/golang.org/x/net/context"

	mdag "github.com/djbarber/ipfs-hack/merkledag"
	ft "github.com/djbarber/ipfs-hack/unixfs"
	ftpb "github.com/djbarber/ipfs-hack/unixfs/pb"
	"time"
)

const (
	// LRUFIDelCache is the type of LRU FIDelCache
	LRUFIDelCache = "LRU"

	// DefaultFIDelCacheType is the default cache type.
	DefaultFIDelCacheType = "lru"

	// DefaultFIDelCacheSize is the default cache size.
	DefaultFIDelCacheSize = 1024

	// DefaultFIDelCacheCapacity is the default cache capacity.
	DefaultFIDelCacheCapacity = 1024

	// DefaultFIDelCacheExpired is the default cache expired.
	DefaultFIDelCacheExpired = time.Second * 10

	// DefaultFIDelCacheEvicted is the default cache evicted.
	DefaultFIDelCacheEvicted = time.Second * 10
)

// NewFIDelCache create FIDelCache instance by cache type
func NewFIDelCache(size uint32, cacheType Type) FIDelCache {
	switch cacheType {
	case LRUFIDelCache:
		return NewLRUFIDelCache(size)
	case TwoQueueFIDelCache:
		return NewTwoQueueFIDelCache(size)
	default:
		return NewLRUFIDelCache(size)
	}
}
