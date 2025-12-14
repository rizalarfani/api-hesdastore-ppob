package rabbitmq

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	config  Config
}

type IConnection interface {
	GetChannel() *amqp091.Channel
	IsConnected() bool
	Reconnect() error
	Close() error
}

func NewConnection(cfg Config) (IConnection, error) {
	conn := &Connection{
		config: cfg,
	}

	if err := conn.connect(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Connection) connect() error {
	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		c.config.Username,
		c.config.Password,
		c.config.Host,
		c.config.Port,
		c.config.VHost,
	)

	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return fmt.Errorf("failed to connect RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}

	c.conn = conn
	c.channel = channel

	log.Println("RabbitMQ connected")

	go c.handleConnectionClose()

	return nil
}

func (c *Connection) GetChannel() *amqp091.Channel {
	return c.channel
}

func (c *Connection) IsConnected() bool {
	return c.conn != nil && c.channel != nil
}

func (c *Connection) Close() error {
	if c.channel != nil {
		if err := c.channel.Close(); err != nil {
			return fmt.Errorf("failed to close channer: %w", err)
		}
	}

	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return fmt.Errorf("failed to close connection: %w", err)
		}
	}

	return nil
}

func (c *Connection) Reconnect() error {
	c.Close()

	if err := c.connect(); err != nil {
		return err
	}

	return nil
}

func (c *Connection) handleConnectionClose() {
	closeErr := make(chan *amqp091.Error)
	c.conn.NotifyClose(closeErr)

	err := <-closeErr
	if err != nil {
		log.Printf("⚠️ RabbitMQ connection closed: %v", err)
		log.Println("Reconnecting...")
		c.Reconnect()
	}
}
