package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenBody struct {
	Token string `json:"token"`
}

func generateToken() string {
	return "PASS"
}

func Authen(c *gin.Context) {
	session := sessions.Default(c)

	session.Set("key", "value")

	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	session.Delete("key")

	session.Clear()

	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	var getBody AuthenBody
	if err := c.BindJSON(&getBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	TOKEN := "eyJhbGciOiJIUzI1NiJ9.eyJtb2JpbGVObyI6Im1wUTBSMTJHTzAzNmY4ckVCbmZqVTg4OWwyczNnZGlGQUVzcCtNRWUrNzQ9IiwidGltZXN0YW1wIjoiMjAyMi0wMS0xNFQxMzowMDowNSswNzowMCJ9.gUvmq2MI9DAa5-AgWAX8DE7tL2elCD7VW8g-2gtYz9g"

	if getBody.Token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
		return
	}

	if getBody.Token != TOKEN {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid token value"})
		return
	}

	key := generateToken()
	c.JSON(http.StatusCreated, gin.H{"code": "201", "message": "success", "token": key})
}
