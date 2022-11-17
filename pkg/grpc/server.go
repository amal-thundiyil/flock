package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	exec "os/exec"
	"strings"
	"sync"
	"time"

	"github.com/amal-thundiyil/flock/pkg/proto"

	"github.com/go-co-op/gocron"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedJobServiceServer
}

// var m sync.Mutex
// var wg sync.WaitGroup

var scheduler gocron.Scheduler

var m sync.Mutex
var wg sync.WaitGroup
var server_port string

func StartGrpcServer(port string) {
	wg.Add(1)
	server_port = port

	scheduler = *gocron.NewScheduler(time.UTC)

	fmt.Println("GRPC server listening on port 4040")

	listener, err := net.Listen("tcp", port)
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
	fmt.Println(request.Config.GetCommand())
	arr := strings.Split(request.Config.GetCommand(), " ")
	cmd := exec.Command(arr[0], arr[1:]...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("Error while executing the code")
		fmt.Println(err)
		return
	}

	fmt.Println("From port :" + server_port)
	fmt.Println("===================Output Of Cron Job===================")
	fmt.Println("Output of cron job: " + string(stdout))
	fmt.Println("===================End Of Job Output===================\n\n")
}

func (s *server) ScheduleJob(ctx context.Context, request *proto.JobRequest) (*proto.JobResponse, error) {
	m.Lock()
	ScheduleMutex := func(request *proto.JobRequest) {
		fmt.Println("Curent Time: ", time.Now())
		fmt.Println("======================Scheduling Cron Job======================")
		fmt.Println("Cron Job Details")
		fmt.Printf("Filecontent: %s\n", request.GetFileBody())
		fmt.Printf("Filename: %s\n", request.GetName())
		fmt.Printf("CronCommand: %s\n", request.GetCronSchedule())
		fmt.Println("======================End Of Job Details======================\n\n")

		scheduler.Cron(request.GetCronSchedule()).Do(RunJob, request)

		scheduler.StartAsync()
		time.Sleep(6 * time.Second)
		m.Unlock()
	}
	go ScheduleMutex(request)
	wg.Wait()

	return &proto.JobResponse{Body: "Job Scheduled"}, nil
}
