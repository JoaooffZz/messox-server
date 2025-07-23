package servicejwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type ServiceJWT struct {
	KeyPem []byte
}

func (s *ServiceJWT) Build(userID string) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(s.KeyPem)
	if err != nil {
		return "", fmt.Errorf("build jwt: %v", err)
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userID,
	})

	sigJwy, err := claim.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("build jwt: %v", err)
	}
	return sigJwy, nil
}

func (s *ServiceJWT) Verify(tokenStr string) (bool, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(s.KeyPem)
	if err != nil {
		return false, fmt.Errorf("verify jwt: erro ao carregar chave privada: %v", err)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("verify jwt: invalid sign method: %v", token.Header["alg"])
		}
		return &key.PublicKey, nil
	})
	if err != nil {
		return false, fmt.Errorf("verify jwt: %v", err)
	}

	return token.Valid, nil
}


