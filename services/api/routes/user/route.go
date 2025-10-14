package user

import (
	portsDB "ports/db"
	login "user/login"
	register "user/register"

	"github.com/gin-gonic/gin"
)

type RoutesUser struct {
	Login login.RouteLogin
	Register register.RouteRegister
}

func New(eng *gin.Engine, db portsDB.DB, keyPem *[]byte) RoutesUser {
	return RoutesUser{
		Login: login.RouteLogin{
			Eng: eng,
			DB: db,
			KeyPem: keyPem,
		},
		Register: register.RouteRegister{
			Eng: eng,
			DB: db,
			KeyPem: keyPem,
		},
	}
}