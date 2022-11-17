package main

import (
	"github.com/amal-thundiyil/flock/pkg/grpc"
)

func main() {
	grpc.ClientConnect("4040", "echo Hello World")
}
