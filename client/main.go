package main

import (
	"context"
	pb "grpc-example/api/user"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	request := &pb.CreateRequest{
		User: &pb.Model{
			Email: "cuontb",
		},
	}
	resp, err := client.CreateUser(context.Background(), request)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.GetMessage(), "--", resp.GetStatus)
	// To do something with resp from instance server response
}
