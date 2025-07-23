package main

import (
	"fmt"
	"os"
	utils "utils"
)

func TestGenerateKeyPem() {
	fmt.Printf("test: initiating key generation...\n")
	err := utils.GenerateKeyPem("../private/key.pem")
	if err != nil {
		fmt.Printf("test: error\n")
		return
	}
	fmt.Printf("test: sucess!\n")
}

func TestGetKeyPem() {
	fmt.Printf("test: initiating  getter key-pem...\n")
	keyPem, err := utils.GetKeyPem(os.Getenv("PATH_KEY_PEM"))
	if err != nil {
		fmt.Printf("test: error -> %v\n", err)
	}
	fmt.Printf("test: sucess!, keyPem: %v\n", keyPem)
}