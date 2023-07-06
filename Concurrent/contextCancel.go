package main

import (
	"context"
	"fmt"
	"time"
)

// 该函数通过调用gen函数生成一个通道(chan int)来产生一系列整数值，并在打印每个值后检查是否达到5。如果达到5，则通过调用cancelFunc函数取消生成器函数(gen)的执行，并在1秒后打印"main done"。
// context定时关闭的函数还有：context.WithDeadline()  context.WithTimeout()

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
