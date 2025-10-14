package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	KeyPem []byte
}

func (m *JWT) CreateToken(userID int) (*string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(m.KeyPem)
	if err != nil {
		return nil, fmt.Errorf("build jwt: %v", err)
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userID,
	})

	sigJwt, err := claim.SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("build jwt: %v", err)
	}
	return &sigJwt, nil
}

func (m *JWT) AuthToken(tokenStr string) (*int, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(m.KeyPem)
	if err != nil {
		return nil, fmt.Errorf("uploud keypem failed: %v", err)
	}
    
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			alg, ok := token.Header["alg"].(string)
			if !ok {
				alg = "not-found algorithm"
			}
			return nil, &InvalidSignMethodError{Alg: alg}
		}
		return &key.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, fmt.Errorf("token is nil")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("claims are not of type jwt.MapClaims")
	}
    
	cKey := "user_id"
	uid, ok := claims["user_id"].(float64)
	fmt.Printf("\nID: %s", uid)
	if !ok {
		return nil, &NotFoundUserIDError{ClaimKey: cKey}
	}
	userID := int(uid)

	return &userID, nil
}