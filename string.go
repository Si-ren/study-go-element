package main

import (
	"fmt"
	"strconv"
	"strings"
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

	//数字转换进制，返回对应的字符串
	tempStr = strconv.FormatInt(111, 2)
	fmt.Println(tempStr)
	tempStr = strconv.FormatInt(111, 16)
	fmt.Println(tempStr)

	//字符串中查找子字符串
	bl := strings.Contains(str, "hello")
	fmt.Println(bl)

	//判断字符串是否相等，"==" 区分大小写
	bl = strings.EqualFold("abc", "ABC")
	fmt.Println(bl)
	fmt.Println("abc" == "ABC")

	//字符串替换，数字n代表希望替换几个，-1表示全部替换
	var newStr string = strings.Replace(str, "hello", "byebye", 1)
	fmt.Println(newStr)

	//将字符串分割成字符串数组
	strArr := strings.Split(str, ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}

	//字符串全大小写
	//strings.ToLower()
	//strings.ToUpper()

	//去除字符串两边空格
	tempStr = strings.TrimSpace("  aaa bbb ccc   ")
	fmt.Println(tempStr)

	//将字符串左右两边指定的字符去掉
	//strings.Trim(str,"世界")  hello,
	//去除左边的 strings.TrimLeft() , strings.TrimRight()
	tempStr = strings.Trim(str, "世界")
	fmt.Println(tempStr)

	//判断字符串是否以指定的字符串开头
	//strings.HasPrefix("https://127.0.0.1","http") true
	//以字符串结尾  func HasSuffix(s string, suffix string) bool
	bl = strings.HasPrefix("https://127.0.0.1", "http")
	fmt.Println(bl)

}
