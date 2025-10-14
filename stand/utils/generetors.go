package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	r "math/rand"
	"os"
	"time"
)

type Generetors struct {
	Path string
}

func (g  *Generetors)KeyPem() error {
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
    file, err := os.Create(g.Path)
    if err != nil {
        return fmt.Errorf("error creating file: %v", err)
    }
    defer file.Close()

    if err := pem.Encode(file, pemBlock); err != nil {
        return fmt.Errorf("error writing key to file: %v", err)
    }

    fmt.Printf("✅ Pem Key save in %v", g.Path)
    return nil
}

func (g* Generetors)ApiKey(maxSize int) error {
	// Fonte aleatória com seed baseada no tempo atual
	r.Seed(time.Now().UnixNano())

	// Conjunto de caracteres possíveis
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Cria arquivo
	file, err := os.Create(g.Path)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Gera conteúdo e escreve no arquivo
	for i := 0; i < maxSize; i++ {
		randomChar := charset[r.Intn(len(charset))]
		_, err := file.WriteString(string(randomChar))
		if err != nil {
			return err
		}
	}
    
	fmt.Printf("✅ Api Key save in %v", g.Path)
	return nil
}