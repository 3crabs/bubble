package main

import (
	"context"
	pb "github.com/3crabs/bubble/task"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"net"
)

type server struct {
	tasks map[string]*pb.TaskInfo
}

func (s *server) GetTasks(_ context.Context, _ *emptypb.Empty) (*pb.TaskInfos, error) {
	var tasks []*pb.TaskInfo
	if s.tasks == nil {
		s.tasks = make(map[string]*pb.TaskInfo)
	}
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return &pb.TaskInfos{Tasks: tasks}, status.New(codes.OK, "").Err()
}

func (s *server) AddTask(_ context.Context, in *pb.TaskCreateInfo) (*wrapperspb.StringValue, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while generating task id: %v", err)
	}
	if s.tasks == nil {
		s.tasks = make(map[string]*pb.TaskInfo)
	}
	task := pb.TaskInfo{Id: out.String(), Name: in.Name, Created: timestamppb.Now(), Done: false}
	s.tasks[task.Id] = &task
	return wrapperspb.String(task.Id), status.New(codes.OK, "").Err()
}

func (s *server) DoneTask(_ context.Context, in *wrapperspb.StringValue) (*emptypb.Empty, error) {
	if s.tasks == nil {
		s.tasks = make(map[string]*pb.TaskInfo)
	}
	value, exists := s.tasks[in.Value]
	if exists {
		return &emptypb.Empty{}, status.New(codes.NotFound, "task not found").Err()
	}
	value.Done = true
	s.tasks[value.Id] = value
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}

func (s *server) DeleteTask(_ context.Context, in *wrapperspb.StringValue) (*emptypb.Empty, error) {
	if s.tasks == nil {
		s.tasks = make(map[string]*pb.TaskInfo)
	}
	value, exists := s.tasks[in.Value]
	if exists {
		return &emptypb.Empty{}, status.New(codes.NotFound, "task not found").Err()
	}
	s.tasks[value.Id] = nil
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
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
