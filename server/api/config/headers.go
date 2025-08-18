package config

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	accept = "Accept"
	acceptH = "application/json"

	auth = "Authorization"
	authH = "Bearer "
)

func AuthHeader(c *gin.Context) (*string, bool) {
	acceptHeader := c.GetHeader(accept)
	if (acceptHeader != acceptH) {
		c.JSON(http.StatusUnauthorized, gin.H{"accept": "malformed"})
		return nil, false
	}
	authHeader := c.GetHeader(auth)
	if !strings.HasPrefix(authHeader, authH) {
		c.JSON(http.StatusUnauthorized, gin.H{"authorization": "missing"})
		return nil, false
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	return &token, true
}
