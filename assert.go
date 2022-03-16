package main

import "fmt"

type Stu struct {
}

// TypeJudge 断言的最佳实践
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case float32:
			fmt.Printf("第%v个参数类型是 float32 ,值为 %v\n", index, x)
		case int:
			fmt.Printf("第%v个参数类型是 int ,值为 %v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数类型是 float64 ,值为 %v\n", index, x)
		case bool:
			fmt.Printf("第%v个参数类型是 bool ,值为 %v\n", index, x)
		case string:
			fmt.Printf("第%v个参数类型是 string ,值为 %v\n", index, x)
		case Stu:
			fmt.Printf("第%v个参数类型是 Stu ,值为 %v\n", index, x)
		case *Stu:
			fmt.Printf("第%v个参数类型是 *Stu ,值为 %v\n", index, x)
		default:
			fmt.Printf("第%v个参数类型 不确定 ,值为 %v\n", index, x)
		}
	}
}

func main() {
	var x interface{}
	var a float64 = 1.1
	x = a //空接口接收任意类型的值

	//使用类型断言,x已被赋为float64类型,如果断言为别的类型,那么就会报错
	y, err := x.(float64)
	fmt.Println(y, err) //1.1 true
	fmt.Println("============")

	var n1 int = 1
	var n2 float32 = 2.2
	var n3 string = "33"
	var n4 bool = true
	var n5 Stu = Stu{}
	var n6 *Stu = &Stu{}
	TypeJudge(n1, n2, n3, n4, n5, n6)

}
