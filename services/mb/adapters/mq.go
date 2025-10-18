package adapters

import (
	"context"
	"fmt"
	mbp "mb/ports"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchangeWS = "ws.events.msg"
)

// broker
type RabbitMQBroker struct{}

func (mb *RabbitMQBroker) Run(url string) (*RabbitMQHandler, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &RabbitMQHandler{conn: conn}, nil
}

// handler
type RabbitMQHandler struct {
	conn    *amqp.Connection
	builder RabbitMQBuildRoutingKeys
}

func (mbh *RabbitMQHandler) CreatePubWsEvents(msg mbp.WsPubMsg) {
	ch, _ := mbh.conn.Channel()
	defer ch.Close()
	ch.Publish(
		exchangeWS,
		msg.RoutingKey,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  msg.ContentType,
			Body:         msg.Body,
		},
	)
}

func (mbh *RabbitMQHandler) OpenConsumerWsEvents(ctx context.Context, con mbp.WsConsumer) (<-chan []byte, error) {
	ch, _ := mbh.conn.Channel()

	err := ch.ExchangeDeclare(
		exchangeWS,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		con.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		con.RoutingKey,
		exchangeWS,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msg := make(chan []byte)
	go func(goCtx context.Context) {
		defer func() {
			ch.Close()
			close(msg)
		}()
		for {
			select {
			case <-goCtx.Done():
				return
			case m, ok := <-msgs:
				if !ok {
					return
				}
				msg <- m.Body
				// m.Ack(false)
			}
		}
	}(ctx)

	return msg, nil
}

func (mbh *RabbitMQHandler) Build() mbp.BuildRoutingKeys {
	return &mbh.builder
}

// building
type RabbitMQBuildRoutingKeys struct{}

func (mbb *RabbitMQBuildRoutingKeys) PubWsEvents(senderID, consumerID int) string {
	return fmt.Sprintf("sender.%d.consumer.%d", senderID, consumerID)
}
func (mbb *RabbitMQBuildRoutingKeys) ConWsEvents(consumerID int) string {
	return fmt.Sprintf("*.*.consumer.%d", consumerID)
}
