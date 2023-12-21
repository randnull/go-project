package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
	"time"
)

type TripData struct {
	TripID   string `json:"trip_id"`
	DriverID string `json:"driver_id"`
}

type MessagePayload struct {
	ID              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	DataContentType string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            TripData  `json:"data"`
}

func main() {
	ctx := context.Background()

	logger := log.Default()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:     []string{"127.0.0.1:29092"},
		Topic:       "demo",
		Async:       true,
		Logger:      kafka.LoggerFunc(logger.Printf),
		ErrorLogger: kafka.LoggerFunc(logger.Printf),
		BatchSize:   2000,
	})
	defer writer.Close()

	m := MessagePayload{
		ID:              "284655d6-0190-49e7-34e9-9b4060acc261",
		Source:          "/driver",
		Type:            "trip.command.accept",
		DataContentType: "application/json",
		Time:            time.Now(),
		Data: TripData{
			TripID:   "284655d6-0190-49e7-34e9-9b4060acc260",
			DriverID: "42d6142f-c27e-4c73-bd3a-0051c440aecb",
		},
	}

	messageBytes, err := json.Marshal(m)
	err = writer.WriteMessages(ctx, kafka.Message{Key: []byte(strconv.Itoa(0)), Value: messageBytes})
	if err != nil {
		log.Fatal(err)
	}
}
