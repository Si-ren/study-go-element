package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// 基于协程的共享内存
// 基于go语言内置的同步特性
// 其他协程向持有共享数据的该单协程发送消息，防止数据受损。
// 此类协程被称作为监控协程
// 建议使用监控协程做数据共享，因为实现更加安全、整洁，更接近于go语言的哲学
var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Println(value)
		case readValue <- value:
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give an integer!")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Go to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()
	var w sync.WaitGroup
	for i := 0; i < 15; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			set(rand.Intn(n * 10))
		}()
	}
	w.Wait()
	fmt.Println("Last Value: ", read())
}
