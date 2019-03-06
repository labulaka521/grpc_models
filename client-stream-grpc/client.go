package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-models/client-stream-grpc/proto"
	"log"
)

func main() {
	// grpC服务器地址
	addr := "0.0.0.0:2333"

	// 建立链接 不使用认证
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect Server Error: %v\n", err)
	}
	defer conn.Close()

	// 创建gRPC客户端实例
	grpcClient := pb.NewUserServiceClient(conn)

	// 调用服务端函数
	stream, err := grpcClient.GetUserInfo(context.Background())

	// 向服务端发送数据
	var i int32
	for i = 1; i < 4; i++ {
		err := stream.Send(&pb.UserRequest{ID: i})
		if err != nil {
			log.Fatalf("Send Error: %v\n", err)
		}
	}

	// 接收服务端
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Recevice Resp Error: %v\n", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", resp)
}
