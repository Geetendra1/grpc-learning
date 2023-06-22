package main

import (
	"context"
	pb "grpc-learning/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	r, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Clement"})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", r.Result)
}
