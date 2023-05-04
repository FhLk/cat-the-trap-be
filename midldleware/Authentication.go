package midldleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authentication(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")
	//body := c.Request.Body

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		return
	}
}

//TOKEN := "eyJhbGciOiJIUzI1NiJ9.eyJtb2JpbGVObyI6Im1wUTBSMTJHTzAzNmY4ckVCbmZqVTg4OWwyczNnZGlGQUVzcCtNRWUrNzQ9IiwidGltZXN0YW1wIjoiMjAyMi0wMS0xNFQxMzowMDowNSswNzowMCJ9.gUvmq2MI9DAa5-AgWAX8DE7tL2elCD7VW8g-2gtYz9g"

func validateToken(token string) error {
	TOKEN := "PASS"
	if token == "" {
		return fmt.Errorf("token should not be empty")
	} else if token != TOKEN {
		return fmt.Errorf("token is not invalid")
	}
	return nil
}
