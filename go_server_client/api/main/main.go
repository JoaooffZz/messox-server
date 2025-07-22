package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engGIN := gin.Default()

	engGIN.GET("/ws")

	engGIN.Run(":8080")
}