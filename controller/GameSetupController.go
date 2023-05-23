package controller

import (
	"cat-the-trap-back-end/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SetupBody struct {
	Level int `json:"level"`
}

var board [][]map[string]interface{}

func Setup(c *gin.Context) {
	var getBody SetupBody
	turn = 0
	token = "TokenCheck00"
	if err := c.BindJSON(&getBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if getBody.Level >= 4 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if c.Request.ContentLength == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
		return
	}
	board = service.GameSetup(getBody.Level)

	service.StartSession(board, getBody.Level)

	c.JSON(http.StatusOK, gin.H{
		"board":   board,
		"turn":    0,
		"timeOut": false,
		"token":   "TokenCheck00",
		"canPlay": true,
		"level":   getBody.Level,
	})
}

func Reset(c *gin.Context) {
	var getBody SetupBody
	if err := c.BindJSON(&getBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if getBody.Level >= 4 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	board = service.ResetBoard(getBody.Level)
	turn = 0
	token = "TokenCheck00"
	c.JSON(http.StatusOK, gin.H{
		"board":   board,
		"turn":    0,
		"timeOut": false,
		"token":   "TokenCheck00",
		"canPlay": true,
		"level":   getBody.Level,
	})
}
