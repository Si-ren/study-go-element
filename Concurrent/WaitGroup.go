package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 10
	var waitGroup sync.WaitGroup
	fmt.Println("waitGroup Start")
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			fmt.Println(i)
		}(i)
	}
	fmt.Println("waitGroup wait")
	waitGroup.Wait()
	fmt.Println("waitGroup Done")
}
