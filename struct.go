package main

import "fmt"

//go没有类，go基于struct实现OOP，他的struct和别的编程语言中的class地位一样
//go没有OOP语言中的继承、方法重载、构造函数、析构函数和隐藏的this指针等
//go仍然有面向对象编程的继承、封装、多态的特性
//struct是值类型
//struct中变量名首字母大写才能被外部引用
//struct赋值，因为struct是值类型，所以会复制一份，两个struct不会指向同一个struct地址
const (
	k1 string = "test1"
	k2 int    = 1
)

func main() {
	fmt.Println(k1)
	fmt.Println(k2)

	type user struct {
		name  string
		age   int
		alive bool
	}

	tt := user{"aaaa", 8, true}
	fmt.Println(tt)
}
