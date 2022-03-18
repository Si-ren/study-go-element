package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//当启 动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示和数据库的最大链接数， 0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接的代码，链接哪个ip的redis
			return redis.Dial("tcp", "192.168.137.141:6379")
		},
	}
}

func main() {
	//先从pool取出一个链接
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", "name", "Tom~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	//取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)

	//一定保证取redis连接时，连接池未关闭
}
