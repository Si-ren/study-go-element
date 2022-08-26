package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func testError() {
	//使用defer + recover 获取异常，但不会退出程序
	defer func() {
		//recover内置函数，捕获到异常
		if err := recover(); err != nil {
			fmt.Println("ERROR: ", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func tryError(str string) (err error) {
	if str == "file1" {
		return nil
	} else {
		//返回自定义错误
		return errors.New("文件错误。。")
	}
}

func func1() error {
	return errors.New("this is func1 error")
}

// errors.Wrap 和 WithMessage可以传递error的上下文
func func2() error {
	err := func1()
	//return errors.Wrap(err,"this is func2 error")
	return errors.WithMessage(err, "this is func3 error")

}

func func3() error {
	err := func2()
	//return errors.Wrap(err,"this is func3 error")
	return errors.WithMessage(err, "this is func3 error")
}

// 协程处理发生error 不中断main函数
func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		x()
	}()
}

func main() {
	//testError()
	//err := tryError("file")
	//中断程序并打印错误

	Go(func() {
		fmt.Println("This is go routine")
		panic("go routine err")
	})

	err := func3()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		os.Exit(1)
	}

	panic(err)
	//fmt.Println(err)
}
