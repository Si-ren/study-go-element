package main

import (
	"context"
	"log"
	"net"
	"net/http"
	protobuf "study-go-element/protobuf/gen/go"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

type Service struct{}

func (s *Service) GetStudentInfo(context.Context, *protobuf.FindStudent) (*protobuf.Student, error) {
	return &protobuf.Student{Name: "siri", Sex: 1, Height: 173}, nil
}

func (s *Service) mustEmbedUnimplementedGetStudentInfoServer() {}
func main() {
	go startGRPCGateway()
	lis, _ := net.Listen("tcp", ":8081")
	s := grpc.NewServer()
	protobuf.RegisterGetStudentInfoServer(s, &Service{})
	log.Fatal(s.Serve(lis))
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			protojson.MarshalOptions{
				AllowPartial:    true,
				EmitUnpopulated: true,
			},
			protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	))
	err := protobuf.RegisterGetStudentInfoHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("error registering getStudentInfoHandler: %v", err)
	}
	err = http.ListenAndServe(":8082", mux)
}
