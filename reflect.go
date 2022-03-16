package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	age  int
}

type Teacher struct {
	name string
	age  int
}

func reflectTestInt(a interface{}) {
	//通过反射获取传入的变量的type,kind,value
	//1.先获取到类型 reflct.Typeof
	rType := reflect.TypeOf(a)
	fmt.Println("rType=", rType)

	//获取传入变量的value
	rValue := reflect.ValueOf(a)
	fmt.Println("rValue=", rValue)

	//通过反射修改原本值
	rValue.Elem().SetInt(5)
	fmt.Println("rValue=", rValue)

	//获取变量对应的kind,kind可能和type一样,结构体的type是Student,但是kind是struct
	//rType.Kind()
	//rValue.Kind()
	kind1 := rType.Kind()
	kind2 := rValue.Kind()
	fmt.Printf("kind = %v, kind = %v\n", kind1, kind2)
	fmt.Printf("rValue=%v,rValue type=%T\n", rValue, rValue)

	switch a.(type) {
	case float32:
		fmt.Printf("参数类型是 float32 ,值为 %v\n", a)
	case int:
		fmt.Printf("参数类型是 int ,值为 %v\n", a)
	default:
		fmt.Printf("参数类型 不确定 ,值为 %v\n", a)
	}
}

func reflectTestStruct(a interface{}) {
	rType := reflect.TypeOf(a)
	fmt.Println("rType=", rType)
	rValue := reflect.ValueOf(a)
	fmt.Println("rValue=", rValue)
	fmt.Printf("rValue=%v,rValue type=%T\n", rValue, rValue)

	//直接转换,assert最佳实践
	switch a.(type) {
	case Student:
		stu := a.(Student)
		fmt.Printf("参数类型是 Student ,值为 %v\n", stu.name)
	case Teacher:
		tea := a.(Teacher)
		fmt.Printf("参数类型是 Teacher ,值为 %v\n", tea.name)
	default:
		fmt.Printf("参数类型 不确定 ,值为 %v\n", a)
	}
}

func main() {
	var num int = 1
	//通过反射传入指针修改值
	reflectTestInt(&num)
	fmt.Println(num)
	//stu := Student{"Siri",18}
	tea := Teacher{
		name: "LSL",
		age:  16,
	}
	reflectTestStruct(tea)
}
