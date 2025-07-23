module ws/master

go 1.22.1

replace ws/connection => ../connection

replace ws/models => ../models

require (
	ws/connection v0.0.0-00010101000000-000000000000
	ws/models v0.0.0-00010101000000-000000000000
)

require github.com/gorilla/websocket v1.5.3 
