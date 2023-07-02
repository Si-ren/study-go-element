package main

import "fmt"

//channel本质是一个数据结构-队列
//线程安全
//channel是有类型的
//管道一定要有读写,如果缺少读或者写,那么就会报错deadlock
// struct{}信号通道有点是不可向其中发送数据，从而防止bug和错误的出现
//如果从关闭的通道中读取数据，那么将会得到该通道类型的零值

func main() {
	// new和make有啥区别
	// new 用于分配值类型的内存，并返回指向该类型零值的指针；而 make 用于分配引用类型的内存，并进行初始化，返回该类型的引用。

	//创建一个存放1个int的管道
	//如果要传入多个类型,那么就channel的类型就必须为interface{}
	//interface类型的channel取出时要进行类型断言,否则取出的类型为interface不能使用
	var court = make(chan int, 1)

	//存放一个数据
	court <- 1

	//如果管道已满,多存一个会报错,fatal error: all goroutines are asleep - deadlock!
	//court <- 2

	//go func() {
	//	ball,ok := <-court
	//	fmt.Println(ball,ok)
	//}()

	//取出一个数据
	//如果管道数据全部取出,再取就会报错,报告deadlock
	//<- court表示取出丢掉
	ball, ok := <-court
	close(court)

	ball2, ok2 := <-court
	fmt.Println(ball, ok)
	fmt.Println(ball2, ok2)

	var intChan = make(chan int, 100)
	for i := 0; i <= 10; i++ {
		intChan <- i
	}
	//遍历管道
	//不能使用管道的容量去遍历,取一个少一个,容量会大于存在的数据,会下标越界
	//在遍历时,如果channel没有关闭,则会出现报错-->fatal error: all goroutines are asleep - deadlock!
	//因为没有关闭时,程序会认为可能有数据持续写入,因此就会等待,没有数据写入就会死锁了.
	//如果遍历时,管道已经关闭,则会正常遍历数据,遍历完成后,就会退出遍历
	close(intChan)
	for v := range intChan {
		fmt.Println("v=", v)
	}
}
