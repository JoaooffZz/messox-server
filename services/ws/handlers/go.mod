module ws/handlers

go 1.24.7

replace ws/models => ../models

replace ws/connection => ../connection

replace mb/ports => ../../mb/ports

require (
	mb/ports v0.0.0-00010101000000-000000000000
	ws/connection v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	ws/models v0.0.0-00010101000000-000000000000 // indirect
)
