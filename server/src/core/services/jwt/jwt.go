package servicejwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type ServiceJWT struct {
	KeyPem []byte
	token *jwt.Token
}

func (s *ServiceJWT) Build(userID string) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(s.KeyPem)
	if err != nil {
		return "", fmt.Errorf("build jwt: %v", err)
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userID,
	})

	sigJwt, err := claim.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("build jwt: %v", err)
	}
	return sigJwt, nil
}

func (s *ServiceJWT) Verify(tokenStr string) (bool, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(s.KeyPem)
	if err != nil {
		return false, fmt.Errorf("verify jwt: erro ao carregar chave privada: %v", err)
	}

	s.token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("verify jwt: invalid sign method: %v", token.Header["alg"])
		}
		return &key.PublicKey, nil
	})
	if err != nil {
		return false, fmt.Errorf("verify jwt: %v", err)
	}

	return s.token.Valid, nil
}

func (s *ServiceJWT) GetUserIDFromJWT() (string, error) {
	if s.token == nil {
		return "", fmt.Errorf("get user id: token is nil")
	}

	claims, ok := s.token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("get user id: claims are not of type jwt.MapClaims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("get user id: user_id not found in claims")
	}

	return userID, nil
}
