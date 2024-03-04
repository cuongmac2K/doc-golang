package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	produceMessage([]string{"10.3.52.78:9093"}, "logs", "xin chao xin chao")
}

func produceMessage(brokers []string, topic string, message string) error {
	producerConfig := sarama.NewConfig()
	producerConfig.Version = sarama.V2_1_0_0
	producerConfig.Net.SASL.Enable = true
	producerConfig.Net.SASL.User = "admin"
	producerConfig.Net.SASL.Password = "admin-secret"
	producerConfig.Net.SASL.Handshake = true
	producerConfig.Net.TLS.Enable = false
	producerConfig.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, producerConfig)
	if err != nil {
		return fmt.Errorf("failed to create producer: %w", err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
