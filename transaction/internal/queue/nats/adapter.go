package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type NatsConnection struct {
	conn *nats.Conn
}

func NewNatsConnection(url string) (*NatsConnection, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}
	return &NatsConnection{conn: conn}, nil
}

func (nc *NatsConnection) Close() {
	if nc.conn != nil {
		nc.conn.Close()
	}
}

type NatsPublisher struct {
	natsConn *NatsConnection
}

func NewNatsPublisher(conn *NatsConnection) *NatsPublisher {
	return &NatsPublisher{natsConn: conn}
}

func (p *NatsPublisher) Publish(subject string, data []byte) error {
	return p.natsConn.conn.Publish(subject, data)
}

type NatsSubscriber struct {
	natsConn *NatsConnection
}

func NewNatsSubscriber(conn *NatsConnection) *NatsSubscriber {
	return &NatsSubscriber{natsConn: conn}
}

func (s *NatsSubscriber) Subscribe(subject string, callback func(data []byte)) error {
	_, err := s.natsConn.conn.Subscribe(subject, func(msg *nats.Msg) {
		callback(msg.Data)
	})
	return err
}
