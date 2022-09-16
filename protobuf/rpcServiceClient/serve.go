package main

import (
	"context"
	"log"
	"net"
	protobuf "study-go-element/protobuf/gen/go"

	"google.golang.org/grpc"
)

type Service struct{}

func (s *Service) GetStudentInfo(context.Context, *protobuf.FindStudent) (*protobuf.Student, error) {
	return &protobuf.Student{Name: "siri", Sex: 1, Height: 173}, nil
}

func (s *Service) mustEmbedUnimplementedGetStudentInfoServer() {}
func main() {
	lis, _ := net.Listen("tcp", ":8081")
	s := grpc.NewServer()
	protobuf.RegisterGetStudentInfoServer(s, &Service{})
	log.Fatal(s.Serve(lis))
}
