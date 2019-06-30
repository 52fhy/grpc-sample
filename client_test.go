package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-sample/Pb_Go"
	"testing"
)

func TestExec(t *testing.T)  {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		t.Errorf("dial error: %v\n", err)
	}
	defer conn.Close()

	// 实例化客户端
	client := pb.NewGreeterClient(conn)

	// 调用服务

	user := pb.User{}
	user.Id = 1
	user.Name = "test"
	result, err := client.SayHello(context.Background(), &user)
	if err != nil {
		t.Errorf("grpc error: %v\n", err)
	}
	t.Logf("Recevied: %v\n", result)
}
