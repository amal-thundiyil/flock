package main

import (
	"context"
	"fmt"
	"net"

	proto "github.com/Deadcoder11u2/go-chat/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{
	proto.UnimplementedAddServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	fmt.Println("Hello from the multiplication")
	a, b := request.GetA(), request.GetB()

	result := a * b

	fmt.Println(result)

	return &proto.Response{Result: result}, nil
}
