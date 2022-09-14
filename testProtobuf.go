package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	protobuf "study-go-element/protobuf"
)

func main() {
	student := protobuf.Student{Name: "siri", Sex: 1, Height: 173}
	fmt.Println(student)
	b, _ := proto.Marshal(&student)
	fmt.Println(b)
	fmt.Printf("%X\n", b)
	var stu1 protobuf.Student
	proto.Unmarshal(b, &stu1)
	fmt.Println(stu1)
}
