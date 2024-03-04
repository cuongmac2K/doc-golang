package main

import (
	"context"
	"crypto/sha1"
	postges_connect "doc/database/kafka/postges-connect"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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
				dataSend := map[string]interface{}{}
				json.Unmarshal(message.Value, &dataSend)
				fmt.Println(dataSend)
				var logs []LogData
				byteLog, _ := json.Marshal(dataSend["ArrLogs"])
				_ = json.Unmarshal(byteLog, &logs)
				fmt.Println("log ", logs)
				arrlogEntry := []LogsInsert{}
				for _, i := range logs {
					id := strconv.Itoa(time.Now().UTC().Nanosecond())
					hash := sha1.Sum([]byte(id))
					uuid := hex.EncodeToString(hash[:])
					logEntry := LogsInsert{
						ID:         id,
						UUID:       uuid,
						Agent_UUID: dataSend["AgentUuid"].(string),
						Data:       i.Log,
						Created_At: time.Now().UTC(),
					}
					arrlogEntry = append(arrlogEntry, logEntry)
				}
				err := InsertMany(context.Background(), postges_connect.Postgres, "logs", arrlogEntry)
				time.Sleep(time.Millisecond / 2)
				if err != nil {
					fmt.Println("", err)
				}
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
	err := postges_connect.GetDBConnection()
	topic := "logs"
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
			if err := consumerGroup.Consume(ctx, []string{topic}, &consumer); err != nil {
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

func InsertMany(ctx context.Context, conn *pgx.Conn, tableName string, logs []LogsInsert) error {
	// Tạo một transaction
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Chuẩn bị câu lệnh INSERT
	stmtName := fmt.Sprintf("insert_statement_%s", tableName)
	_, err = tx.Prepare(ctx, stmtName, fmt.Sprintf("INSERT INTO %s (id, uuid,agent_uuid,data,created_at) VALUES ($1, $2, $3, $4, $5)", tableName))
	if err != nil {
		return err
	}
	// Chèn từng dòng vào PostgreSQL
	for _, log := range logs {
		_, err := conn.Exec(ctx, stmtName, log.ID, log.UUID, log.Agent_UUID, log.Data, log.Created_At)
		if err != nil {
			return err
		}
	}
	// Commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
