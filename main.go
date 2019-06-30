package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-sample/Pb_Go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	addr = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, u *pb.User) (*pb.Response, error) {
	return &pb.Response{ErrCode: 0, ErrMsg: "success", Data: map[string]string{"name": "Hello " + u.Name}}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("%s server start at %s\n", time.Now(), addr)

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
