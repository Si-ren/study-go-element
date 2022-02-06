package main

import (
	"fmt"
)

// ***接口中所有的方法都没有方法体，即都是没有实现的方法
// ***接口中不能出现变量
// ***自定义类型能实现多个接口
// ***接口类型默认是一个指针(引用类型),如果没有对interface初始化就使用,那么会输出nil

//继承的价值: 解决代码复用性和可维护性
//接口的价值: 设计,设计好各种规范,让其他自定义类型去实现这些方法
//接口比继承更灵活,在一定程度上实现代码解耦

type Usb interface {
	Stop()
	Start()
}

// Machine ***接口继承
// ***要实现Machine方法,就要将Usb这个接口中的方法都实现
type Machine interface {
	Usb
	// MachineStartUp 接口本身的方法
	MachineStartUp()
}

// Phone 声明结构体
type Phone struct {
}

// Start Phone实现接口
func (p *Phone) Start() {
	fmt.Println("手机已连接...")
}

// Stop Phone实现接口
func (p *Phone) Stop() {
	fmt.Println("手机已拔出...")
}

// Camera test ctrl+F1  && Shift+Alt+Enter
type Camera struct {
}

func (c *Camera) Start() {
	fmt.Println("照相机已连接...")
}

// Stop Phone实现接口
func (c *Camera) Stop() {
	fmt.Println("照相机已拔出...")
}

type Computer struct {
}

// Working ***只要是实现了 Usb接口(指的是 实现了Usb接口声明***所有***方法)
func (c *Computer) Working(usb Usb) {
	//通过usb变量来调用方法
	usb.Start()
	usb.Stop()
}

func main() {
	var phone = &Phone{}
	var camera = &Camera{}
	var computer = &Computer{}
	computer.Working(phone)
	computer.Working(camera)

	// ***一个自定义类型只有实现了某个接口，才能将自定义类型的实例(变量)赋给接口类型
	var testUsb Usb = phone
	testUsb.Start() //手机已连接...
	testUsb.Stop()  //手机已拔出...

	fmt.Println("================")

	// 如果是空接口interface{},空接口没有任何方法,所以***所有类型***都实现了空接口

	type T interface {
	}
	var t T = *phone
	fmt.Println(t)

	var t2 interface{} = *camera
	fmt.Println(t2)

	// ***所有类型都实现了空接口***,所以可以把任何一个变量赋值给空接口
	var num01 int = 1
	t2 = num01
	fmt.Println(t2)

}
