package app

import (
	"fmt"
	"strconv"
	"time"
)

// QuackerConfig - Configuration of AMQP
//server
type QuackerConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Exchange string
	Topic    string
	Interval string // Interval - Seconds between two publish
	DataFile string // DataFile - Data template file path
}

// Quacker - The quacker class.
type Quacker struct {
	config  QuackerConfig
	builder DataBuilder
}

// NewQuacker - Create a new Quacker object
func NewQuacker(config QuackerConfig) Quacker {
	return Quacker{
		config:  config,
		builder: NewDataBuilder(DataBuilderConfig{Path: config.DataFile}),
	}
}

// Close - Close the quacker mission
func (q *Quacker) Close() {
}

// Start - Start to subscribe MQTT and tranfer data into Pgsql
func (q *Quacker) Start() error {
	fmt.Printf("Quacker starting...\n")

	reliable := true
	ctag := "amqp-quacker"
	key := q.config.Topic
	exchangeType := "direct"
	exchange := q.config.Exchange
	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		q.config.Username,
		q.config.Password,
		q.config.Host,
		q.config.Port,
	)
	producer, err := NewProducer(amqpURI, exchange, exchangeType, key, ctag, reliable)
	if err != nil {
		return err
	}

	interval, err := strconv.Atoi(q.config.Interval)
	if err != nil {
		return err
	}

	payload := ""

	fmt.Printf("AMQP server %s:%s\n", q.config.Host, q.config.Port)
	fmt.Printf("Exchange: %s\n", q.config.Exchange)
	fmt.Println("Publisher Started to: " + q.config.Topic)
	for true {
		fmt.Printf("%s ---- Publish ----\n", time.Now())
		payload = q.getPayload()
		producer.Publish(q.config.Exchange, q.config.Topic, payload)
		time.Sleep(time.Millisecond * time.Duration(interval))
	}

	fmt.Println("Publisher Disconnected")

	return nil
}

// getPayload - Get payload.
func (q *Quacker) getPayload() string {
	payload, err := q.builder.Make()
	if err != nil {
		panic(err)
	}
	return payload
}
