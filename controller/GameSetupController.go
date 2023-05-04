package controller

import (
	"cat-the-trap-back-end/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(c *gin.Context) {
	board, blocks, destination, err := service.GameSetup()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"board":       board,
		"blocks":      blocks,
		"destination": destination,
	})
}

func aStar(c *gin.Context) {

}
