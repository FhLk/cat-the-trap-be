package controller

import (
	"cat-the-trap-back-end/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SetupBody struct {
	Level int `json:"level"`
}

var board [][]map[string]interface{} = service.GameSetup()

func Setup(c *gin.Context) {
	var getBody SetupBody
	if err := c.BindJSON(&getBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if c.Request.ContentLength == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"board": board,
		"turn":  0,
		"token": "TokenCheck00",
	})
}
