package register

import (
	"errors"
	middHaders "middleware/headers"
	middJwt "middleware/jwt"
	"net/http"
	portsDB "ports/db"
	utils "utils"

	"github.com/gin-gonic/gin"
)
const (accept = "application/json")
type Request struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Response struct {
	Token string `json:"token" binding:"required"`
	Name string `json:"name" binding:"required"`
	Profile []byte `json:"profile" binding:"required"`
	Bio string `json:"bio" binding:"required"`
}

type RouteRegister struct {
	Eng *gin.Engine
	DB portsDB.DB
	KeyPem *[]byte
}

func (r *RouteRegister)Run() {
	r.Eng.POST("/user/register", func(ctx *gin.Context){
		
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
        pass := utils.Sha512String(req.Password)
		user, err := r.DB.NewUser(req.Name, pass)
		if err != nil {
		    var ve *portsDB.ValidationError
			if errors.As(err, &ve){
				ctx.JSON(http.StatusConflict, gin.H{
					"field": ve.Field,
					"message": ve.Msg,
				})
				return
			}
			var sle *portsDB.StringLengthError
			if errors.As(err, &sle){
				ctx.JSON(http.StatusRequestEntityTooLarge, gin.H{
					"field": sle.Field,
					"message": "exceeded size",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		jwt := middJwt.JWT{KeyPem: *r.KeyPem}
		token, err := jwt.CreateToken(user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		
		ctx.JSON(http.StatusCreated, Response{
			Token: *token,
			Name: user.Name,
			Profile: user.Profile,
			Bio: user.Bio,
		})
	})
}