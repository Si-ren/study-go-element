package main

import (
	"fmt"
	pattern "mytest/pattern"
)

func main() {
	var factory = pattern.NewFactory("factory01", 10, true)
	//使用包中小写开头的struct
	fmt.Println(*factory)
	//获取struct中小写开头的成员变量
	fmt.Println(factory.GetUse())
}
