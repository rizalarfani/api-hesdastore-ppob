package rabbitmq

import (
	"context"
)

type Manager struct {
	conn      IConnection
	publisher IPublisher
}

type IManager interface {
	GetConnectioin() IConnection
	GetPublisher() IPublisher
	GetConsumer() IConsumer
	Close() error
	SetupWebhook(ctx context.Context) error
}

func NewManager(cfg Config) (IManager, error) {
	conn, err := NewConnection(cfg)

	if err != nil {
		return nil, err
	}

	publisher := NewPublisher(conn)

	return &Manager{
		conn:      conn,
		publisher: publisher,
	}, nil
}

func (m *Manager) GetConnectioin() IConnection {
	return m.conn
}

func (m *Manager) GetPublisher() IPublisher {
	return m.publisher
}

func (m *Manager) GetConsumer() IConsumer {
	return NewConsumer(m.conn)
}

func (m *Manager) Close() error {
	return m.conn.Close()
}

func (m *Manager) SetupWebhook(ctx context.Context) error {
	publiser := m.GetPublisher()
	consumer := m.GetConsumer()

	if err := publiser.DeclareExchange("transaction.events", "topic"); err != nil {
		return err
	}

	if err := consumer.DeclareQueue("transaction.update.webhook", true); err != nil {
		return err
	}

	if err := consumer.BindQueue("transaction.update.webhook", "transaction.events", "transaction.update"); err != nil {
		return err
	}

	return nil
}
