package main

import (
	"cat-the-trap-back-end/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controller.ListAlbums)
	router.GET("/albums/:id", controller.ShowAlbum)

	router.Run("localhost:8080")
}
