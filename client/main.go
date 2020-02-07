package main

import (
	"context"
	pb "grpc-example/api/user"
	"io"
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
	log.Println(resp.GetMessage(), "--", resp.GetStatus())
	stream, _ := client.GenInt(context.Background(), &pb.CreateGenRequest{})
	for {
		in, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(in.Number)
	}
	// To do something with resp from instance server response
}
