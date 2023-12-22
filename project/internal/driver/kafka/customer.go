package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
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

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{"127.0.0.1:29092"},
			Topic:          "demo",
			GroupID:        "my-group",
			SessionTimeout: time.Second * 6,
		})
		defer reader.Close()

		for {
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Fatal(err)
			}
			var message TripEvent
			if err := json.Unmarshal(msg.Value, &message); err != nil {
				log.Printf("Failed to decode JSON: %s\n", err)
				continue
			}

			fmt.Printf("Received message: %+v\n", message.Data.To)
			//в бд
			time.Sleep(300 * time.Millisecond)
		}
	}()

	<-ctx.Done()
}
