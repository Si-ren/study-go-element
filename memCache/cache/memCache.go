package cache

import (
	"memCache/utils"
	"time"
)

var (
	_ Cache = (*memCache)(nil)
)

type memCache struct {
	// 最大内存大小
	MaxMemorySize uint64
	// 当前已使用内存
	UsedMemorySize uint64
}

func (c *memCache) SetMaxMemory(size string) bool {
	num, err := utils.ParseSize(size)
	if err != nil {
		return false
	}

	return true
}

func (c *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	return false
}

func (c *memCache) Get(key string) (interface{}, bool) {
	return nil, false
}

func (c *memCache) Del(key string) bool {
	return false
}

func (c *memCache) Exists(key string) bool {
	return false
}

func (c *memCache) Flush() bool {
	return false
}

func (c *memCache) Keys() int64 {
	return 0
}

func NewMemCache() *memCache {
	cache := &memCache{}
	return cache
}
