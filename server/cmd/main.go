package main

import (
	adapterDB "adapters/db"
	connDB "db/connection"
	"fmt"
	"os"
	portsDB "ports/db"
	rUser "routes/user"
	utils "utils"

	"github.com/gin-gonic/gin"
)

type ConfigCMD struct {
	EngGIN *gin.Engine
	DB portsDB.DB
	KeyPem *[]byte
}

func main() {
	fmt.Print("\ninit cmd...\n")
	cmd, err := new()
	if err != nil {
		fmt.Print("\ncmd failed!\n")
		return
	}
	fmt.Print("\ncmd sucess!\n")

	routesUser := rUser.New(cmd.EngGIN, cmd.DB, cmd.KeyPem)
	routesUser.Login.Run()

	cmd.EngGIN.Run(":8080")
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

	return &ConfigCMD{
		EngGIN: engGIN,
		DB: adpDB,
		KeyPem: &keyPem,
	}, nil
}

