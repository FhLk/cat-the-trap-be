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
	Time  bool   `json:"time"`
	Turn  int    `json:"turn"`
	Token string `json:"token"`
	Level int    `json:"level"`
}

//var turn int = 0
//var token string = "TokenCheck00"

func Play(c *gin.Context) {
	var getBody playBody
	if err := c.BindJSON(&getBody); err != nil {
		fmt.Println("wow1")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	turn := getBody.Turn - 1
	token := fmt.Sprintf("TokenCheck0%d", turn)

	if getBody.Token != token {
		fmt.Println(getBody.Token)
		fmt.Println(token)
		fmt.Println("wow2")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if !(getBody.Turn > turn && getBody.Turn < turn+2) {
		fmt.Println("wow3")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if getBody.Level == 0 {
		fmt.Println("wow4")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if c.Request.ContentLength == 0 {
		fmt.Println("wow5")
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
		fmt.Println("wow6")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	} else if err != nil {
		fmt.Println("wow7")
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

//func Time(c *gin.Context) {
//	var getBody timeBody
//	if err := c.BindJSON(&getBody); err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
//		return
//	}
//
//	if getBody.Token != token {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
//		return
//	}
//
//	if !(getBody.Turn > turn && getBody.Turn < turn+2) {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
//		return
//	}
//
//	if getBody.Level == 0 {
//		fmt.Println("wow")
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
//		return
//	}
//
//	if c.Request.ContentLength == 0 {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
//		return
//	}
//	turn = getBody.Turn
//	board, newToken, err := service.TimeOut(turn, board)
//	if board != nil && err == nil && newToken == "" {
//		c.JSON(http.StatusOK, gin.H{
//			"board":   board,
//			"turn":    turn,
//			"timeOut": true,
//			"token":   token,
//			"canPlay": false,
//			"level":   getBody.Level,
//		})
//		return
//	} else if board == nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
//		return
//	} else if err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
//		return
//	}
//
//	token = newToken
//	c.JSON(http.StatusOK, gin.H{
//		"board":   board,
//		"turn":    turn,
//		"timeOut": false,
//		"token":   token,
//		"canPlay": true,
//		"level":   getBody.Level,
//	})
//}
