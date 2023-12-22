package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"project/modals"
	"strconv"
	"time"
)

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

	m := modals.KafkaMessage{
		ID:              "284655d6-0190-49e7-34e9-9b4060acc261",
		Source:          "/trip",
		Type:            "trip.event.created",
		DataContentType: "application/json",
		Time:            time.Date(2023, 11, 9, 17, 31, 0, 0, time.UTC),
		Data: modals.TripData{
			TripID:  "e82c42d6-b86f-4e2a-93a2-858413acb148",
			OfferID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN0cmluZyIsImZyb20iOnsibGF0IjowLCJsbmciOjB9LCJ0byI6eyJsYXQiOjAsImxuZyI6MH0sImNsaWVudF9pZCI6InN0cmluZyIsInByaWNlIjp7ImFtb3VudCI6OTkuOTUsImN1cnJlbmN5IjoiUlVCIn19.fg0Bv2ONjT4r8OgFqJ2tpv67ar7pUih2LhDRCRhWW3c",
			Price: modals.Money{
				Currency: "RUB",
				Amount:   100,
			},
			Status: "DRIVER_SEARCH",
			From: modals.Latlngtiteral{
				Lat: 0,
				Lng: 0,
			},
			To: modals.Latlngtiteral{
				Lat: 0,
				Lng: 0,
			},
		},
	}
	//fmt.Println("success")
	messageBytes, err := json.Marshal(m)
	err = writer.WriteMessages(ctx, kafka.Message{Key: []byte(strconv.Itoa(0)), Value: messageBytes})
	if err != nil {
		log.Fatal(err)
	}
}
