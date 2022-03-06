package service

import (
	"mytest/CoustomerManagerSystem/model"
)

// CustomerService 该CustomerService ,完成对Customer的操作
//增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片有多少客户
	//该字段后，还可以作为新客户的id+1
	customerNum int
}

func NewCustomerService() *CustomerService {
	CustomerService := &CustomerService{}
	CustomerService.customerNum = 1
	var customer model.Customer = model.NewCustomer(1, "Siri", "男", 20, "13391110017", "siri@qq.com")
	CustomerService.customers = append(CustomerService.customers, customer)
	return CustomerService

}

// List 列出列表
func (c *CustomerService) List() []model.Customer {
	return c.customers
}

func NewCustomer(name string, gender string,
	age int, phone string, email string) model.Customer {
	return model.Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// Add 用户添加到customers
//!!!一定要指针，因为CustomerService永远是第一次创建的
func (c *CustomerService) Add(customer model.Customer) bool {
	//id顺序规则，即添加顺序
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}
