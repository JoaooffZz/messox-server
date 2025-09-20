package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512String(s string) string {
    hash := sha512.Sum512([]byte(s))
    return hex.EncodeToString(hash[:])
}