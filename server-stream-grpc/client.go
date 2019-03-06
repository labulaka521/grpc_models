package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "grpc-models/server-stream-grpc/proto"
)

// 服务端流式RPC客户端

func main() {
	addr := "0.0.0.0:2333"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect Server Error: %v", err)
	}
	defer conn.Close()

	// 创建gRPC客户端实例
	grpcClient := pb.NewUserServiceClient(conn)

	// 调用服务端的函数
	req := pb.UserRequest{ID:1}
	stream, err := grpcClient.GetUserInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("Recevie Resp Error: %v", err)
	}

	// 接收数据
	for {
		resp, err := stream.Recv()
		if err == io.EOF { // EOF表示服务端数据发送完毕
			break
		}
		if err != nil {
			log.Fatalf("Receive Error: %v", err)
		}
		log.Printf("[RECEVIED RESPONSE]: %v\n", resp)
	}

}
