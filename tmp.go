package main

import (
	"fmt"
	"sync"
)

func main() {
	var s sync.Map
	s.Store("map", &map[string]string{"testKey": "testValue"})
	fmt.Println(s.Load("map"))
}
