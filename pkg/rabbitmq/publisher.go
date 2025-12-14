package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	conn IConnection
	log  *logrus.Logger
}

type IPublisher interface {
	Publish(ctx context.Context, exchange, routingKey string, message interface{}) error
	DeclareExchange(exchange, exchangeType string) error
}

func NewPublisher(conn IConnection) IPublisher {
	logger := logrus.New()

	return &Publisher{
		conn: conn,
		log:  logger,
	}
}

func (p *Publisher) DeclareExchange(exchange, exchangeType string) error {
	channel := p.conn.GetChannel()

	err := channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	p.log.Info("Exchange declared successfully")

	return nil
}

func (p *Publisher) Publish(ctx context.Context, exchange, routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marchal message: %w", err)
	}

	channel := p.conn.GetChannel()

	err = channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp091.Persistent,
			Timestamp:    time.Now(),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
