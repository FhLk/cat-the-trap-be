package controller

import (
	"cat-the-trap-back-end/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type playBody struct {
	SessionID string `json:"sessionID"`
	Turn      int    `json:"turn"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Block     string `json:"block"`
	Token     string `json:"token"`
	Level     int    `json:"level"`
}

type timeBody struct {
	SessionID string `json:"sessionID"`
	Time      bool   `json:"time"`
	Turn      int    `json:"turn"`
	Token     string `json:"token"`
	Level     int    `json:"level"`
}

func Play(c *gin.Context) {
	var getBody playBody
	if err := c.BindJSON(&getBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	turn := getBody.Turn - 1
	token := fmt.Sprintf("TokenCheck0%d", turn)

	if getBody.Token != token {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if !(getBody.Turn > turn && getBody.Turn < turn+2) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if getBody.Level == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if c.Request.ContentLength == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
		return
	}
	turn = getBody.Turn
	board, newToken, err := service.UpdateBoardSessions(getBody.SessionID, getBody.X, getBody.Y, turn, getBody.Block)
	if board != nil && err == nil && newToken == "" {
		c.JSON(http.StatusOK, gin.H{
			"board":   board,
			"turn":    turn,
			"timeOut": true,
			"token":   token,
			"canPlay": false,
			"level":   getBody.Level,
		})
		return
	} else if board == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	token = newToken
	c.JSON(http.StatusOK, gin.H{
		"board":   board,
		"turn":    turn,
		"timeOut": false,
		"token":   token,
		"canPlay": true,
		"level":   getBody.Level,
	})
}

func Time(c *gin.Context) {
	var getBody timeBody
	if err := c.BindJSON(&getBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	turn := getBody.Turn - 1
	token := fmt.Sprintf("TokenCheck0%d", turn)

	if getBody.Token != token {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if !(getBody.Turn > turn && getBody.Turn < turn+2) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if getBody.Level == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if c.Request.ContentLength == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
		return
	}
	turn = getBody.Turn
	board, newToken, err := service.TimeOutSessions(getBody.SessionID, turn)
	if board != nil && err == nil && newToken == "" {
		c.JSON(http.StatusOK, gin.H{
			"board":   board,
			"turn":    turn,
			"timeOut": true,
			"token":   token,
			"canPlay": false,
			"level":   getBody.Level,
		})
		return
	} else if board == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	token = newToken
	c.JSON(http.StatusOK, gin.H{
		"board":   board,
		"turn":    turn,
		"timeOut": false,
		"token":   token,
		"canPlay": true,
		"level":   getBody.Level,
	})
}
