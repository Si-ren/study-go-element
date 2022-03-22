package main

import (
	"fmt"
	"mytest/chatRoom/common/message"
	"mytest/chatRoom/my_server/utils"
	"net"
)

var (
	transfer = utils.Transfer{}
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()
	if err != nil {
		fmt.Println("net Listen err: ", err)
	}

	for {
		//等待客户端连接
		fmt.Println("等待客户端连接...")

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen Accept() err: ", err)
		} else {
			fmt.Println("Accept() suc con= ", conn.RemoteAddr().String())
		}
		//准备一个协程为客户端服务
		go processMessage(conn)

	}
}

func processMessage(conn net.Conn) {
	transfer.Conn = conn
	mes, err := transfer.ReadPkg()
	if err != nil {
		return
	}
	switch mes.Type {

	}
}
