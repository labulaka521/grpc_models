package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "grpc-models/two-way-stream-grpc/proto"
	"time"
)

// 双向数据流gRPC客户端

func main() {
	addr := "0.0.0.0:2333"

	//建立连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect Server Error: %v", err)
	}
	defer conn.Close()

	// 创建gRPC客户端实例
	grpcClient := pb.NewUserServiceClient(conn)
	stream, err := grpcClient.GetUserInfo(context.Background())
	if err != nil {
		log.Fatalf("Receive Stream Error: %v", err)
	}

	//  向服务端发送数据流并处理响应流
	var i int32
	for i = 1; i < 4; i ++ {
		stream.Send(&pb.UserRequest{ID: i})
		time.Sleep(1 * time.Second)
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("Resp Error: %v", err)
		}
		log.Printf("[RECEVIED RESPONSE]: %v\n", resp)
	}
}
