package main

import (
	"fmt"
	"memCache/cache"
	"memCache/cache_server"
)

func main() {
	cache := cache.NewMemCache()
	cache.SetMaxMemory("100MB")
	cache.Set("int", 1)
	cache.Set("bool", false)
	cache.Set("data", map[string]interface{}{"a": 1})
	cache.Get("int")
	cache.Del("int")
	cache.Flush()
	cache.Keys()

	cache1 := cache_server.NewMemCache()
	cache1.SetMaxMemory("100MB")
	cache1.Set("int", 2)
	fmt.Println(cache1.Get("int"))
}
