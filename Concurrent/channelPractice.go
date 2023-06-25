package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 这边有个疑问？为啥这个要全局变量
// ***map不支持并发修改***如果要使用map并发修改，那么就需要使用sync.map

//使用for select时，单独的break不能退出for循环，得使用break LOOP:的方式退出
//LOOP:
//	for{
//		select{
//		case <- ch1:
//			break LOOP
//		case <- ch2:
//			break LOOP
//		case default:
//			break LOOP
//		}
//	}

var (
	wg sync.WaitGroup
)

// intChan存放要计算的数
func putNum(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 800; i++ {
		intChan <- i
	}
	close(intChan)
}

// 计算素数并把素数放入primeChan管道中
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	defer wg.Done()
	fmt.Println("开始数据筛选")
	var flag bool
	for {
		flag = true
		num, ok := <-intChan
		if !ok {
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}

	exitChan <- true
}

func main() {
	runtime.GOMAXPROCS(4)
	wg.Add(5)
	intChan := make(chan int, 800)
	//放入素数的结果
	primeChan := make(chan int, 200)
	exitChan := make(chan bool, 4)
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	fmt.Println("数据筛选完成")
	go func() {
		fmt.Println("等带筛选结束")
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
		close(exitChan)
	}()

	wg.Wait()

	fmt.Println("开始打印素数")
	for i := range primeChan {
		fmt.Println(i)
	}
	wg.Wait()
	fmt.Println("查找素数结束")

}
