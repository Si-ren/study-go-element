package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 进程是个执行环境，包含指令、用户数据、部分系统数据，以及运行期内获取的其他资源类型
// 程序表示一个文件，文件包含了指令和数据，用于初始化指令和进程的用户数据部分
// 线程由进程创建，包含自身的控制流和栈
// 协程位于线程中
// go运行期也包含其自身的调度器，负责通过m:n调度机制执行协程。
// 通过多路复用技术，m个协程使用n个操作系统线程运行。
// *** go调度器是负责go程序中协程的执行方式和顺序的一个组件 ***
// 为什么协程会比线程快？
// 协程切换是完全在用户态进行，线程切换涉及特权模式切换，需要在内核空间完成
// 协程切换：
// 		1、只涉及基本的CPU上下文切换，把当前协程的 CPU 寄存器状态保存起来，然后将需要切换进来的协程的 CPU 寄存器状态加载的 CPU 寄存器上。
//		2、完全在用户态进行
// CPU上下文切换：一堆寄存器，里面保存了 CPU运行任务所需要的信息：从哪里开始运行（%rip：指令指针寄存器，标识 CPU 运行的下一条指令），栈顶的位置（%rsp： 是堆栈指针寄存器，通常会指向栈顶位置），当前栈帧在哪（%rbp 是栈帧指针，用于标识当前栈帧的起始位置）以及其它的CPU的中间状态或者结果（%rbx，%r12，%r13，%14，%15 等等）。
// 线程切换：
//		1、线程的调度只有拥有最高权限的内核空间才可以完成，所以线程的切换涉及到用户空间和内核空间的切换，也就是特权模式切换，然后需要操作系统调度模块完成线程调度（taskstruct），而且除了和协程相同基本的 CPU 上下文，还有线程私有的栈和寄存器等，说白了就是上下文比协程多一些，
//
//作者：暗淡了乌云
//链接：https://www.zhihu.com/question/308641794/answer/572499202

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
