package main

import (
  "fmt"
	"context"
	"log"
	"time"

	go_rpc "github.com/Deadcoder11u2/go-chat/proto"
	"google.golang.org/grpc"
)

func main() {

  conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
  if err != nil {
    log.Fatal("Error while connecting to the server")
  }

  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  client := go_rpc.NewAddServiceClient(conn)

  a := int64(19)
  b := int64(20)

  req := &go_rpc.Request{A: a, B: b}


  res, err := client.Multiply(ctx, req)
  fmt.Println(res)
  defer conn.Close()
}
