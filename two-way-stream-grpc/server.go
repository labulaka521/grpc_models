package main

import (
	"google.golang.org/grpc"
	pb "grpc-models/two-way-stream-grpc/proto"
	"io"
	"log"
	"net"
)

var users = map[int32]pb.UserResponse{
	1: {Name: "WANG", Age: 12},
	2: {Name: "LI", Age: 13},
	3: {Name: "TAO", Age: 14},
}

type bidrectionalStreamServer struct{}

func (s *bidrectionalStreamServer) GetUserInfo(stream pb.UserService_GetUserInfoServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		u := users[req.ID]
		err = stream.Send(&u)
		if err != nil {
			return err
		}
		log.Printf("[RECEVIED REQUEST]: %v\n", req)
	}
	return nil
}

func main() {
	addr := "0.0.0.0:2333"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Listen Error: %v", err)
	}
	log.Println("Server Listen:", addr)

	// 创建gRPC服务器实例
	grpcServer := grpc.NewServer()

	// 向gRPC服务器注册服务
	pb.RegisterUserServiceServer(grpcServer, &bidrectionalStreamServer{})

	// start gRPC server
	grpcServer.Serve(listener)
}
