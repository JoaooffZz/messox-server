package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateKeyPem(path string) error {
    // Gera uma chave RSA de 2048 bits
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return fmt.Errorf("erro genereta key: %v", err)
    }

    // Codifica a chave em formato DER (PKCS#1)
    keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

    // Cria o bloco PEM
    pemBlock := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: keyBytes,
    }

    // Cria o arquivo e escreve o conteúdo PEM
    file, err := os.Create(path)
    if err != nil {
        return fmt.Errorf("error creating file: %v", err)
    }
    defer file.Close()

    if err := pem.Encode(file, pemBlock); err != nil {
        return fmt.Errorf("error writing key to file: %v", err)
    }

    fmt.Println("✅ Key RSA save in 'private_key.pem'")
    return nil
}