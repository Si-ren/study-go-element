package main

import (
	"testing" //引入testing的框架包
)

//***文件名格式必须为XXXX_test.go

//***函数名必须以Test开头
//***形参类型必须是*testing.T,细节请看手册
//一个测试用例文件中，可以有多个测试用例函数
//go test -v ，加上-v表示不论运行正确或者错误都会输出日志

//***测试单个文件，一定要加上原文件
//go test .\testing.go .\testing_test.go -v
//=== RUN   TestAddUpper
//testing_test.go:23: AddUpper(10)执行正常...
//--- PASS: TestAddUpper (0.00s)
//PASS
//ok      command-line-arguments  0.034s

//测试单个方法 go test -v -test.run TestAddUpper

//要给测试用例，去测试addUpper是否正确
func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 45 {
		//输出错误
		t.Fatalf("AddUpper(10)执行错误，期望值：55，实际值：%v\n", res)
	}
	//输出日志
	t.Logf("AddUpper(10)执行正常...")
}
