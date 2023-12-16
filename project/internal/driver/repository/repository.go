package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"project/modals"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db_collection *mongo.Collection
}

func NewRepository(URI string) (*Repository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	mongoDb := client.Database("driver")
	collectionTrip := mongoDb.Collection("trip")

	return &Repository{
		db_collection: collectionTrip,
	}, nil
}

func (storage *Repository) GetListTrip() (*modals.Trip, error) {
	fmt.Print("not implement")
	return nil, nil
}

func (storage *Repository) GetTripById(id string) (*modals.Trip, error) {
	var answerTrip modals.Trip

	filter := bson.M{"id": id}

	err := storage.db_collection.FindOne(context.TODO(), filter).Decode(&answerTrip)

	if err != nil {
		log.Fatal(err)
	}

	return &answerTrip, nil
}

func (storage *Repository) CancelTrip(id string) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"status": "Canceled"}}

	_, err := storage.db_collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (storage *Repository) AcceptTrip(id string) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"status": "Accepted"}}

	_, err := storage.db_collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (storage *Repository) StartTrip(id string) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"status": "Started"}}

	_, err := storage.db_collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
