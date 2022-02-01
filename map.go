package main

import (
	"fmt"
	"os"
)

func main() {
	//map的声明和注意事项
	//使用map前，需要先make，make的作用就是map分配数据空间
	//map的value经常是struct类型，更适合复杂的数据类型
	//第一种map使用方式
	var map01 map[string]string
	map01 = make(map[string]string, 10)
	map01["key01"] = "value01"
	map01["key02"] = "value02"
	map01["key07"] = "value07"
	map01["key08"] = "value08"
	//查找一个值，如果值存在，那么exists变量为true，不存在则为false
	val, exists := map01["key01"]
	fmt.Println("key01对应的值：", val, "exists对应的值: ", exists)
	//map删除一个key
	//如果key不在，那么delete也不会报错，也不会删除
	delete(map01, "key01")
	fmt.Println("map01的元素有: ", map01)
	//遍历map，需要使用for_range
	for k, v := range map01 {
		fmt.Println("key为: ", k, "  value为: ", v)
	}
	//如果要排序，那么就得新建一个slice或者数组，或者别的数据结构，在那儿key排序好后，再用for循环输出value

	//第二种map使用方式
	map02 := map[string]string{
		"key03": "value03",
		"key04": "value04",
	}
	fmt.Println(map02)

	//第三种map使用方式
	map03 := make(map[string]string)
	map03["key05"] = "value05"
	map03["key06"] = "value06"
	fmt.Println(map03)

	var arrStu [3]map[string]string

	arrStu[0] = make(map[string]string, 2)
	arrStu[0]["name"] = "XiRui Wong"
	arrStu[0]["sex"] = "男"
	arrStu[1] = make(map[string]string, 2)
	arrStu[1]["name"] = "ShiLan Luo"
	arrStu[1]["sex"] = "女"
	arrStu[2] = make(map[string]string, 2)
	arrStu[2]["name"] = "Jue Luo"
	arrStu[2]["sex"] = "女"
	fmt.Println(arrStu)

	os.Exit(2)
}
