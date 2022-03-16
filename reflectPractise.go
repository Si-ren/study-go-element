package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:" name"`
	Age   int    `json:"monster age"`
	Score float32
	Sex   string
}

//方法，显示s的值
func (s Monster) Print() {
	fmt.Println("---start----")
	fmt.Println(s)
	fmt.Println("---end----")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	//获取类型
	typ := reflect.TypeOf(a)
	//获取值
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println(" expect struct")
		return
	}
	//获取到该结构体有多少字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值为=%v\n", i, val.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段有tag标签就显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}
	//获取结构体方法数量
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods \n", numOfMethod)
	// var params []reflect.value
	//按照函数名排序方法,这里使用的是第一个方法,从0开始
	val.Method(1).Call(nil)
	//调用结构体的第1个方法Methop(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果，返回的结果是[]reflect.Value
}

func main() {
	//创建了一个Monster实例
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	//将Monster实例传递给TestStruct函数
	TestStruct(a)
}
