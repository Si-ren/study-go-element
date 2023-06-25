package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	//使用互斥锁可以解决同时写问题
	lock sync.Mutex
)

func main() {
	//设置使用的CPU
	//1.8以下需要手动设置CPU核数
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}()
	wg.Wait()
	fmt.Println("Wait to Finish")

	fmt.Println("Complete Program ")

}
