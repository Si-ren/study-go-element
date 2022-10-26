package service

import (
	"mytest/CoustomerManagerSystem/model"
)

// CustomerService 该CustomerService ,完成对Customer的操作
// 增删改查
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

func (c *CustomerService) Delete(id int) bool {
	index := c.FindById(id)
	//如果index为-1，说明没有这用户
	if index == -1 {
		return false
	}
	//从切片删除一个元素,即元素的左边加上元素的右边，跳过这个元素
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

// FindById 根据id查找客户在切片中的位置
func (c *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(c.customers); i++ {
		if c.customers[i].Id == id {
			index = i
		}
	}

	return index
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
// !!!一定要指针，因为CustomerService永远是第一次创建的
func (c *CustomerService) Add(customer model.Customer) bool {
	//id顺序规则，即添加顺序
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}
