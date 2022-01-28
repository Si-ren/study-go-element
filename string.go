package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "hello,世界"

	//遍历string
	r := []rune(str)
	fmt.Println(str)
	fmt.Println(r)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//字符串转整数,例如 "123" --> 123
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("Success: ", n)
	}

	//整数转字符串
	numToStr := strconv.Itoa(123)
	fmt.Printf("numToStr Type is %T, numToStr is %s\n", numToStr, numToStr)

	//字符串转byte
	var bytes = []byte(str)
	fmt.Printf("bytes=%v\n", bytes)

	//byte转字符串
	tempStr := string([]byte{97, 98, 99})
	fmt.Println(tempStr)
}
