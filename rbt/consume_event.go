package rbt

import (
	"encoding/json"

	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/strs"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

func ConsumeRabbitEvent[T any](
	fluentUrl string,
	rabbitUrl string,
	serverName string,
	exchangeName string,
	channelType string,
	queueName string,
	routingKey string,
	processIncomingMessage func(log.LogRusEntry, *InputMessage[T]),
) error {
	logger := RbtLogger(fluentUrl, serverName, uuid.NewV4().String())

	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		logger.Error(err)
		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		exchangeName, // name
		channelType,  // type
		false,        // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,   // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil)
	if err != nil {
		logger.Error(err)
		return err
	}

	mgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	go func() {
		for d := range mgs {
			newMessage := &InputMessage[T]{}
			body := d.Body

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				l := RbtLogger(fluentUrl, "cartgo", getCorrelationId(newMessage)).
					WithField(log.LOG_FIELD_RABBIT_ACTION, "consume").
					WithField(log.LOG_FIELD_RABBIT_EXCHANGE, exchangeName).
					WithField(log.LOG_FIELD_RABBIT_QUEUE, queueName)

				l.Info("Incoming :", string(body))

				processIncomingMessage(l, newMessage)

				if err := d.Ack(false); err != nil {
					l.Info("Failed ACK :", strs.ToJson(newMessage), err)
				} else {
					l.Info("Consumed :", strs.ToJson(newMessage))
				}
			} else {
				logger.Error(err)
			}
		}
	}()

	logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

type InputMessage[T any] struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       T
}

func getCorrelationId[T any](c *InputMessage[T]) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}
