package main

import (
	"context"
	helloworld "github.com/zhaoleigege/datastruct/grpc/hello"
	"google.golang.org/grpc"
	"log"
)

func main(){
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v1", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "test"})
	if err != nil {
		log.Fatalf("could not greet: %v1", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
