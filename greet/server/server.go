package main

import pb "grpc-learning/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
