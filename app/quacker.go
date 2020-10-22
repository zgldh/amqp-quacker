package app

import (
	"fmt"
	"log"
	"math"
	"os/exec"
	"strconv"
	"time"
)

// QuackerConfig - Configuration of RabbitMQ
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

// sendMessage - Send message to amqp server
func (q *Quacker) sendMessage(message string) {
	command := fmt.Sprintf(
		"amqpc -u amqp://%s:%s@%s:%s/ -g 1 -i %s -n 1 -p %s %s %s",
		q.config.Username,
		q.config.Password,
		q.config.Host,
		q.config.Port,
		q.config.Interval,
		q.config.Exchange,
		q.config.Topic,
		message,
	)

	cmd := exec.Command(command)
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error %v", err)
	}
}

// Start - Start to subscribe MQTT and tranfer data into Pgsql
func (q *Quacker) Start() error {
	fmt.Printf("Quacker starting...\n")

	payload := ""

	fmt.Printf("RabbitMQ server %s:%s\n", q.config.Host, q.config.Port)
	fmt.Printf("Exchange: %s\n", q.config.Exchange)
	fmt.Println("Publisher Started to: " + q.config.Topic)
	for true {
		fmt.Printf("%s ---- Publish ----\n", time.Now())
		payload = q.getPayload()
		q.sendMessage(payload)
		fmt.Println(payload)
		time.Sleep(time.Second * time.Duration(interval))
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
