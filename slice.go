package main

import (
	"fmt"
	"os"
)

func main() {
	//切片时引用类型，所以在函数传递时，是引用传递
	//用make创建切片，[]string代表类型，5代表长度，其实还有第三个参数，capacity容量，如果指定了容量，那么一定cap>=len
	//通过make创建的切片对应的数组是由make底层维护，对外不可见，即只能通过slice去访问各个元素
	//切片初始化时，仍然不能越界。范围在[0-len(arr)]之间，但是可以动态增长
	//var slice = arr[0:end] 可以简写: var slice = arr[:end]
	//var slice = arr[start:len(arr)] 可以简写: var slice = arr[start:]
	//var slice = arr[0:len(arr)] 可以简写: var slice = arr[:]
	slice01 := make([]string, 5)

	slice01[1] = "AAAAA"
	slice01[2] = "BBBBB"
	slice01[3] = "cccccc"
	slice01[4] = "ddddd"

	fmt.Println(slice01)
	fmt.Println("切片的地址: ", &slice01)
	fmt.Println("切片第一个数据的地址: ", &slice01[0])
	fmt.Println("切片最后一个数据的地址: ", &slice01[4])
	fmt.Println()
	//切片是左闭右开，所以"2:3"的意思是[2,3),第三个3代表容量
	slice02 := slice01[2:3:3]
	slice02[0] = "test"
	fmt.Println("slice01的数据：", slice01)
	fmt.Println("slice02的数据", slice02)

	//当slice使用append扩容，底层会把原slice复制一份
	//append本质就是对数组进行扩容
	//append会创建一个扩容后的newArray
	//将slice包含的元素copy到newArray
	//slice重新引用newArray
	slice02 = append(slice02, "11111")
	fmt.Println("slice01的数据：", slice01)
	fmt.Println("slice02的数据", slice02)
	slice02[0] = "testAppend"
	fmt.Println("slice01的数据：", slice01)
	fmt.Println("slice02的数据", slice02)

	fmt.Println(len(slice01))
	fmt.Println(len(slice02))
	fmt.Println(cap(slice01))
	fmt.Println(cap(slice02))

	//定义一个切片，直接指定具体数组，原理类似make
	var slice03 []string = []string{"111", "222", "333"}
	fmt.Println(slice03)
	fmt.Println("slice03的长度", len(slice03))
	fmt.Println("slice03的容量", cap(slice03))
	fmt.Println()

	//copy()内置函数，参数只能为切片
	//切片为相对独立的空间
	slice04 := []int{1, 2, 3, 4, 5}
	slice05 := make([]int, 10)
	copy(slice05, slice04)
	fmt.Println("slice05: ", slice05)
	/*
		for _, value := range slice{
			fmt.Printf("Value: %s\n", value)
		}
		for _, value := range slice02{
			fmt.Printf("Value: %s\n", value)
		}
	*/
	os.Exit(2)
}
