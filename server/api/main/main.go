package main

import (
	// modelsWS "ws/models"

	routeWS "routes/web_socket"
	connWS "ws/connection"

	"github.com/gin-gonic/gin"
)

func main() {
	engGIN := gin.Default()

    hub := connWS.NewHub()
	go hub.Run()
	routeWS.Run(engGIN, hub)

	// engGIN.GET("/ws", func(ctx *gin.Context) {
	// 	var handler modelsWS.Handler
	// 	if err := ctx.ShouldBindJSON(&handler); err != nil {
	// 		ctx.JSON(404 , gin.H{"error":"not authorized"})
	// 		return
	// 	}
	// 	// chamar servico de autorizacao de token aqui

	// 	// chamar servico de websocket aqui

	// })

	engGIN.Run(":8080")
}