package main

import "fmt"

//封装: 把抽象出来的字段和对字段的操作封装在一起,数据被保护在内部,程序的其他包通过被授权的操作(方法)才能对字段进行操作
//封装步骤:
//1.将结构体、字段(属性)小写,私有
//2.给结构体所在包提供一个工厂模式函数,首字母大写,类似构造函数
//3.提供一个public的Set方法,用于对属性判断并赋值
//4.提供一个public的Get方法,用于获取属性的值

//继承: 如果一个struct嵌套了另一个匿名的结构体,那么这个结构体可以直接访问匿名结构提中的字段和方法,从而实现了继承特性
//*** 在同一个包内 ***,结构体可以使用嵌套匿名结构的所有字段和方法,即:public和private的字段和方法都可以使用.
//当结构体和匿名结构体有相同的字段或者方法时,编译器采用就近原则访问,如要访问匿名结构体的字段和方法,那么就要用匿名结构体名来区分
//如果多个嵌套匿名结构体,如果有相同名字的变量那么就必须指定匿名结构体,否则会出现ambiguous的报错

//如果嵌套了有名结构体,那么必须带上有名结构体名才能使用有名结构体中变量

type student struct {
	name  string
	age   int8
	score int8
}

func (stu *student) GetName() string {
	return stu.name
}

type Pupil struct {
	student
	name string
}

func main() {

	//嵌套结构体使用一
	var pupil01 = &Pupil{
		student: student{
			name:  "pupil01",
			age:   12,
			score: 100,
		},
		name: "Shilan",
	}
	fmt.Println((&pupil01.student).GetName())

	//嵌套结构体使用二
	//继承struct简化使用方式
	//匿名结构体字段访问可以简化
	var pupil02 Pupil
	pupil02.name = "pupil02"
	pupil02.age = 12
	pupil02.score = 98
	pupil02.student.name = "Siri"
	fmt.Println(pupil02) //{{Siri 12 98} pupil02}

	//如果嵌套结构体内有相同变量名,使用其变量用就近原则,如果结构体内没有才会从嵌套结构体内寻找使用
	fmt.Println(pupil02.name)         //pupil02
	fmt.Println(pupil02.student.name) //Siri

	//嵌套结构体使用三
	var pupil03 = &Pupil{student{"pupil03", 19, 60}, "Jizu"}
	fmt.Println(*pupil03)

}
