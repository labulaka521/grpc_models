package main

import (
	"google.golang.org/grpc"
	pb "grpc-models/client-stream-grpc/proto"
	"io"
	"log"
	"net"
)

var users = map[int32]pb.UserResponse{
	1: {Name: "wang", Age: 9},
	2: {Name: "li", Age: 10},
	3: {Name: "tao", Age: 11},
}

type clientsideStreamServer struct {}

func (s *clientsideStreamServer) GetUserInfo(stream pb.UserService_GetUserInfoServer) error {
	var lastID int32
	for {
		req, err := stream.Recv()
		// 接收客户端数据流发送完毕
		if err == io.EOF {
			// 返回最后一个ID的用户信息
			if u, ok := users[lastID]; ok {
				stream.SendAndClose(&u)
				return nil
			}
		}
		lastID = req.ID
		log.Printf("[RECEVIED REQUEST]: %v\n", req)
		log.Printf("[Send Response]: Name: %v Age:%v\n", users[lastID].Name, users[lastID].Age)
	}
	return nil
}

func main() {
	//  微服务监听的地址
	addr := "0.0.0.0:2333"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Listen Error: %v\n", err)
	}
	log.Println("Server Listen On Address:Port: ", addr)

	// 创建gRPC实例
	grpcServer := grpc.NewServer()

	// 向gRPC服务器注册服务
	pb.RegisterUserServiceServer(grpcServer, &clientsideStreamServer{})

	// 启动gRPC服务器
	grpcServer.Serve(listener)
}