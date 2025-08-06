package models

type Message struct {
	Type string `json:"type"`
	From string `json:"from"`
	To string `json:"to"`
	Content string `json:"content"`
}