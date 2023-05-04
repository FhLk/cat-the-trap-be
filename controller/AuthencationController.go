package controller

import (
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
	var getBody AuthenBody
	TOKEN := "eyJhbGciOiJIUzI1NiJ9.eyJtb2JpbGVObyI6Im1wUTBSMTJHTzAzNmY4ckVCbmZqVTg4OWwyczNnZGlGQUVzcCtNRWUrNzQ9IiwidGltZXN0YW1wIjoiMjAyMi0wMS0xNFQxMzowMDowNSswNzowMCJ9.gUvmq2MI9DAa5-AgWAX8DE7tL2elCD7VW8g-2gtYz9g"
	if err := c.BindJSON(&getBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if c.Request.ContentLength == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
		return
	}

	if getBody.Token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if getBody.Token != TOKEN {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid token value"})
		return
	}

	key := generateToken()
	c.JSON(http.StatusCreated, gin.H{"code": "201", "message": "success", "token": key})
}
