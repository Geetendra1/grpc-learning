package main

import (
	pb "grpc-learning/blog/proto"
)

type BlogItem struct {
	ID       string `gorm:"primarykey"`
	Content  string
	Title    string
	AuthorID string
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID,
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}
