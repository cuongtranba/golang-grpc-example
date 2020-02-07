package main

import (
	"context"
	pb "grpc-example/api/user"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{
		Message: "create ok",
		Status:  "true",
	}, nil
}
func (s *server) GetUsers(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Status: req.Keyword,
	}, nil
}
func (s *server) FindUser(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	return &pb.FindResponse{
		User: &pb.Model{
			Id: req.Id,
		},
	}, nil
}

// GenInt(ctx context.Context, in *CreateGenRequest, opts ...grpc.CallOption) (User_GenIntClient, error)

func (s *server) GenInt(req *pb.CreateGenRequest, stream pb.User_GenIntServer) error {
	var count int32 = 1
	for {
		stream.Send(&pb.GenResponse{Number: count})
		count++
		time.Sleep(time.Second * 2)
	}
}

func main() {
	port := ":8080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServer(srv, &server{})
	log.Printf("listen on %s", port)
	srv.Serve(lis)
}
