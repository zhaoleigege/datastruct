package main

import (
	"context"
	"fmt"
	helloworld "github.com/zhaoleigege/datastruct/grpc/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Printf(req.Name)
	return &helloworld.HelloReply{Message: "ok"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v1", err)
	}
	s := grpc.NewServer()

	helloworld.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v1", err)
	}
}
