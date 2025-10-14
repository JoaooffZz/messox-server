package login

import (
	"net/http"

	utils "utils"

	middHaders "middleware/headers"
	middJwt "middleware/jwt"
	portsDB "ports/db"

	"github.com/gin-gonic/gin"
)
const (accept = "application/json")
type Request struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Response struct {
	Token string `json:"token" binding:"required"`
}

type RouteLogin struct {
	Eng *gin.Engine
	DB portsDB.DB
	KeyPem *[]byte
}

func (r *RouteLogin)Run() {
	r.Eng.POST("/user/login", func(ctx *gin.Context){
		
		headers := middHaders.HeaderAPI{Ctx: ctx}
		_, hae := headers.AuthHTTP(accept)
		if hae != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"header-field": hae.Field, 
				"message": hae.Msg,
			})
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
			user, err = r.DB.GetUser(name)
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

					jwt := middJwt.JWT{KeyPem: *r.KeyPem}
					token, err := jwt.CreateToken(user.ID)
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, nil)
						return
					}

					ctx.JSON(http.StatusOK, Response{Token: *token})
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