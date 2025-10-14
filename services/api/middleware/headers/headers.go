package headers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type HeaderAPI struct {
	Ctx *gin.Context
}

type HeaderAPIError struct {
	Field string
	Msg   string
}

func (h *HeaderAPIError) Error() string {
	return fmt.Sprintf("Field: %s, Error: %s", h.Field, h.Msg)
}

func (h *HeaderAPI) AuthHTTP(expectedAccept string) (*string, *HeaderAPIError) {
	acceptHeader := h.Ctx.GetHeader("Accept")

	// if expectedAccept != "*/*" {
	if acceptHeader != expectedAccept {
		return nil, &HeaderAPIError{Field: "Accept", Msg: "malformed"}
	}
	// }
	authHeader := h.Ctx.GetHeader("Authorization")
	fmt.Printf("HEADER AUTH: %s", authHeader)
	if !strings.HasPrefix(authHeader, "Bearer") {
		return nil, &HeaderAPIError{Field: "Authorization", Msg: "missing (token)"}
	}

	tokenPrefix := strings.TrimPrefix(authHeader, "Bearer")
	token := strings.ReplaceAll(tokenPrefix, " ", "")

	return &token, nil
}

func (h *HeaderAPI) AuthWs() (*string, *HeaderAPIError) {
	upgradeHeader := h.Ctx.GetHeader("Upgrade")
	if upgradeHeader != "websocket" {
		return nil, &HeaderAPIError{Field: "Upgrade", Msg: "not is websocket"}
	}

	connectionHeader := h.Ctx.GetHeader("Connection")
	if connectionHeader != "Upgrade" {
		return nil, &HeaderAPIError{Field: "Connection", Msg: "missing (upgrade)"}
	}

	authHeader := h.Ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer") {
		return nil, &HeaderAPIError{Field: "Authorization", Msg: "missing (token)"}
	}

	token := strings.TrimPrefix(authHeader, "Bearer")
	token = strings.ReplaceAll(token, " ", "")

	return &token, nil
}
