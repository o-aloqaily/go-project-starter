package kafka

import (
	"testing"

	"github.com/Shopify/sarama"
)

// Test connect is working properly
func TestConnect(t *testing.T) {

	broker := sarama.NewMockBrokerAddr(t, 1, "localhost:9092")
	// Print errors if any
	if broker == nil {
		t.Error("Cannot start mock kafka broker")
	}

	topic := "test-topic"
	mockFetchResponse := sarama.NewMockFetchResponse(t, 1)
	mockFetchResponse.SetMessage(topic, 0, 0, sarama.ByteEncoder([]byte{0x41, 0x42}))

	broker.SetHandlerByMap(map[string]sarama.MockResponse{
		"MetadataRequest": sarama.NewMockMetadataResponse(t).
			SetBroker(broker.Addr(), broker.BrokerID()).
			SetLeader(topic, 0, broker.BrokerID()),
		"OffsetRequest": sarama.NewMockOffsetResponse(t).
			SetOffset(topic, 0, sarama.OffsetOldest, 0).
			SetOffset(topic, 0, sarama.OffsetNewest, 2),
		"FetchRequest": mockFetchResponse,
	})

	consumer := NewConsumer()
	if consumer == nil {
		t.Error("Error creating consumer")
	}

	err := consumer.Connect(&ConnParams{
		BrokerAddress: broker.Addr(),
		Protocol:      "tcp",
		Topic:         topic,
		Partition:     0,
	})

	if err != nil {
		t.Error("Error connecting to the kafka broker: ", err)
	}
}
