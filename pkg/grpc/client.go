package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/amal-thundiyil/flock/pkg/proto"
	"google.golang.org/grpc"
)

func ClientConnect() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error while connecting to the server")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := proto.NewJobServiceClient(conn)

	req := &proto.JobRequest{FileBody: "for i in range(10):\n\tprint(i, end = \",\")\nprint(\"Done\")", Name: "Main.py", CronSchedule: "*/1 * * * *", Executor: proto.JobRequest_PYTHON, Config: &proto.JobRequest_ExecutorConfig{Command: "python3 Main.py"}}

	res, err := client.ScheduleJob(ctx, req)
	fmt.Println(res.GetBody())
	defer conn.Close()
}
