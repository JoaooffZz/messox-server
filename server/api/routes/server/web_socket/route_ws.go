package websocket

import (
	// "fmt"
	"net/http"
	// "os"
	// svcJwt "services/jwt"
	// utils "utils"
	connWS "ws/connection"
	ws "ws/master"

	"github.com/gin-gonic/gin"
)

func Run(eng *gin.Engine, hub *connWS.Hub) {
	eng.GET("/ws", func(ctx *gin.Context) {
		var req Request
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.Writer.WriteHeader(http.StatusUnauthorized)
			// ctx.JSON(http.StatusUnauthorized , nil)
			return
		}

		serverWs := ws.ServerWS{
			Id: req.ID,
			Hub: hub,
			R: ctx.Request,
			W: ctx.Writer,
		}
		serverWs.Run(ctx)



		// var handler Handler
		// if err := ctx.ShouldBindJSON(&handler); err != nil {
		// 	ctx.JSON(http.StatusUnauthorized , gin.H{"notAuth":"error binding JSON"})
		// 	return
		// }

		// keyPem, err := utils.GetKeyPem(os.Getenv("PATH_KEY_PEM"))
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, 
		// 		gin.H{"internal": fmt.Sprintf("failed: %v", err)},
		// 	)
		// 	return
		// }

		// service := svcJwt.ServiceJWT{KeyPem: keyPem}
		// isAuth, err := service.Verify(handler.Authorization)
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, 
		// 		gin.H{"internal": fmt.Sprintf("failed: %v", err)},
		// 	)
		// 	return
		// }
		// if !isAuth {
		//     ctx.JSON(http.StatusUnauthorized, 
		// 		gin.H{"notAuth": "invalid token"},
		// 	)
		// 	return
		// }

		// userID, err := service.GetUserIDFromJWT()
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, 
		// 		gin.H{"internal": fmt.Sprintf("failed: %v", err)},
		// 	)
		// 	return
		// }
		

		// chamar servico de websocket aqui

	})
}