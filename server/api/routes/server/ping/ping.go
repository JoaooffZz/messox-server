package ping

import (
	config "api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// If you want to make your server accessible to everyone,
// Just pass api key as nil.
type ServerPing struct {
	Eng *gin.Engine
	ApiKey *string
}
func (s *ServerPing)Run() {
	s.Eng.GET("/ping", func(ctx *gin.Context){

		content := config.AuthHeader(ctx)
		if !content.IsAuth {
			ctx.JSON(http.StatusUnauthorized, content.Header)
			return
		}

		if (content.Token != s.ApiKey) {
			ctx.JSON(http.StatusUnauthorized, nil)
			return 
		}

		ctx.JSON(http.StatusOK, nil)
	})
}