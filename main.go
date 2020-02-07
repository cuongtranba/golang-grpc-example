package main

import (
	"context"
	pb "grpc-example/api/user"
	"log"
	"net"

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
