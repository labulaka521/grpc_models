package main

import (
	"context"
	pb "grpc-models/one-way-rpc/proto"
	"log"
	"net"
	"google.golang.org/grpc"
)

// 数据查询结果
var users = map[int32]pb.UserResponse{
	1: {Name: "wang", Age: 12},
	2: {Name: "litao", Age: 22},
	3: {Name: "labulaka", Age: 34},
}

//实现了user.pb.go 中的UserServiceServer接口
type simpleServer struct{}

func (s *simpleServer) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	if user, ok := users[req.ID]; ok {
		resp = &user
	}
	log.Printf("[RECEVIED REQUEST]: %v\n", req)
	return
}

func main() {
	// 指定微服务的服务端监听地址
	addr := "0.0.0.0:2333"
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Listen Error: %v", err)
	} else {
		log.Println("Server Listen On Port: ", addr)
	}

	// 创建gRPC服务实例
	grpcServer := grpc.NewServer()

	// 向gRPC服务器注册服务
	pb.RegisterUserServiceServer(grpcServer, &simpleServer{})

	// 启动gRPC
	// 阻塞等待客户端的调用
	grpcServer.Serve(listener)

}