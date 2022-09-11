package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	exec "os/exec"

	proto "github.com/Deadcoder11u2/go-chat/proto"

	"github.com/go-co-op/gocron"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{
	proto.UnimplementedJobServiceServer
}


var scheduler gocron.Scheduler

func main() {
	scheduler = *gocron.NewScheduler(time.UTC)

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterJobServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func CreateFile(filename string) (*os.File) {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatal("Cannot create file for cron job")
	}

	return file
}

func (s *server) ScheduleJob(ctx context.Context, request *proto.JobRequest) (*proto.JobResponse, error) {
	fmt.Println("Cron Job Details")
	fmt.Printf("Filecontent: %s\n", request.GetFileContent())
	fmt.Printf("Filename: %s\n", request.GetFileName())
	fmt.Printf("CronCommand: %s\n", request.GetCronCommand())

	file := CreateFile(request.GetFileName())

	file.WriteString(request.GetFileContent())

	file.Close()

	jobRes := proto.JobResponse{Body: "Cron Job scheduled successfully"}
	scheduler.Cron(request.GetCronCommand()).Do(func() {
		cmd := exec.Command("python", request.GetFileName())
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Sprintln("Error while executing the code")
			return
		}

		fmt.Sprintln("Output of cron job: " + string(stdout))
	})

	scheduler.StartAsync()

	return &jobRes, nil
}
