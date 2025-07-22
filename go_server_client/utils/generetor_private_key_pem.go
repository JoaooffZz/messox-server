package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateKeyPem() {
    // Gera uma chave RSA de 2048 bits
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        panic(fmt.Sprintf("Erro ao gerar chave: %v", err))
    }

    // Codifica a chave em formato DER (PKCS#1)
    keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

    // Cria o bloco PEM
    pemBlock := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: keyBytes,
    }

    // Cria o arquivo e escreve o conteúdo PEM
    file, err := os.Create("private_key.pem")
    if err != nil {
        panic(fmt.Sprintf("Erro ao criar arquivo: %v", err))
    }
    defer file.Close()

    if err := pem.Encode(file, pemBlock); err != nil {
        panic(fmt.Sprintf("Erro ao escrever chave no arquivo: %v", err))
    }

    fmt.Println("✅ Chave RSA salva em 'private_key.pem'")
}