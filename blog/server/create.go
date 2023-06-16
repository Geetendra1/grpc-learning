package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "grpc-learning/blog/proto"
)

func (*Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlogs was invoked with %v\n", in)
	in.GetId()

	fmt.Println("IN id", in)

	blog := &BlogItem{
		ID:       in.Id,
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	result := DB.Create(&blog)
	if result.Error != nil {
		errorMessage := fmt.Sprintf("Internal error: %s", result.Error.Error())
		return nil, errors.New(errorMessage)
	}

	return &pb.BlogId{Id: blog.ID}, nil
}
