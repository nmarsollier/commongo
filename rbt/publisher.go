package rbt

import (
	"encoding/json"

	"github.com/nmarsollier/commongo/log"
	"github.com/streadway/amqp"
)

type RabbitPublisher[T any] interface {
	PublishTo(exchange string, routingKey string, data T) error
	Publish(data T) error
}

type rabbitPublisher[T any] struct {
	ch           *amqp.Channel
	log          log.LogRusEntry
	exchangeName string
	channelType  string
	routingKey   string
}

func NewRabbitPublisher[T any](
	log log.LogRusEntry,
	rabbitURL string,
	exchangeName string,
	channelType string,
	routingKey string,
) (RabbitPublisher[T], error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &rabbitPublisher[T]{
		ch:           channel,
		exchangeName: exchangeName,
		channelType:  channelType,
		routingKey:   routingKey,
		log:          log,
	}, nil
}

func (c *rabbitPublisher[T]) exchangeDeclare() error {
	return c.ch.ExchangeDeclare(
		c.exchangeName, // name
		c.channelType,  // type
		false,          // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
}

func (c *rabbitPublisher[T]) Publish(data T) error {
	return c.PublishTo(c.exchangeName, c.routingKey, data)
}

func (c *rabbitPublisher[T]) PublishTo(exchange string, routingKey string, data T) error {
	err := c.exchangeDeclare()
	if err != nil {
		c.log.Error(err)
		return err
	}

	send := publishMessage[T]{
		CorrelationId: c.log.CorrelationId(),
		Message:       data,
	}

	body, err := json.Marshal(send)
	if err != nil {
		c.log.Error(err)
		return err
	}

	err = c.ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Body: body,
		})

	if err != nil {
		c.log.Error(err)
		return err
	}

	c.log.Info("Rabbit publish", string(body))

	return nil
}

type publishMessage[T any] struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       T      `json:"message" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg" `
}
