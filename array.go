package main

import "fmt"

func main() {
	//如果缺少了"..."，那么就是切片，例如[]int
	//长度是数据类型的一部分，所以 [2]int,和[3]int是不同的数据类型
	var numsArray01 = [...]int{1, 2, 3}
	fmt.Println(numsArray01)

	var numsArray02 [3]int = [3]int{4, 5, 6}
	fmt.Println(numsArray02)

	var numsArray03 = [3]int{7, 8, 9}
	fmt.Println(numsArray03)

	var numsArray04 [3]int = [...]int{1: 10, 0: 11, 2: 12}
	fmt.Println(numsArray04)

}
