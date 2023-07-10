package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

// 先运行程序，然后go tool trace  <trace.out文件> ，就会出现以下信息，用浏览器打开就可以了
// 2023/07/10 11:03:46 Parsing trace...
// 2023/07/10 11:03:46 Splitting trace...
// 2023/07/10 11:03:46 Opening browser. Trace viewer is listening on http://127.0.0.1:59692

func printStatus(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc: ", mem.Alloc)
	fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc)
	fmt.Println("mem.NumAlloc: ", mem.NumGC)
	fmt.Println("------")
}

func main() {
	// 首先需要创建一个文件，保存go tool trace实用程序的跟踪数据
	f, err := os.Create("C:\\Users\\sirius\\traceFile.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	// 进行跟踪处理
	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()
	var mem runtime.MemStats
	printStatus(mem)
	for i := 0; i < 3; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
		}
	}
	printStatus(mem)
	for i := 0; i < 5; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed")
		}
	}
	printStatus(mem)

}
