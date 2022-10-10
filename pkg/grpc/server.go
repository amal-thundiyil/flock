package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	exec "os/exec"
	"time"

	"github.com/amal-thundiyil/flock/pkg/proto"

	"github.com/go-co-op/gocron"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedJobServiceServer
}

var scheduler gocron.Scheduler

func StartGrpcServer() {
	scheduler = *gocron.NewScheduler(time.UTC)

	fmt.Println("GRPC server listening on port 4040")

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterJobServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}

}

func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatal("Cannot create file for cron job")
	}

	return file
}

func RunJob(request *proto.JobRequest) {
	cmd := exec.Command(request.GetExecutor().String(), request.Name)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("Error while executing the code")
		return
	}

	fmt.Println("Output of cron job: " + string(stdout))
}

func (s *server) ScheduleJob(ctx context.Context, request *proto.JobRequest) (*proto.JobResponse, error) {
	fmt.Println("Cron Job Details")
	fmt.Printf("Filecontent: %s\n", request.GetFileBody())
	fmt.Printf("Filename: %s\n", request.GetName())
	fmt.Printf("CronCommand: %s\n", request.GetCronSchedule())

	file := CreateFile(request.GetName())

	file.WriteString(request.GetFileBody())

	file.Close()

	jobRes := proto.JobResponse{Body: "Cron Job scheduled successfully"}
	scheduler.Cron(request.GetCronSchedule()).Do(RunJob, request)

	scheduler.StartAsync()

	return &jobRes, nil
}
