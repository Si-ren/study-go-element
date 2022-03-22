package main

import (
	"errors"
	"fmt"
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

func main() {
	testError()

	err := tryError("file")
	//中断程序并打印错误
	panic(err)
	//fmt.Println(err)
}
