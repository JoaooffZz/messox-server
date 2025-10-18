package models

type WsEvent struct {
	Sender   *int   `json:"sender,omitempty"`
	Adderess *int   `json:"adderess"`
	Type     string `json:"type"`
	Date     string `json:"date"`
	Data     string `json:"data"` // base 64
}

type TypeMessage struct {
	Msg string `json:"message"`
}

type TypeConnectionRequest struct{}
