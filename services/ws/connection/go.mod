module ws/connection

go 1.24.7

replace mb/ports => ../../mb/ports

replace ws/models => ../models

require (
	github.com/gorilla/websocket v1.5.3
	mb/ports v0.0.0-00010101000000-000000000000
	ws/models v0.0.0-00010101000000-000000000000
)
