package main

import (
	"context"
	"log"
	"net"

	"github.com/Dinesh789kumar12/grpc-go/greet"
	"google.golang.org/grpc"
)

type server struct {
	greet.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, in *greet.GreetRequest) (*greet.GreetResponse, error) {
	req := in.GetReqGreet()
	res := greet.GreetResponse{ResGreet: "hello " + req}
	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50055")
	if err != nil {

		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Started listening on port 50055")

	s := grpc.NewServer()
	greet.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {

		log.Fatalf("Failed to serve: %v", err)

	}
}
