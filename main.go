package main

import (
	"github.com/amal-thundiyil/flock/pkg/grpc"
  "os"
)

func main() {
  port := os.Args[1]
  grpc.StartGrpcServer(port)
}
