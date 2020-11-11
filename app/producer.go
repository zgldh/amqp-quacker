package app

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Producer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	tag        string
	done       chan error
	dryRun     bool
}

func NewProducer(amqpURI, exchange, exchangeType, key, ctag string, dryrun bool) (*Producer, error) {
	p := &Producer{
		connection: nil,
		channel:    nil,
		tag:        ctag,
		done:       make(chan error),
		dryRun:     dryrun,
	}

	var err error

	if dryrun {
		fmt.Println("Dry Run")
	} else {

		fmt.Printf("Connecting to %s", amqpURI)
		p.connection, err = amqp.Dial(amqpURI)
		if err != nil {
			return nil, fmt.Errorf("Dial: ", err)
		}

		fmt.Printf("Getting Channel ")
		p.channel, err = p.connection.Channel()
		if err != nil {
			return nil, fmt.Errorf("Channel: ", err)
		}

		fmt.Printf("Declaring Exchange (%s)", exchange)
		if len(exchange) > 0 {
			if err := p.channel.ExchangeDeclare(
				exchange,     // name
				exchangeType, // type
				true,         // durable
				false,        // auto-deleted
				false,        // internal
				false,        // noWait
				nil,          // arguments
			); err != nil {
				return nil, fmt.Errorf("Exchange Declare: %s", err)
			}
		}

	}
	return p, nil
}

func (p *Producer) Publish(exchange, routingKey, body string) error {
	if p.dryRun {
		return nil
	}

	if err := p.channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: ", err)
	}

	return nil
}

// One would typically keep a channel of publishings, a sequence number, and a
// set of unacknowledged sequence numbers and loop until the publishing channel
// is closed.
func confirmOne(ack, nack chan uint64) {
	fmt.Printf("waiting for confirmation of one publishing")

	select {
	case tag := <-ack:
		fmt.Printf("confirmed delivery with delivery tag: %d", tag)
	case tag := <-nack:
		fmt.Printf("failed delivery of delivery tag: %d", tag)
	}
}
