package models

type Message struct {
	FromID string `json:"from_id"`
	ToID string `json:"to_id"`
	Msg string `json:"message"`
}