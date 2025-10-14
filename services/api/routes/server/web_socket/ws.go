package ws

import (
	"errors"
	"fmt"
	middHaders "middleware/headers"
	middJwt "middleware/jwt"
	"net/http"
	portsDB "ports/db"
	connWs "ws/connection"
	ws "ws/master"

	"github.com/gin-gonic/gin"
)

type RouteWs struct {
	Eng *gin.Engine
	DB portsDB.DB
	Hub *connWs.Hub
	KeyPem *[]byte
}
func (r *RouteWs)Run() {
	r.Eng.GET("server/ws", func(ctx *gin.Context){
		headers := middHaders.HeaderAPI{Ctx: ctx}
		token, hae := headers.AuthWs()
		if hae != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"header-field": hae.Field, 
				"message": hae.Msg,
			})
			return
		}

		jwt := middJwt.JWT{KeyPem: *r.KeyPem}
		fmt.Printf("TOKEN: %s\n", *token)
		ID, err := jwt.AuthToken(*token)
		if err != nil {
			var isme *middJwt.InvalidSignMethodError
			if errors.As(err, &isme) {
				fmt.Printf("ERRO IS -> %s", isme.Error())
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": isme.Error(),
				})
				return
			}
			var ntuide *middJwt.NotFoundUserIDError
			if errors.As(err, &ntuide) {
				fmt.Printf("ERRO IS -> %s", ntuide.Error())
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "token invalid",
				})
				return
			}
			fmt.Printf("ERRO IS -> %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		server := ws.ServerWS{
			UserID: *ID,
			Hub: r.Hub,
			W: ctx.Writer,
			R: ctx.Request,
		}
		server.Run()
	})
}