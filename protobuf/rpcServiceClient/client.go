package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	protobuf "protobuf/gen/go"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial("hsc-wechat-hololens-installation-service:9898", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Error("Can't connect rpc server: ", err)
	}
	logrus.Info("Connect rpc server success")
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
