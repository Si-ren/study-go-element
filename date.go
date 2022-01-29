package main

import (
	"fmt"
	"time"
)

func main() {
	//看时间和日期相关函数
	now := time.Now()
	fmt.Printf("%v\n%T\n", now, now)

	//通过now可以获取到的年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//直接通过函数获得
	fmt.Println(time.Now().Year())

	//格式化日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateStr)

	//格式化日期时间，字符串年月日中只能写2006-01-02，必须这么写
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	//时间常量
	//https://pkg.go.dev/time#pkg-constants
	fmt.Println(time.Millisecond)

	//Unix和UnixNano的使用
	fmt.Printf("Unix时间戳: %v\nUnixnano的时间戳: %v\n", now.Unix(), now.UnixNano())
}
