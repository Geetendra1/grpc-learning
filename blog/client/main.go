package main

import (
	"flag"
	"fmt"
	"grpc-learning/blog/controllers"
	pb "grpc-learning/blog/proto"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	fmt.Println("runnig main")

	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.POST("/movies", func(c *gin.Context) { controllers.CreateBlog(c, client) })
	r.GET("/movies", func(c *gin.Context) { controllers.ListAllBlogs(c, client) })
	r.GET("/movies/:id", func(c *gin.Context) { controllers.GetMovie(c, client) })
	r.PUT("/movies/:id", func(c *gin.Context) { controllers.UpdateMovie(c, client) })
	r.DELETE("/movies/:id", func(c *gin.Context) { controllers.DeleteMovie(c, client) })

	r.Run(":5000")
}
