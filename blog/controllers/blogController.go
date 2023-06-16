package controllers

import (
	"fmt"
	pb "grpc-learning/blog/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

var addr string = "0.0.0.0:50051"

func CreateBlog(c *gin.Context, client pb.BlogServiceClient) {
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// if err != nil {
	// 	log.Fatalf("Couldn't connect to client: %v\n", err)
	// }

	// defer conn.Close()
	// client := pb.NewBlogServiceClient(conn)

	var body struct {
		ID       string `gorm:"primarykey"`
		Content  string
		Title    string
		AuthorID string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})
		return
	}
	// fmt.Println("client", client)

	fmt.Println("body", body)

}
