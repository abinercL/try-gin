package route

import (
	c "github.com/abinercl/try-gin/src/controller"
	"github.com/gin-gonic/gin"
)

func Routers() {
	router := gin.Default()
	router.GET("/albums", c.GetAlbums)
	router.GET("/albums/:id", c.GetAlbumByID)
	router.POST("/albums", c.PostAlbums)
	router.Run("localhost:8080")
}
