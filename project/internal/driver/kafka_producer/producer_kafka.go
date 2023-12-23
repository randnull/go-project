package kafka_producer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"project/modals"
	"strconv"
	"time"
)

func Produce_data(trip_id string, act string, trip *modals.Trip) {
	ctx := context.Background()

	logger := log.Default()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:     []string{"127.0.0.1:29092"},
		Topic:       "setter",
		Async:       true,
		Logger:      kafka.LoggerFunc(logger.Printf),
		ErrorLogger: kafka.LoggerFunc(logger.Printf),
		BatchSize:   2000,
	})
	defer writer.Close()

	type_now := "trip.event." + act

	m := modals.TripAnswer{
		ID:              trip.ID,
		Source:          "/trip",
		Type:            type_now,
		DataContentType: "application/json",
		Time:            time.Time{},
		Data: modals.TripAnswerData{
			TripID: trip_id,
		},
	}

	messageBytes, err := json.Marshal(m)
	err = writer.WriteMessages(ctx, kafka.Message{Key: []byte(strconv.Itoa(0)), Value: messageBytes})
	if err != nil {
		log.Fatal(err)
	}
}
