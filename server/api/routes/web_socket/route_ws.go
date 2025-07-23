package websocket

import (
	"fmt"
	"net/http"
	"os"
	svcJwt "services/jwt"
	utils "utils"

	"github.com/gin-gonic/gin"
)

func Websocket(eng *gin.Engine) {
	eng.GET("/ws", func(ctx *gin.Context) {
		var handler Handler
		if err := ctx.ShouldBindJSON(&handler); err != nil {
			ctx.JSON(http.StatusUnauthorized , gin.H{"notAuth":"error binding JSON"})
			return
		}

		keyPem, err := utils.GetKeyPem(os.Getenv("PATH_KEY_PEM"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, 
				gin.H{"internal": fmt.Sprintf("failed: %v", err)},
			)
			return
		}

		service := svcJwt.ServiceJWT{KeyPem: keyPem}
		isAuth, err := service.Verify(handler.Authorization)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, 
				gin.H{"internal": fmt.Sprintf("failed: %v", err)},
			)
			return
		}
		if !isAuth {
		    ctx.JSON(http.StatusUnauthorized, 
				gin.H{"notAuth": "invalid token"},
			)
			return
		}

		// chamar servico de websocket aqui

	})
}