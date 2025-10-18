package ports

import "context"

type MessageBroker interface {
	Run(url string) (error, HandlerMB)
}

type HandlerMB interface {
	CreatePubWsEvents(msg WsPubMsg)
	OpenConsumerWsEvents(ctx context.Context, con WsConsumer) (<-chan []byte, error)
	Build() BuildRoutingKeys
}

type BuildRoutingKeys interface {
	PubWsEvents(senderID, consumerID int) string
	ConWsEvents(consumerID int) string
}
