package service

import "mytest/CoustomerManagerSystem/model"

// CustomerService 该CustomerService ,完成对Customer的操作
//增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片有多少客户
	//该字段后，还可以作为新客户的id+1
	customerNum int
}
