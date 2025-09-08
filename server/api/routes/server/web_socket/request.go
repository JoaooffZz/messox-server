package websocket

type Request struct {
	ID string `form:"id" binding:"required"`
}