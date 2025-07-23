package utils

import (
	"fmt"
	"os"
)

func GetKeyPem(pemPath string) ([]byte, error) {

	if pemPath == "" {
		return nil, fmt.Errorf("PATH_KEY_PEM is not set")
	}

	keyPem, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, fmt.Errorf("error read key PEM file: %v", err)
	}

	if len(keyPem) == 0 {
		return nil, fmt.Errorf("key PEM file is empty")
	}

	return keyPem, nil
}