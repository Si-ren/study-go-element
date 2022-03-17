package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connect, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Dial err: ", err)
	}
	fmt.Println("Dial sucessful: ", connect)
	//客户端发送单行数据
	reader := bufio.NewReader(os.Stdin)
	for {
		//从终端读取一行用户输入，并发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Reader err: ", err)
		}
		//如果输入exit,那么就退出客户端
		if line == "exit\n" {
			break
		}
		n, err := connect.Write([]byte(line))
		if err != nil {
			fmt.Println("Write err: ", err)
		}
		fmt.Printf("客户端发送了%v字节\n", n)
	}

}
