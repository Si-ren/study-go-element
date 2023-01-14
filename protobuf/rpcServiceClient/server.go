package main

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	protobuf "protobuf/gen/go"

	"google.golang.org/grpc"
)

type Service struct{}

func (s *Service) GetIPInfo(context.Context, *emptypb.Empty) (*protobuf.IPInfo, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var ipAddress string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddress = ipnet.IP.String()
			}
		}
	}
	return &protobuf.IPInfo{Ip: ipAddress}, nil
}

func (s *Service) mustEmbedUnimplementedGetStudentInfoServer() {}
func main() {
	//go startGRPCGateway()
	lis, _ := net.Listen("tcp", ":8081")
	s := grpc.NewServer()
	protobuf.RegisterIPServer(s, &Service{})
	log.Fatal(s.Serve(lis))
}

//
//func startGRPCGateway() {
//	c := context.Background()
//	c, cancel := context.WithCancel(c)
//	defer cancel()
//	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
//		runtime.MIMEWildcard, &runtime.JSONPb{
//			protojson.MarshalOptions{
//				AllowPartial:    true,
//				EmitUnpopulated: true,
//			},
//			protojson.UnmarshalOptions{
//				DiscardUnknown: true,
//			},
//		},
//	))
//	//err := protobuf.RegisterGetStudentInfoHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
//	err := protobuf.RegisterGetIPInfoHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
//
//	if err != nil {
//		log.Fatalf("error registering getStudentInfoHandler: %v", err)
//	}
//	err = http.ListenAndServe(":8082", mux)
//}
