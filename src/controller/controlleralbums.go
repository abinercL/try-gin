package controller

import (
	"net/http"

	"github.com/abinercl/try-gin/src/lib"
	"github.com/abinercl/try-gin/src/model"
	"github.com/gin-gonic/gin"
)

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum model.Album

	//call bindJSON to bind the received JSON to
	//newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	//add the new album to the slice.
	lib.Albums = append(lib.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over  the list of albums, looking for
	//an  album whose  ID  value matches  the parameter
	for _, a := range lib.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album  not found"})
}

// getAlbums responds whit the list of all albums as JSON
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, lib.Albums)
}
