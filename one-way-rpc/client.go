package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "grpc-models/one-way-rpc/proto"
)

// gRPC客户端

func main() {
	// gRPC服务器地址
	addr := "0.0.0.0:2333"

	//建立连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect Server Error: %v", err)
	}
	defer conn.Close()

	// 创建服务端的函数
	grpcClient := pb.NewUserServiceClient(conn)

	// 调用服务端的函数
	req := pb.UserRequest{ID: 2}
	resp, err := grpcClient.GetUserInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("Recevice Resp Error: %v\n", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", resp)
}
