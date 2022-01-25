package main

import "fmt"

func GetSum(n1 int, n2 int) int {
	return n1 + n2
}

func TestFuncVar(funcvar func(int, int) int, n1 int, n2 int) int {
	return funcvar(n1, n2)
}
func Args(args ...int) int {
	var sum int = 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
func main() {
	a := GetSum
	var n1 int = 1
	var n2 int = 2
	var n3 int = a(n1, n2)
	fmt.Println(n3)
	var n4 int = TestFuncVar(GetSum, n1, n2)
	fmt.Println(n4)

	n5 := Args(1, 1, 1, 1, 1, 1, -1)
	fmt.Println(n5)
}
