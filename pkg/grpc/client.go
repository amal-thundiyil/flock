package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"mod/pkg/proto"
	"google.golang.org/grpc"
)

func clientConnect() {

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error while connecting to the server")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := proto.NewJobServiceClient(conn)

	req := &proto.JobRequest{FileContent: "print(\"Hello World\")", FileName: "main.py", CronCommand: "*/1 * * * *"}

	res, err := client.ScheduleJob(ctx, req)
	fmt.Println(res)
	defer conn.Close()
}
