package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// 官网安装所需工具：https://graphviz.gitlab.io/_pages/Download/Download_windows.html
// go  tool pprof  /tmp/cpuProfile.txt //进入pprof
// top10 --cum //查看每个函数的累计时间
// list main.main  //list <包名>.<函数名>查看某个特定函数所发生的情况，可以获得函数的详细信息
// pdf 生成一个好看的pdf图
func fibo1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Millisecond)
	return int64(fibo1(n-1)) + int64(fibo1(n-2))
}

func fibo2(n int) int {
	fn := make(map[int]int)
	for i := 0; i < n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	time.Sleep(50 * time.Millisecond)
	return fn[n-1]
}

func main() {
	cpufile, err := os.Create("/tmp/cpuProfile.txt")
	if err != nil {
		logrus.Error(err)
		return
	}
	pprof.StartCPUProfile(cpufile)
	defer pprof.StopCPUProfile()

	for i := 0; i < 50; i++ {
		n := fibo2(i)
		fmt.Print(n, " ")
	}
	fmt.Println()
	for i := 0; i < 50; i++ {
		n := fibo2(i)
		fmt.Print(n, " ")

	}
	fmt.Println()
	runtime.GC()

	memory, err := os.Create("/tmp/mempryProfile.txt")
	for i := 0; i < 50; i++ {
		s := make([]byte, 50000)
		if s == nil {
			fmt.Println("Operation failed")

		}
		time.Sleep(50 * time.Millisecond)

	}
	err = pprof.WriteHeapProfile(memory)
	if err != nil {
		fmt.Println(err)
		return
	}

}
