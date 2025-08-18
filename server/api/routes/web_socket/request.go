package websocket

type Handler struct {
	Accept string `json:"accept"`
	Authorization string `json:"authorization"`
}

type Request struct {
	ID string `form:"id" binding:"required"`
}