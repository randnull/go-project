package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
	"time"
)

type Price struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type TripData struct {
	TripID  string   `json:"trip_id"`
	OfferID string   `json:"offer_id"`
	Price   Price    `json:"price"`
	Status  string   `json:"status"`
	From    Location `json:"from"`
	To      Location `json:"to"`
}

type TripEvent struct {
	ID              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	DataContentType string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            TripData  `json:"data"`
}

func producer() {
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

	m := TripEvent{
		ID:              "284655d6-0190-49e7-34e9-9b4060acc261",
		Source:          "/trip",
		Type:            "trip.event.created",
		DataContentType: "application/json",
		Time:            time.Date(2023, 11, 9, 17, 31, 0, 0, time.UTC),
		Data: TripData{
			TripID:  "e82c42d6-b86f-4e2a-93a2-858413acb148",
			OfferID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN0cmluZyIsImZyb20iOnsibGF0IjowLCJsbmciOjB9LCJ0byI6eyJsYXQiOjAsImxuZyI6MH0sImNsaWVudF9pZCI6InN0cmluZyIsInByaWNlIjp7ImFtb3VudCI6OTkuOTUsImN1cnJlbmN5IjoiUlVCIn19.fg0Bv2ONjT4r8OgFqJ2tpv67ar7pUih2LhDRCRhWW3c",
			Price: Price{
				Currency: "RUB",
				Amount:   100,
			},
			Status: "DRIVER_SEARCH",
			From: Location{
				Lat: 0,
				Lng: 0,
			},
			To: Location{
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
