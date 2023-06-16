package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "grpc-learning/blog/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Server struct {
	pb.BlogServiceServer
}

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func main() {
	dsn := os.Getenv("DB_STRING")
	fmt.Println("db string", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// DB.AutoMigrate(Movie{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
