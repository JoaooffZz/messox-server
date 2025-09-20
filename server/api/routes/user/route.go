package user

import (
	portsDB "ports/db"
	login "user/login"

	"github.com/gin-gonic/gin"
)

type RoutesUser struct {
	Login login.UserLogin
}

func New(eng *gin.Engine, db portsDB.DB, keyPem *[]byte) RoutesUser {
	return RoutesUser{
		Login: login.UserLogin{
			Eng: eng,
			DB: db,
			KeyPem: keyPem,
		},
	}
}