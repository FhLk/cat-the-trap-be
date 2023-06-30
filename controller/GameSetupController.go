package controller

import (
	"cat-the-trap-back-end/service"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SetupBody struct {
	SessionID string `json:"sessionID"`
	Level     int    `json:"level"`
}

func Setup(c *gin.Context) {
	var getBody SetupBody
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

	board, session := service.StartSession(getBody.Level)

	c.JSON(http.StatusOK, gin.H{
		"sessionID": session,
		"board":     board,
		"turn":      0,
		"timeOut":   false,
		"token":     hashToken("TokenCheck") + "00",
		"canPlay":   true,
		"level":     getBody.Level,
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
	service.EndSession(getBody.SessionID)
	board, session := service.StartSession(getBody.Level)
	c.JSON(http.StatusOK, gin.H{
		"sessionID": session,
		"board":     board,
		"turn":      0,
		"timeOut":   false,
		"token":     hashToken("TokenCheck") + "00",
		"canPlay":   true,
		"level":     getBody.Level,
	})
}

func hashToken(o string) string {
	hash := sha256.New()

	hash.Write([]byte(o))

	checksum := hash.Sum(nil)

	checksumStr := hex.EncodeToString(checksum)
	return checksumStr
}
