package main

import (
	"fmt"
	"math/rand"
	"time"
)

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")

	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Printf("C()! %d \n", rand.Intn(10))
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})
	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)
	//此处先执行A，因为先关闭的是 x channel，只有A 传参了x。
	//最后多次调用C不会报错原因，因为C函数不会关闭任何通道。

	close(x)
	time.Sleep(2 * time.Second)
	/*
	   输出为：
	      A()!
	      B()!
	      C()!
	      C()!
	      C()!
	*/
}
