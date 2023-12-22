package repository

import (
	"context"
	"fmt"
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
	fmt.Println("Testing data begin...")
	//Testing data begin... !!!!! NEED TO BE REMOVED IN PRODUCT !!!!
	trips := []interface{}{
		modals.Trip{
			ID:       "1",
			DriverID: "dwdwdadawdwa",
			UserId:   "user_1",
			From: modals.Latlngtiteral{
				Lat: 40.7128,
				Lng: -74.0060,
			},
			To: modals.Latlngtiteral{
				Lat: 34.0522,
				Lng: -118.2437,
			},
			Price: modals.Money{
				Amount:   25.0,
				Currency: "USD",
			},
			Status: "completed",
		},
		modals.Trip{
			ID:       "2",
			DriverID: "driver_2",
			UserId:   "user_2",
			From: modals.Latlngtiteral{
				Lat: 34.0522,
				Lng: -118.2437,
			},
			To: modals.Latlngtiteral{
				Lat: 37.7749,
				Lng: -122.4194,
			},
			Price: modals.Money{
				Amount:   30.0,
				Currency: "USD",
			},
			Status: "in_progress",
		},
	}

	_, err = collectionTrip.InsertMany(context.TODO(), trips)
	if err != nil {
		fmt.Println("bad")
	}
	//Testing data end
	fmt.Println("Testing data end")
	fmt.Println("Database is ready!")
	return &DriverRepository{
		dbCollection: collectionTrip,
	}
}

func (storage *DriverRepository) GetListTrip(user_id string) (*[]modals.Trip, error) {
	log.Fatal("err")
	fmt.Print("not implement")
	return nil, nil
}

//func (storage *DriverRepository) PutNewTrip()

func (storage *DriverRepository) GetTripById(user_id string, trip_id string) (*modals.Trip, error) {
	fmt.Println("GetTripById with user_id: trip_id")
	fmt.Println(user_id, trip_id)

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

	return &answerTrip, nil
}

func (storage *DriverRepository) CancelTrip(user_id string, trip_id string, reason string) error {
	objectTripId, err := primitive.ObjectIDFromHex(trip_id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectTripId}
	update := bson.M{"$set": bson.M{"status": "Canceled"}}

	idTripCorrect := true // TODO()
	idUserCorrect := true // TODO()

	if !idUserCorrect {
		return errors.UserNotFoundInBD
	}

	if !idTripCorrect {
		return errors.TripNotFound
	}

	_, err = storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return errors.FailedToCancelTrip
	}

	return err
}

func (storage *DriverRepository) EndTrip(user_id string, trip_id string) error {
	objectTripId, err := primitive.ObjectIDFromHex(trip_id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectTripId}
	update := bson.M{"$set": bson.M{"status": "Ended"}}

	idTripCorrect := true // TODO()
	idUserCorrect := true // TODO()

	if !idUserCorrect {
		return errors.UserNotFoundInBD
	}

	if !idTripCorrect {
		return errors.TripNotFound
	}

	_, err = storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return errors.FailedToEndTrip
	}

	return err
}

func (storage *DriverRepository) AcceptTrip(user_id string, trip_id string) error {
	objectTripId, err := primitive.ObjectIDFromHex(trip_id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectTripId}
	update := bson.M{"$set": bson.M{"status": "Accepted"}}

	idTripCorrect := true // TODO()
	idUserCorrect := true // TODO()

	if !idUserCorrect {
		return errors.UserNotFoundInBD
	}

	if !idTripCorrect {
		return errors.TripNotFound
	}

	_, err = storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return errors.FailedToAcceptTrip
		log.Fatal(err)
	}

	return err
}

func (storage *DriverRepository) StartTrip(user_id string, trip_id string) error {
	objectTripId, err := primitive.ObjectIDFromHex(trip_id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectTripId}
	update := bson.M{"$set": bson.M{"status": "Started"}}

	idTripCorrect := true // TODO()
	idUserCorrect := true // TODO()

	if !idUserCorrect {
		return errors.UserNotFoundInBD
	}

	if !idTripCorrect {
		return errors.TripNotFound
	}

	_, err = storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return errors.FailedToStartTrip
		log.Fatal(err)
	}

	return err
}

func (storage *DriverRepository) PutNewTrip(trip modals.Trip) error {
	//trip.ID = primitive.NewObjectID().Hex()
	//log.Fatal("we")
	_, err := storage.dbCollection.InsertOne(context.TODO(), trip)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
