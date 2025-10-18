module ws/main

go 1.24.7

replace mb/ports => ../../mb/ports

replace mb/adapters => ../../mb/adapters

replace ws/handlers => ../handlers

replace ws/models => ../models

replace ws/connection => ../connection

require (
	mb/adapters v0.0.0-00010101000000-000000000000
	ws/handlers v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
	mb/ports v0.0.0-00010101000000-000000000000 // indirect
	ws/connection v0.0.0-00010101000000-000000000000 // indirect
	ws/models v0.0.0-00010101000000-000000000000 // indirect
)
