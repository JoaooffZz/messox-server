package config

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	accept = "Accept"
	acceptType = "application/json"

	auth = "Authorization"
	authH = "Bearer "
)

type HeaderContent struct {
	Token *string
	IsAuth bool
	Header map[string]string
}

func AuthHeader(c *gin.Context) HeaderContent {
	acceptReceived := c.GetHeader(accept)
	if (acceptReceived != acceptType) {
		return HeaderContent{
			Token: nil,
			IsAuth: false,
			Header: map[string]string{"accept": "malformed"},
		}
	}

	authHeader := c.GetHeader(auth)
	if !strings.HasPrefix(authHeader, authH) {
		return HeaderContent{
			Token: nil,
			IsAuth: false,
			Header: map[string]string{"authorization": "missing"},
		}
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	return HeaderContent{
		Token: &token,
		IsAuth: true,
		Header: nil,
	}
}
