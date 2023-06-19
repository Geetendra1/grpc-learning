package controllers

import (
	pb "grpc-learning/blog/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"Title"`
	Genre string `json:"genre"`
}

func CreateBlog(c *gin.Context, client pb.MovieServiceClient) {
	var movie Movie
	err := c.ShouldBind(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	data := &pb.Movie{
		Title: movie.Title,
		Genre: movie.Genre,
	}
	res, err := client.CreateMovie(c, &pb.CreateMovieRequest{
		Movie: data,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"movie": res.Movie,
	})
}

func ListAllBlogs(c *gin.Context, client pb.MovieServiceClient) {
	res, err := client.GetMovies(c, &pb.ReadMoviesRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movies": res.Movies,
	})

}

func GetMovie(c *gin.Context, client pb.MovieServiceClient) {
	id := c.Param("id")
	res, err := client.GetMovie(c, &pb.ReadMovieRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movie": res.Movie,
	})

}

func UpdateMovie(c *gin.Context, client pb.MovieServiceClient) {
	var movie Movie
	id := c.Param("id")

	err := c.ShouldBind(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := client.UpdateMovie(c, &pb.UpdateMovieRequest{
		Movie: &pb.Movie{
			Id:    id,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movie": res.Movie,
	})
	return

}

func DeleteMovie(c *gin.Context, client pb.MovieServiceClient) {
	id := c.Param("id")
	res, err := client.DeleteMovie(c, &pb.DeleteMovieRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if res.Success == true {
		c.JSON(http.StatusOK, gin.H{
			"message": "Movie deleted successfully",
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error deleting movie",
		})
		return
	}
}
