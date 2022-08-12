package pram

import "time"

const (
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
func NewFIDelCache(size int, cacheType Type) FIDelCache {
	switch cacheType {
	case LRUFIDelCache:
		return NewLRUFIDelCache(size)
	case TwoQueueFIDelCache:
		return NewTwoQueueFIDelCache(size)
	default:
		return NewLRUFIDelCache(size)
	}
}
