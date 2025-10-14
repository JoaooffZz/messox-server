package utils

import (
	"fmt"
	"os"
)

func GetGenericKey(path string) ([]byte, error) {
	if path == "" {
		return nil, fmt.Errorf("PATH is not set")
	}

	key, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error read key in file: %v", err)
	}

	if len(key) == 0 {
		return nil, fmt.Errorf("key file is empty")
	}

	return key, nil
}