package main

import (
	"google.golang.org/grpc"
	pb "grpc-models/server-stream-grpc/proto"
	"log"
	"net"
)

var users = map[int32]pb.UserResponse{
	1: {Name: "wang", Age: 12},
	2: {Name: "li", Age: 13},
	3: {Name: "tao", Age: 14},
}

type serverSideStreamServer struct{}

func (s *serverSideStreamServer) GetUserInfo(req *pb.UserRequest, stream pb.UserService_GetUserInfoServer) error {
	// 响应数据流
	for _, user := range users {
		stream.Send(&user)
	}
	log.Printf("[RECEIVED REQUEST]: %v\n", req)
	return nil
}

func main() {
	addr := "0.0.0.0:2333"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Listen Error: %v", err)
	}
	log.Println("Server Listen:", addr)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &serverSideStreamServer{})

	grpcServer.Serve(listener)
}

