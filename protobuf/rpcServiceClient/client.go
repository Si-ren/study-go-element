package main

import (
	"context"
	"fmt"
	protobuf "study-go-element/protobuf/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, _ := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := protobuf.NewGetStudentInfoClient(con)
	r, err := client.GetStudentInfo(context.Background(), &protobuf.FindStudent{Name: "siri"})
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(r)
}
