package controller

import (
	"cat-the-trap-back-end/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type playBody struct {
	Turn  int    `json:"turn"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Token string `json:"token"`
}

var turn int = 0
var token string = "TokenCheck00"

func Play(c *gin.Context) {
	var getBody playBody
	if err := c.BindJSON(&getBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if c.Request.ContentLength == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
		return
	}
	if getBody.Token != token {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if !(getBody.Turn > turn && getBody.Turn < turn+2) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	turn = getBody.Turn
	board, newToken := service.UpdateBoard(getBody.X, getBody.Y, turn, board)
	token = newToken
	c.JSON(http.StatusOK, gin.H{
		"board": board,
		"turn":  turn,
		"token": token,
	})
}

func Reset(c *gin.Context) {
	board = service.ResetBoard()
	turn = 0
	token = "TokenCheck00"
	c.JSON(http.StatusOK, gin.H{
		"board": board,
		"turn":  0,
		"token": "TokenCheck00",
	})
}
