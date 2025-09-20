package register

import (
	config "api/config"
	"errors"
	"net/http"

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
	Name string `json:"name" binding:"required"`
	Profile []byte `json:"profile" binding:"required"`
	Bio string `json:"bio" binding:"required"`
}

type UserRegister struct {
	Eng *gin.Engine
	DB portsDB.DB
	KeyPem *[]byte
}

func (u *UserRegister)Run() {
	u.Eng.POST("/user/register", func(ctx *gin.Context){
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
       
		user, err := u.DB.NewUser(req.Name, req.Password)
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
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		jwt := middleware.MiddlewareJWT{KeyPem: *u.KeyPem}
		token, err := jwt.CreateToken(user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		
		ctx.JSON(http.StatusCreated, Response{
			Token: token,
			Name: user.Name,
			Profile: user.Profile,
			Bio: user.Bio,
		})
	})
}