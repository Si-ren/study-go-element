package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//os.Args是一个string的切片，用来存储所有的命令行参数
	fmt.Println("命令行参数：", len(os.Args))
	//遍历Args
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("args[%v] = %v\t", i, os.Args[i])
	}
	fmt.Println()

	//用flag包来解析命令行参数
	var user string
	var password string
	var host string
	//如果参数为int类型，有IntVar的方法
	//方法参数1为接收变量，参数2为命令行中“-u”等，参数3为默认值，参数4为用法
	flag.StringVar(&user, "u", "", "用户名，不能为空")
	flag.StringVar(&password, "p", "", "密码，不能为空")
	flag.StringVar(&host, "h", "localhost", "url，不能为空")

	//转换结果
	flag.Parse()

	//输出
	fmt.Println("user: ", user)
	fmt.Println("password: ", password)
	fmt.Println("host: ", host)

}
