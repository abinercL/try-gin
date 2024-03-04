package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
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

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over  the list of albums, looking for
	//an  album whose  ID  value matches  the parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album  not found"})
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
