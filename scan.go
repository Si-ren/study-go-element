package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var name string
	var age byte
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Printf("姓名:%v  年龄:%v", name, age)

}
