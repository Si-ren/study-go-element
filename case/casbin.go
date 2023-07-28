package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		logrus.Error(err)
	}

	sub := "zhangsan" // 想要访问资源的用户。
	obj := "data1"    // 将被访问的资源。
	act := "read"     // 用户对资源执行的操作。
	added, err := e.AddPolicy("zhangsan1", obj, act)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(added)
	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理err
	}

	if ok == true {
		// 允许alice读取data1
		logrus.Info("Successful")
	} else {
		// 拒绝请求，抛出异常
		logrus.Error("Error")
	}

	// 您可以使用BatchEnforce()来批量执行一些请求
	// 这个方法返回布尔切片，此切片的索引对应于二维数组的行索引。
	// 例如results[0] 是{"alice", "data1", "read"}的结果
	// results, err := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
}
