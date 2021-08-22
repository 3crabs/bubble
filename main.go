package main

import (
	"context"
	pb "github.com/3crabs/bubble/task"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"net"
)

type server struct{}

func (s server) GetTasks(ctx context.Context, in *emptypb.Empty) (*pb.TaskInfos, error) {
	panic("implement me")
}

func (s server) AddTask(ctx context.Context, in *pb.TaskCreateInfo) (*wrapperspb.StringValue, error) {
	panic("implement me")
}

func (s server) DoneTask(ctx context.Context, in *wrapperspb.StringValue) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s server) DeleteTask(ctx context.Context, in *wrapperspb.StringValue) (*emptypb.Empty, error) {
	panic("implement me")
}

func main() {
	port := ":8080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskServer(s, &server{})

	log.Printf("starting listener on port %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild to serve: %v", err)
	}
}
