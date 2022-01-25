package main

import (
	"fmt"
	"strconv"
)

var (
	a, b, c int64
)

func main() {
	a = 10
	var f float32 = float32(a)
	fmt.Printf("i=%T\n", f)
	b = 999999
	var i int16 = int16(b)
	fmt.Println(i)
	var str string = strconv.FormatInt(b, 2)
	fmt.Printf("str type is %T,%v", str, str)
}
