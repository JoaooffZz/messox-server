package jwt

import "fmt"

type InvalidSignMethodError struct {
	Alg string
}
func (i *InvalidSignMethodError)Error() string {
	return fmt.Sprintf("jwt invalid, Algorithm: %s,", i.Alg)
}

type NotFoundUserIDError struct {
	ClaimKey string
}
func (n *NotFoundUserIDError)Error() string {
	return fmt.Sprintf("claim(%s) not found in token", n.ClaimKey)
}