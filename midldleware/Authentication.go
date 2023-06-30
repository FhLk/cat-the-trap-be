package midldleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authentication(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		return
	}
}

func validateToken(token string) error {
	TOKEN := "0af1223668c978dfef6b0b5a6ce3361abd6f7d46c4e6f13ee8f491a326b39328"
	if token == "" {
		return fmt.Errorf("token should not be empty")
	} else if token != TOKEN {
		return fmt.Errorf("token is not invalid")
	}
	return nil
}
