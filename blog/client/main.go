package main

import (
	"grpc-learning/blog/controllers"
	pb "grpc-learning/blog/proto"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

// type Something struct {
// 	Client pb.BlogServiceClient
// }

func main() {
	var r = gin.Default()

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()
	client := pb.NewBlogServiceClient(conn)
	r.POST("/create", func(c *gin.Context) { controllers.CreateBlog(c, client) })

	// id := createBlog(c, s)

	// fmt.Println("id", id)
	// createBlog(c)
}
