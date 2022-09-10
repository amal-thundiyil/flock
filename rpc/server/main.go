package server

import (
	"context"
	"log"
	"net"

	"github.com/Deadcoder11u2/go-chat/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func main() {
	listner, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Printf("Error while starting the server on port 4040 %v", err)
	}

	srv := grpc.NewServer()
	go_rpc.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listner); e != nil {
		log.Printf("Error occured while serving in the port 4040 %v", e)
	}
}

func (s *server) Add(ctx context.Context,request *go_rpc.Request) (*go_rpc.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &go_rpc.Response{Result: result}, nil
}


func (s *server) Multiply(ctx context.Context,request *go_rpc.Request) (*go_rpc.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &go_rpc.Response{Result: result}, nil
}
