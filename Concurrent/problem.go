package main

import (
	"fmt"
	"runtime"
)

// 一些前置概念
// STW： stop the world，即停顿产生时整个应用程序线程都会暂停，没有任何响应，有点像卡死的感觉。
// 1.13之前协程之间调度都是非抢占式的。无函数调用的死循环 goroutine 会一直占据一个 P，GC 需要等待所有 goroutine 停止才得以执行，从而会导致 GC 延迟。
// 在go1.14后，引入了基于系统信号的异步抢占调度，即对占用执行权时间过长的协程发送信号通知，收到信号的G就主动让出CPU，自己重新进入P队列等待下次调度。

func main() {
	var i byte
	// 这里i的类型为byte，所以当到255时再+1，就溢出了，又为0，所以这里是死循环
	go func() {
		for i = 0; i <= 255; i++ {

		}
	}()
	fmt.Println("Leave goroutine")
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Good Bye")
}
