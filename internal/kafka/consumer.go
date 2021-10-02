// Package kafka implements the Kafka consumer to consume messages from a broker
package kafka

import (
	"context"
	"sync"

	kafka "github.com/segmentio/kafka-go"
)

// Consumer is the kafka consumer interface
type Consumer interface {
	Connect(connParams interface{}) error
}

type consumer struct {
	Conn *kafka.Conn
}

// ConnParams is the struct to specify the parameters required to
// connect to the Kafka Broker.
type ConnParams struct {
	BrokerAddress string
	Protocol      string
	Topic         string
	Partition     int
}

// consumer is a singleton consumer instance
var c *consumer
var once sync.Once

// NewConsumer is the constructor for the kafkaConsumer type
func NewConsumer() Consumer {
	once.Do(func() {
		c = &consumer{}
	})
	return c
}

// Connect is the function to connect to the Kafka broker
func (c *consumer) Connect(connParams interface{}) error {
	params := connParams.(*ConnParams)
	if c.Conn == nil {
		// TODO inject a logger into the package and log an error here
	}

	// Dial the kafka broker with the connection parameters
	conn, err := kafka.DialLeader(
		context.Background(),
		params.Protocol,
		params.BrokerAddress,
		params.Topic,
		params.Partition,
	)

	// Print errors if any
	if err != nil {
		// TODO inject a logger into the package and log the error here
		return err
	}

	// If connection is successful assign the conn object to the singleton consumer
	c.Conn = conn

	// Return nil if no error occurs
	return nil
}
