package main

import (
	"fmt"
)

//go没有类，go基于struct实现OOP，他的struct和别的编程语言中的class地位一样
//go没有OOP语言中的继承、方法重载、构造函数、析构函数和隐藏的this指针等
//go仍然有面向对象编程的继承、封装、多态的特性
//struct是值类型
//*** struct中变量名首字母大写才能被外部引用
//*** struct赋值，因为struct是值类型，所以会复制一份，两个struct不会指向同一个struct地址

// User struct声明
type User struct {
	Name  string `json:"name,omitempty"` //``代表struct tag,该tag可以通过反射机制获取,常见的使用场景是序列化和反序列化
	Age   int    `json:"age,omitempty"`
	Alive bool   `json:"alive,omitempty"`
}

//函数与方法的区别
//如果是普通函数，接收者为值类型时候，不能将指针类型的数据直接传递，反之亦然
//如果是方法，接收者为值类型时，可以用指针类型的变量调用方法，反之亦然
//对于函数,func test(p *point){} ,不能传值类型:test(p),得这样用test(&p)

// PrintAge struct中的方法
//这里传值是值传递,如果是要引用传递,那么struct中要写成指针: func (user *User) PrintAge(){}
//在使用方法时,p.method和(&p).method都能使用,主要看定义方法时,使用的struct是值还是引用.
func (user User) PrintAge() {
	fmt.Println(user.Name, "的年龄为: ", user.Age)
}

//struct自定义String，这样使用 fmt.Println(&user04) 的输出为  Name: user04  Age: 20  Alive: true
func (user *User) String() string {
	var str string = fmt.Sprintf("Name: %v  Age: %v  Alive: %v", user.Name, user.Age, user.Alive)
	return str
}

//main
func main() {

	//struct使用方式一,直接声明使用
	user01 := User{"aaaa", 8, true}
	user01.Name = "user02"
	fmt.Println(user01)

	//struct使用方式二,返回的是指针类型!!!
	var user02 *User = new(User)
	//因为user02是一个指针,因此可以使用标准的字段赋值方式
	//不过go做了简化, (*user02).name = "user02" 也可以写成 user02.name="user02"
	//因为go的设计者,为了程序员使用方便,底层会对 user02.name="user02" 进行处理,会把user02加上取值运算符 (*user02).name = "user02"
	(*user02).Name = "user02"
	(*user02).Age = 11
	(*user02).Alive = true
	fmt.Println(*user02)

	//struct使用方式三,返回的是指针类型!!!
	var user03 *User = &User{"user03", 15, false}
	//(*user03).Name = "user03"
	//(*user03).Age = 15
	//(*user03).Alive = false
	fmt.Println(*user03)

	//struct使用方式四
	var user04 User
	user04 = User{"user04", 20, true}
	fmt.Println(user04)

	fmt.Println(&user04)
	user04.PrintAge()

}
