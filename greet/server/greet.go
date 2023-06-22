package main

import (
	"context"
	pb "grpc-learning/greet/proto"
	"log"
)

func (*Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet was invoked %v", in)
	return &pb.GreetResponse{Result: "Hello" + in.FirstName}, nil
}
