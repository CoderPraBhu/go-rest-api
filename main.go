package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// serialize to json with lowercase field name
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "One Love", Artist: "Blue", Price: 10.99},
	{ID: "2", Title: "Legend", Artist: "Bob Marley", Price: 12.99},
	{ID: "3", Title: "Nothing but the best", Artist: "Bob Dylan", Price: 12.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
