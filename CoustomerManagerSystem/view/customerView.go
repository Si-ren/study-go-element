package main

import (
	"fmt"
	"mytest/CoustomerManagerSystem/service"
)

type customerView struct {
	//定义必要字段
	key string
	//表示是否循环显示菜单
	loop bool

	customerService *service.CustomerService
}

func (receiver *customerView) list() {
	customers := receiver.customerService.List()
	fmt.Println("----------客户列表----------")
	fmt.Printf("编号\t姓名\t性别\t年龄\t电话\t\t邮箱\n")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("---------客户列表完成---------")
}

func (receiver *customerView) add() {
	fmt.Println("----------添加客户----------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	customer := service.NewCustomer(name, gender, age, phone, email)
	if receiver.customerService.Add(customer) {
		fmt.Println("----------添加完成----------")
	}
}

func (receiver *customerView) mainMenu() {
	for {
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("          1 添加客户")
		fmt.Println("          2 修改客户")
		fmt.Println("          3 删除客户")
		fmt.Println("          4 客户列表")
		fmt.Println("          5 退出")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&receiver.key)
		switch receiver.key {
		case "1":
			receiver.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			//list函数是此对象的函数，所以要用this，不能使用结构体名
			receiver.list()
		case "5":
			fmt.Println("退出")
			receiver.loop = false
		default:
			fmt.Println("输入有误，请重新输入...")
		}
		if !receiver.loop {
			fmt.Println("你退出了客户关系管理系统")
			break
		}

	}
}

func main() {
	//在主函数中创建一个CustomerView并运行显示主菜单..
	customerView := customerView{key: "", loop: true}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
