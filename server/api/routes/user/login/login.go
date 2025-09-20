package login

import (
	"net/http"

	config "api/config"
	utils "utils"

	middleware "middleware/jwt"
	portsDB "ports/db"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Response struct {
	Token string `json:"token" binding:"required"`
}

type UserLogin struct {
	Eng *gin.Engine
	DB portsDB.DB
	KeyPem *[]byte
}

func (u *UserLogin)Run() {
	u.Eng.POST("/user/login", func(ctx *gin.Context){
		contentHeader := config.AuthHeader(ctx)
		if (!contentHeader.IsAuth) {
			ctx.JSON(http.StatusBadRequest, contentHeader.Header)
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		var user *portsDB.User
		done := make(chan bool)
		notFound := make(chan bool)
		errChan := make(chan error)

		go func(name string, done chan bool, notFound chan bool, errChan chan error) {
			var err error
			user, err = u.DB.GetUser(name)
			if err != nil {
				errChan <- err
				return
			}
			if user == nil {
				notFound <- true
				return
			}
			done <- true
		}(req.Name, done, notFound, errChan)

		for {
			select {
			case isDone := <-done:
				if isDone {
					passSha := utils.Sha512String(req.Password)
					if user.Password != passSha {
						ctx.JSON(http.StatusUnauthorized, nil)
						return
					}

					jwt := middleware.MiddlewareJWT{KeyPem: *u.KeyPem}
					token, err := jwt.CreateToken(user.ID)
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, nil)
						return
					}

					ctx.JSON(http.StatusOK, Response{Token: token})
					return
				}

			case isNotFound := <-notFound:
				if isNotFound {
					ctx.JSON(http.StatusNotFound, gin.H{"user": req.Name})
					return
				}

			case errResult := <-errChan:
				if errResult != nil {
					ctx.JSON(http.StatusInternalServerError, nil)
					return
				}
			}
		}
	})
}