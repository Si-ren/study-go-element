package main

import (
	"fmt"
	"memCache/cache"
	"time"
)

func main() {
	cache0 := cache.NewMemCache()
	cache0.SetMaxMemory("100MB")
	cache0.Set("int", 1, 2*time.Second)
	cache0.Set("bool", false, 2*time.Second)
	cache0.Set("data", map[string]interface{}{"a": 1}, 2*time.Second)
	cache0.Get("int")
	cache0.Del("int")
	cache0.Flush()
	cache0.Keys()
	fmt.Println(cache0.UsedMemorySize)
	fmt.Println(cache0.Exists("data"))
	fmt.Println(cache0.Get("data"))

	time.Sleep(3 * time.Second)
	//cache1 := cache.NewMemCache()
	//cache1.SetMaxMemory("100MB")
	//cache1.Set("int", 2, 2*time.Second)
	//fmt.Println(cache1.Get("int"))
}
