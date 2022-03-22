package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"mytest/chatRoom/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8000]byte
}

func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {
	//读取客户端发来的数据长度
	//由于int32占4字节,所以读取4字节
	_, err = transfer.Conn.Read(transfer.Buf[0:4])
	if err != nil {
		fmt.Println("Conn Read err: ", err)
		return
	}
	//把四字节转换为int32
	pkglen := binary.BigEndian.Uint32(transfer.Buf[:4])

	//读取pkglen个字节
	_, err = transfer.Conn.Read(transfer.Buf[:pkglen])
	if err != nil {
		fmt.Println("Conn Read err: ", err)
		return
	}
	//由于传输过来的是json,所以需要把数据转序列化
	err = json.Unmarshal(transfer.Buf[:pkglen], &mes)
	if err != nil {
		fmt.Println("Json Unmarshal err: ", err)
		return
	}
	return
}

func (transfer *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度,需要把uint32,转换为字节
	var pkglen uint32
	pkglen = uint32(len(data))
	binary.BigEndian.PutUint32(transfer.Buf[0:4], pkglen)
	_, err = transfer.Conn.Write(transfer.Buf[0:4])
	if err != nil {
		fmt.Println("Con Write err: ", err)
		return
	}

	//再发送数据
	_, err = transfer.Conn.Write(data)
	if err != nil {
		fmt.Println("Con Write err: ", err)
		return
	}
	return
}
