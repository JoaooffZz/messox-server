package ping

import (
	config "api/config"
	"net/http"
	"os"
	utils "utils"

	"github.com/gin-gonic/gin"
)

const (
	path = "API_KEY_PATH"
)

func Run(eng *gin.Engine) {
	eng.GET("/ping", func(ctx *gin.Context){

		token, isAuth := config.AuthHeader(ctx)
		if !isAuth {
			return
		}

		keyBytes, err := utils.GetGenericKey(os.Getenv(path))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		key := string(keyBytes)

		if (*token != key) {
			ctx.JSON(http.StatusUnauthorized, nil)
			return 
		}

		ctx.JSON(http.StatusOK, nil)
	})
}