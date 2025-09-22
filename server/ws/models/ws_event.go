package models

type WsEvent struct {
	Sender *int `json:"sender,omitempty"`
	Adderess *int `json:"adderess"`
	Date string `json:"date"`
	Type string `json:"type"`
	Data string `json:"data"` 
} 