package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Consumer struct {
	ready         chan bool
	ctx           context.Context
	tracingClient *http.Client
}

func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

type LogData struct {
	Date float64 `json:"date"`
	Log  string  `json:"log"`
}

func convertNameAgent(data []byte) map[string]interface{} {
	a := map[string]interface{}{}
	_ = json.Unmarshal(data, &a)
	return a
}

type LogsInsert struct {
	ID         string
	UUID       string
	Agent_UUID string
	Data       string
	Created_At time.Time
	//Update_At  time.Time
	//Delete_At  time.Time
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				fmt.Println("message channel was closed")
				//transaction.End()
				return nil
			} else {
				fmt.Println(string(message.Value))
				if strings.Contains(string(message.Value), "err") {
					//produceMessage([]string{"10.3.52.78:9093"}, "devops-bizfly-activity-log-err-staging", message.Value)
				}
				session.MarkMessage(message, "")
			}
		case <-session.Context().Done():
			return nil
		}
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}

func main() {
	topic := "devops-bizfly-activity-log-staging"
	topic1 := "devops-bizfly-activity-log-err-staging"
	brokers := []string{"10.3.52.78:9093"}

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Version = sarama.V2_1_0_0
	kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	kafkaConfig.Net.SASL.Enable = true
	kafkaConfig.Net.SASL.User = "admin"
	kafkaConfig.Net.SASL.Password = "admin-secret"
	kafkaConfig.Net.SASL.Handshake = true
	kafkaConfig.Net.TLS.Enable = false

	ctx, cancel := context.WithCancel(context.Background())
	consumer := Consumer{
		ready: make(chan bool),
		ctx:   ctx,
	}
	consumerGroup, err := sarama.NewConsumerGroup(brokers, fmt.Sprintf("%v-group-%v", topic, "dev"), kafkaConfig)
	if err != nil {
		log.Fatalf("failed to create consumer group: %v", err)
	}
	defer consumerGroup.Close()

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := consumerGroup.Consume(ctx, []string{topic, topic1}, &consumer); err != nil {
				if err == sarama.ErrClosedConsumerGroup {
					return
				}
				log.Printf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	keepRunning := true
	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("Terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("Terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(consumerGroup, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()

	log.Println("Waiting for the consumer group to finish processing...")
	<-time.After(10 * time.Second)
	log.Println("Shutdown complete")
}
func produceMessage(brokers []string, topic string, message []byte) error {
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
