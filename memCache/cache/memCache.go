package cache

import (
	"fmt"
	"memCache/utils"
	"sync"
	"time"
	"unsafe"
)

var (
	_ Cache = (*memCache)(nil)
)

type memCache struct {
	// 最大内存大小
	MaxMemorySize    uint64
	MaxMemorySizeStr string
	// 当前已使用内存
	UsedMemorySize uint64

	values sync.Map
}

type memCacheValue struct {
	val        interface{}
	size       uint
	expireTime time.Time
}

func (c *memCache) SetMaxMemory(size string) bool {
	num, err := utils.ParseSize(size)
	if err != nil {
		return false
	}
	c.MaxMemorySize = num
	c.MaxMemorySizeStr = size
	return true
}

func (c *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	v := &memCacheValue{
		val:        val,
		size:       uint(unsafe.Sizeof(val)),
		expireTime: time.Now().Add(expire),
	}
	c.UsedMemorySize = c.UsedMemorySize + uint64(unsafe.Sizeof(val))
	c.values.Store(key, v)
	fmt.Printf("Store key: %s  value: %s\n", key, val)
	return true
}

func (c *memCache) Get(key string) (interface{}, bool) {
	v, ok := c.values.Load(key)
	fmt.Printf("this is key: %s %#v\n", key, v)

	if !ok {
		fmt.Println("Get value error")
		return nil, false
	}

	value, _ := v.(*memCacheValue)
	if time.Now().Before(value.expireTime) {
		return v, true

	}
	fmt.Println("key expired")
	return nil, false
}

func (c *memCache) Del(key string) bool {
	v, ok := c.values.Load(key)
	if !ok {
		fmt.Println("key not exist")
		return false
	}
	c.UsedMemorySize = c.UsedMemorySize - uint64(v.(*memCacheValue).size)
	c.values.Delete(key)
	return false
}

func (c *memCache) Exists(key string) bool {
	_, ok := c.values.Load(key)
	if !ok {
		fmt.Printf("key:%s not exists \n", key)
	}
	return true
}

func (c *memCache) Flush() bool {
	c.values.Range(func(key, value any) bool {
		c.values.Delete(key)
		return true
	})
	return true
}

func (c *memCache) Keys() uint64 {
	var i uint64
	c.values.Range(func(key, value any) bool {
		i++
		return true
	})
	return i
}

func NewMemCache() *memCache {
	cache := &memCache{}
	return cache
}
