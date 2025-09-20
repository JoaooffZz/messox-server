package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type MiddlewareJWT struct {
	KeyPem []byte
}

func (m *MiddlewareJWT) CreateToken(userID int) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(m.KeyPem)
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

func (m *MiddlewareJWT) AuthToken(tokenStr string) (bool, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(m.KeyPem)
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

func (m *MiddlewareJWT) GetUserIdInToken(tokenStr string) (*int, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(m.KeyPem)
	if err != nil {
		return nil, fmt.Errorf("verify jwt: erro ao carregar chave privada: %v", err)
	}
    
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("verify jwt: invalid sign method: %v", token.Header["alg"])
		}
		return &key.PublicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("verify jwt: %v", err)
	}

	if token == nil {
		return nil, fmt.Errorf("get user id: token is nil")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("get user id: claims are not of type jwt.MapClaims")
	}

	userID, ok := claims["user_id"].(int)
	if !ok {
		return nil, fmt.Errorf("get user id: user_id not found in claims")
	}

	return &userID, nil
}