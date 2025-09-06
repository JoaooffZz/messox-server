package models

type WsEvent struct {
	Sender *string `json:"sender,omitempty"`
	Adderess string `json:"adderess"`
	Date int64 `json:"date"`
	ContentType string `json:"content_type"`
	Data string `json:"data"` 
} 