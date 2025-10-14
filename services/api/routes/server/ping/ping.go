package ping

import (
	"fmt"
	middHeaders "middleware/headers"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	accept = "*/*"
)

// If you want to make your server accessible to everyone,
// Just pass api key as nil.
type RoutePing struct {
	Eng    *gin.Engine
	ApiKey *string
}

func (r *RoutePing) Run() {
	r.Eng.GET("server/ping", func(ctx *gin.Context) {

		headers := middHeaders.HeaderAPI{Ctx: ctx}
		token, hae := headers.AuthHTTP(accept)
		if hae != nil {
			fmt.Printf("ERROR IS: %s", hae.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"header-field": hae.Field,
				"message":      hae.Msg,
			})
			return
		}

		if r.ApiKey != nil {
			fmt.Println(*token)
			if *token != *r.ApiKey {
				ctx.JSON(http.StatusUnauthorized, nil)
				return
			}
		}

		ctx.JSON(http.StatusOK, nil)
	})
}
