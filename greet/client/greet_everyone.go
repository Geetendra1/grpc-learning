package main

import (
	"context"
	pb "grpc-learning/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	requests := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	// `waitc := make(chan struct{})` creates an unbuffered channel of type `struct{}`. This channel is
	// used to synchronize the sending and receiving goroutines in the `doGreetEveryone` function. The
	// `struct{}` type is used because it does not allocate any memory, which makes it more efficient than
	// using a channel of type `bool` or `int`. The `waitc` channel is used to signal the end of the
	// receiving goroutine, which in turn signals the sending goroutine to exit as well. The `<-waitc`
	// statement is a blocking statement that waits for a signal from the `waitc` channel, which ensures
	// that the function `doGreetEveryone` does not exit until both goroutines have completed their tasks.
	waitc := make(chan struct{})
	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
		// `close(waitc)` is closing the unbuffered channel `waitc`. This is used to signal the end of the
		// receiving goroutine, which in turn signals the sending goroutine to exit as well. The blocking
		// statement `<-waitc` waits for a signal from the `waitc` channel, which ensures that the function
		// `doGreetEveryone` does not exit until both goroutines have completed their tasks.
		close(waitc)
	}()

	// `<-waitc` is a blocking statement that waits for a signal from the `waitc` channel. It is used to
	// synchronize the two goroutines that are sending and receiving messages through the gRPC stream. The
	// `waitc` channel is closed when the receiving goroutine is done processing all the messages, which
	// signals the sending goroutine to exit as well. The blocking statement ensures that the function
	// `doGreetEveryone` does not exit until both goroutines have completed their tasks.
	<-waitc
}
