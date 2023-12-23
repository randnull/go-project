package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"project/internal/driver/errors"

	"project/modals"

	"go.mongodb.org/mongo-driver/mongo"
)

type DriverRepository struct {
	dbCollection *mongo.Collection
}

func NewDriverRepository(URI string) *DriverRepository {
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
	log.Println("Database is ready!")
	return &DriverRepository{
		dbCollection: collectionTrip,
	}
}

func (storage *DriverRepository) GetListTrip(user_id string) (*[]modals.Trip, error) {
	cur, err := storage.dbCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var trips []modals.Trip

	for cur.Next(context.TODO()) {
		var trip modals.Trip
		err := cur.Decode(&trip)
		if err != nil {
			return nil, err
		}
		trips = append(trips, trip)
	}
	log.Println("GetListTrip - success")

	return &trips, nil
}

func (storage *DriverRepository) GetTripById(user_id string, trip_id string) (*modals.Trip, error) {
	log.Println("GetTripById with user_id: trip_id")
	log.Println(user_id, trip_id)

	objectTripId, err := primitive.ObjectIDFromHex(trip_id)

	if err != nil {
		log.Fatal(err)
	}

	var answerTrip modals.Trip

	filter := bson.M{"_id": objectTripId}

	err = storage.dbCollection.FindOne(context.TODO(), filter).Decode(&answerTrip)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("GetTripById - success")

	return &answerTrip, nil
}

func (storage *DriverRepository) CancelTrip(user_id string, trip_id string, reason string) error {
	objectTripId, err := primitive.ObjectIDFromHex(trip_id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectTripId}
	update := bson.M{"$set": bson.M{"status": "Canceled"}}

	_, err = storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return errors.FailedToCancelTrip
	}
	log.Println("CancelTrip - success")

	return err
}

func (storage *DriverRepository) EndTrip(user_id string, trip_id string) error {
	objectTripId, err := primitive.ObjectIDFromHex(trip_id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectTripId}
	update := bson.M{"$set": bson.M{"status": "Ended"}}

	_, err = storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return errors.FailedToEndTrip
	}
	log.Println("EndTrip - success")
	return err
}

func (storage *DriverRepository) AcceptTrip(user_id string, trip_id string) error {
	objectTripId, err := primitive.ObjectIDFromHex(trip_id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectTripId}
	update := bson.M{"$set": bson.M{"status": "Accepted"}}

	_, err = storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return errors.FailedToAcceptTrip
		log.Fatal(err)
	}
	log.Println("AcceptTrip - success")
	return err
}

func (storage *DriverRepository) StartTrip(user_id string, trip_id string) error {
	objectTripId, err := primitive.ObjectIDFromHex(trip_id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectTripId}
	update := bson.M{"$set": bson.M{"status": "Started"}}

	_, err = storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return errors.FailedToStartTrip
		log.Fatal(err)
	}
	log.Println("StartTrip - success")
	return err
}

func (storage *DriverRepository) PutNewTrip(trip modals.Trip) error {
	_, err := storage.dbCollection.InsertOne(context.TODO(), trip)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("PutNewTrip - success")
	return nil
}

func (storage *DriverRepository) AcceptFromDriver(id string) error {

	objectTripId, err := primitive.ObjectIDFromHex(id)
	println(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectTripId}
	result := storage.dbCollection.FindOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var trip modals.Trip
	err = result.Decode(&trip)
	id_req := trip.ID
	if err != nil {
		return err
	}
	filter = bson.M{"id": id_req, "_id": bson.M{"$ne": objectTripId}}
	_, err = storage.dbCollection.UpdateMany(context.TODO(), filter, bson.M{"$set": bson.M{"status": "Canceled"}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Driver accept")
	return nil
}
