package main

import (
	"fmt"
	"runtime"
)

func main() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {

		}
	}()
	fmt.Println("Leave goroutine")
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Good Bye")
}
