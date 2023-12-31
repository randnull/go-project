package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/modals"
	"syscall"
	"time"
)

func db() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("driver")
	collection := db.Collection("trip")
	return collection
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()
	datab := db()
	client := http.Client{Timeout: 10 * time.Second}
	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{"127.0.0.1:29092"},
			Topic:          "getter",
			GroupID:        "my-group",
			SessionTimeout: time.Second * 6,
		})
		defer reader.Close()

		for {
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Fatal(err)
			}

			var message modals.KafkaMessage
			err = json.Unmarshal(msg.Value, &message)
			if err != nil {
				log.Fatal(err)
			}

			resp := modals.Req{
				Lat:    message.Data.From.Lat,
				Lng:    message.Data.From.Lng,
				Radius: 1000,
			}

			marshalled, err := json.Marshal(resp)
			if err != nil {
				log.Fatal(err)
			}

			req, err := http.NewRequest("GET", "http://localhost:8828/drivers", bytes.NewReader(marshalled))
			if err != nil {
				log.Fatal(err)
			}

			res, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			var drivers []modals.Driver
			err = json.Unmarshal(body, &drivers)
			if err != nil {
				log.Fatal(err)
			}

			for i := 0; i < len(drivers); i++ {
				new_trip := modals.Trip{
					ID:       message.ID,
					DriverID: drivers[i].ID,
					From: modals.Latlngtiteral{
						Lat: message.Data.From.Lat,
						Lng: message.Data.From.Lng,
					},
					To: modals.Latlngtiteral{
						Lat: message.Data.To.Lat,
						Lng: message.Data.To.Lng,
					},
					Price: modals.Money{
						Amount:   message.Data.Price.Amount,
						Currency: message.Data.Price.Currency,
					},
					Status: "DRIVER_GET_REQUEST",
				}
				_, err := datab.InsertOne(context.TODO(), new_trip)
				if err != nil {
					log.Fatal(err)
				}
			}
			log.Println("[INFO]: Get new trip!")
			time.Sleep(300 * time.Millisecond)
		}
	}()

	<-ctx.Done()
}
