package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomString(length int) string {
	// Define os caracteres possíveis
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// Seed para garantir aleatoriedade a cada execução
	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func main() {
	randomChars := randomString(500)
	fmt.Println(randomChars)
	fmt.Println("Tamanho:", len(randomChars)) // Deve imprimir 500
}