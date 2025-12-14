package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type MessageHandler func(ctx context.Context, body []byte) error

type Consumer struct {
	conn IConnection
	log  *logrus.Logger
}

type IConsumer interface {
	DeclareQueue(queueName string, durable bool) error
	BindQueue(queueName, exchange, routingKey string) error
	Consume(ctx context.Context, queueName string, handler MessageHandler) error
	ConsumeWithPrefetch(ctx context.Context, queueName string, prefetchCount int, handler MessageHandler) error
}

func NewConsumer(conn IConnection) IConsumer {
	logger := logrus.New()

	return &Consumer{
		conn: conn,
		log:  logger,
	}
}

func (c *Consumer) DeclareQueue(queueName string, durable bool) error {
	channel := c.conn.GetChannel()

	_, err := channel.QueueDeclare(
		queueName,
		durable,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	c.log.Info("Queue " + queueName + " declared successfully")

	return nil
}

func (c *Consumer) BindQueue(queueName, exchange, routingKey string) error {
	channel := c.conn.GetChannel()

	err := channel.QueueBind(
		queueName,
		routingKey,
		exchange,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	c.log.Info("Queue " + queueName + " bound successfully")

	return nil
}

func (c *Consumer) Consume(ctx context.Context, queueName string, handler MessageHandler) error {
	return c.ConsumeWithPrefetch(ctx, queueName, 5, handler)
}

func (c *Consumer) ConsumeWithPrefetch(ctx context.Context, queueName string, prefetchCount int, handler MessageHandler) error {
	channel := c.conn.GetChannel()

	err := channel.Qos(
		prefetchCount, // prefetch count
		0,             // prefetch size
		false,         // global
	)

	if err != nil {
		return fmt.Errorf("failed to set QoS: %w", err)
	}

	msgs, err := channel.Consume(
		queueName,
		"",    // consumer tag
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		return fmt.Errorf("failed to consume: %w", err)
	}

	c.log.Info("Startted consuming message")

	go c.handleDelivery(ctx, msgs, handler)

	return nil
}

func (c *Consumer) handleDelivery(ctx context.Context, msgs <-chan amqp091.Delivery, handler MessageHandler) error {
	for {
		select {
		case <-ctx.Done():
			c.log.Info("Consumer stopped")
			return nil
		case msg, ok := <-msgs:
			if !ok {
				c.log.Warn("Message channel closed")
				return nil
			}

			c.process(ctx, msg, handler)
		}
	}
}

func (c *Consumer) process(ctx context.Context, msg amqp091.Delivery, handler MessageHandler) {
	c.log.WithFields(logrus.Fields{
		"message_id":  msg.MessageId,
		"routing_key": msg.RoutingKey,
	}).Info("Processing message")

	var prettyJSON map[string]interface{}
	if err := json.Unmarshal(msg.Body, &prettyJSON); err != nil {
		c.log.WithField("body", prettyJSON).Debug("Message Body")
	}

	if err := handler(ctx, msg.Body); err != nil {
		c.log.Error("Failed to handle message")

		if err := msg.Nack(false, true); err != nil {
			c.log.Error("Failed to nack message")
		}

		return
	}

	if err := msg.Ack(false); err != nil {
		c.log.Error("Failed to ack message")
		return
	}

	c.log.WithField("routing_key", msg.RoutingKey).Info("Message processed successfully")
}
