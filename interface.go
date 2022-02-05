package main

import (
	"fmt"
)

type Usb interface {
	Stop()
	Start()
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

// Working 只要是实现了 Usb接口(指的是 实现了Usb接口声明***所有***方法)
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

}
