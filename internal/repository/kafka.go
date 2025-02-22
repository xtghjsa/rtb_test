package repository

import (
	"context"
	"encoding/json"
	"log"
	"test_project/internal/encoding"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	Writer *kafka.Writer
}

// Initiates kafka reader
func NewReaderKafka() *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:5050"},
		Topic:   "test",
	})
	return reader
}

// Returns message from kafka
func ReadKafka(reader *kafka.Reader) (KafkaMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	msg, err := reader.ReadMessage(ctx)
	cancel()
	if err != nil {
		if err == context.DeadlineExceeded {
			time.Sleep(500 * time.Millisecond)
			return KafkaMessage{}, nil
		}
		return KafkaMessage{}, err
	}
	var msgResp KafkaMessage
	err = json.Unmarshal(msg.Value, &msgResp)
	if err != nil {
		log.Printf("Error unmarshalling message: %v\n", err)
		return KafkaMessage{}, err
	}
	return msgResp, nil
}

// Initiates kafka writer
func NewWriterKafka() *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:5050"},
		Topic:   "test",
	})
	return writer
}

type KafkaMessage struct {
	BidID     string
	Price     int64
	EventType string
}

// Writes message to kafka
func (w *KafkaWriter) WriteKafka(decoded *encoding.Tracking, eventType string) error {
	msg := KafkaMessage{
		BidID:     decoded.BidID,
		Price:     decoded.Price,
		EventType: eventType,
	}

	msgMarshalled, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = w.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: msgMarshalled,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
