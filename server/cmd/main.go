package main

import (
	adapterDB "adapters/db"
	connDB "db/connection"
	"fmt"
	"os"
	portsDB "ports/db"
	rServer "routes/server"
	rUser "routes/user"
	utils "utils"
	connWs "ws/connection"

	"github.com/gin-gonic/gin"
)

type ConfigCMD struct {
	EngGIN *gin.Engine
	DB portsDB.DB
	KeyPem *[]byte
	Hub *connWs.Hub
	ApiKey *string
}

func main() {
	fmt.Print("\ninit cmd...\n")
	cmd, err := new()
	if err != nil {
		fmt.Print("\ncmd failed!\n")
		return
	}
	fmt.Print("\ncmd sucess!\n")
	// Hub roda em goroutine
    go cmd.Hub.Run()

	routesUser := rUser.New(cmd.EngGIN, cmd.DB, cmd.KeyPem)
	routesUser.Login.Run()
	routesUser.Register.Run()

	routesServer := rServer.New(cmd.EngGIN, cmd.DB, cmd.KeyPem, cmd.ApiKey, cmd.Hub)
	routesServer.Ping.Run()
	routesServer.Ws.Run()

	// Gin roda na main (bloqueante)
	if err := cmd.EngGIN.Run(":8080"); err != nil {
		panic(err)
	}
}

func new() (*ConfigCMD, error) {
	// init gin
	engGIN := gin.Default()
    
    // init db
	fmt.Print("\ninit db\n")
	configDB := connDB.New()
	conn, err := connDB.GetConn(configDB)
	if err != nil {
		fmt.Printf("\nerror db: %v\n", err)
		return nil, err
	}
	adpDB := adapterDB.New(conn)
    
	// init keys
	fmt.Print("\ninit keys\n")
	keyPem, err := utils.GetGenericKey(os.Getenv("KEY_PEM_PATH"))
	if err != nil {
		fmt.Printf("\nerror get key: %v\n", err)
		return nil, err
	}
    fmt.Print("\ninit hub\n")
	
	hub := connWs.NewHub()
	return &ConfigCMD{
		EngGIN: engGIN,
		DB: adpDB,
		KeyPem: &keyPem,
		Hub: hub,
		ApiKey: nil,
	}, nil
}

