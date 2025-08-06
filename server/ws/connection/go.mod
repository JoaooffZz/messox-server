module ws/connection

go 1.22.1

replace ws/models => ../models

require (
	github.com/gorilla/websocket v1.5.3
	ws/models v0.0.0-00010101000000-000000000000
)
