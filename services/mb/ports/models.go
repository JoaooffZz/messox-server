package ports

type WsPubMsg struct {
	ContentType string
	Body        []byte
	RoutingKey  string
}

type WsConsumer struct {
	QueueName  string
	RoutingKey string
}
