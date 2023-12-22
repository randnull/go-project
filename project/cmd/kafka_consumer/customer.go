package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

type Teacher struct {
	Lat    float64 `json:"Lat"`
	Lng    float64 `json:"Lng"`
	Radius float64 `json:"radius"`
}

type Driver struct {
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Auto string  `json:"auto"`
}

type Price struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
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

			fmt.Printf("Received message: %+v\n", message.Data.From)

			teacher := Teacher{
				Lat:    message.Data.From.Lat,
				Lng:    message.Data.From.Lng,
				Radius: 1000,
			}

			marshalled, err := json.Marshal(teacher)
			if err != nil {
				log.Fatal(err)
			}

			req, err := http.NewRequest("GET", "http://localhost:1512/drivers", bytes.NewReader(marshalled))
			if err != nil {
				log.Fatal(err)
			}

			client := http.Client{Timeout: 10 * time.Second}

			res, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			var drivers []Driver
			err = json.Unmarshal(body, &drivers)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(drivers)
			for i := 0; i < len(drivers); i++ {
				new_trip := modals.Trip{
					ID:       "1",
					DriverID: drivers[i].ID,
					UserId:   "user_1324",
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
				fmt.Println("INSERT")
				_, err = datab.InsertOne(context.TODO(), new_trip)
				if err != nil {
					log.Fatal(err)
				}
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	<-ctx.Done()
}
