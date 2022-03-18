package main

import (
	"fmt"
	redis "github.com/gomodule/redigo/redis"
)

//可以通过第三方包连接redis
//go get github.com/gomodule/redigo/redis
//拉下来的包在GOPATH中，可以通过go env查看gopath的位置

func main() {
	//连接到redis
	conn, err := redis.Dial("tcp", "192.168.137.141:6379")
	if err != nil {
		fmt.Println("redis.Dial err: ", err)
		return
	}

	//可以在redis中get name看到值为siri
	_, err = conn.Do("Set", "name", "siri")
	if err != nil {
		fmt.Println("conn.Do Set err: ", err)
	}

	//从redis获取值，并转换成字符串
	//如果返回值是int类型，那么就使用redis.Int
	name, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("conn.Do get err: ", err)
	}
	fmt.Println(name)

	//2.通过go向redis写 入数据string [key-va1]
	//批量写入
	_, err = conn.Do("HMSet", "user02", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("HMSet err=", err)
		return
	}

	//3.通过go向redis读取数据
	//批量读取，***注意这边是Strings，末尾有s
	r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	//redis批量读取返回的是类似切片的数据结构，所以得用for range遍历
	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

}
