package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string //`json:"monster_name"` //反射机制，序列化后名称变为monster_name
	Age  int
	Sal  float64
}

//将结构体序列化
func testStruct() {
	monster := Monster{Name: "Siri", Age: 18, Sal: 2000.00}

	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("序列化错误 err:", err)
	}
	//Json序列化结果是一串字节码，所以需要string显示
	fmt.Println("monster序列化结果：", string(data))
}

//将map序列化
func testMap() {
	var a map[string]interface{}

	a = make(map[string]interface{})
	a["name"] = "Sherry"
	a["age"] = 10
	a["salary"] = 888888

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化错误 err:", err)
	}
	//Json序列化结果是一串字节码，所以需要string显示
	fmt.Println("monster序列化结果：", string(data))
}

//json反序列化
func unmarshalStruct() {
	str := "{\"Name\":\"Siri\",\"Age\":18,\"Sal\":2000}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化错误 err:", err)
	}
	fmt.Println("反序列化后monster：", monster)
}

func unmarshalMap() {
	str := "{\"age\":10,\"name\":\"Sherry\",\"salary\":888888}"
	var m map[string]interface{}
	//***这边map不需要make，申请空间了，在Unmarshal函数中已经做了
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("反序列化错误 err:", err)
	}
	fmt.Println("反序列化后monster：", m)
}

func main() {
	testStruct()
	testMap()
	unmarshalStruct()
	unmarshalMap()
}
