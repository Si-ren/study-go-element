package main

import (
	"fmt"
	"net"
)

func acceptProcess(con net.Conn) {
	//循环接收客户端发送的数据
	fmt.Println("Wait Client Send Info: ", con.RemoteAddr().String())
	defer con.Close()
	for {
		//创建一个切片存放接收的信息
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息
		//如果客户端没有Write,那么协程就阻塞
		n, err := con.Read(buf)
		if err != nil {
			//fmt.Println("Server Read err: ",err)
			fmt.Println("listen err: ", err)
			return
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Printf(string(buf[:n]))

	}

}
func main() {
	fmt.Println("服务器开始监听...")
	//表示使用的网络协议时tcp
	//0.0.0.0表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		//fmt.Println("listen err:",err)
		return
	}
	//延时关闭接口
	defer listen.Close()

	//循环等待客户端链接
	for {
		//等待客户端连接
		fmt.Println("等待客户端连接...")
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err: ", err)
		} else {
			fmt.Println("Accept() suc con= ", connect.RemoteAddr().String())
		}
		//准备一个协程为客户端服务
		go acceptProcess(connect)

	}

	fmt.Println("listen suc=", listen)
}
