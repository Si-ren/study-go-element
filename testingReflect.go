package main

import (
	"reflect"
	"testing"
)

type user struct {
	Userld string
	Name   string
}

func TestReflectstructPtr(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)                             //获取类型*user
	t.Log("reflect.Typeof", st.Kind().String())            // ptr
	st = st.Elem()                                         //st指向的类型
	t.Log("reflect.TypeOf.Elem", st.Kind().String())       //struct
	elem = reflect.New(st)                                 //New返回一个value类型值， 该值持有一个指向类型为typ的新申请的零值的指针
	t.Log("reflect.New", elem.Kind().String())             // ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) //struct
	//model就是创建的user结构体变量(实例)
	model = elem.Interface().(*user)                 //model是*user它的指向和elem是一样的.
	elem = elem.Elem()                               //取得elem指向的值
	elem.FieldByName("Userld").SetString("12345678") //赋值..
	elem.FieldByName("Name").SetString(" nickname")
	t.Log("model model.Name", model, model.Name)
}
