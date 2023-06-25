package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
	cancelFunc()
	time.Sleep(time.Second)
	fmt.Println("main done")
}

func gen(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("func done")
				return
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
