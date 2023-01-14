package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	protobuf "protobuf/gen/go"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, _ := grpc.Dial("grpc-server:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := protobuf.NewIPClient(con)
	for {
		r, err := client.GetIPInfo(context.Background(), &emptypb.Empty{})
		if err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Println(r)
		time.Sleep(1 * time.Second)
	}

}
