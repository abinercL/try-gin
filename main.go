package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	//call bindJSON to bind the received JSON to
	//newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	//add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// album represents data about a record  album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album slice to seed record album  data.
var albums = []album{
	{ID: "1", Title: "Blue train", Artist: "jhon coltrane", Price: 56.99},
	{ID: "2", Title: "jeru", Artist: "Gerry Muligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vauhan and Cliford Brown", Artist: "Sarah Vauhan", Price: 39.99},
}

// getAlbums responds whit the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
