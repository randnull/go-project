package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"project/internal/driver/errors"

	"project/modals"

	"go.mongodb.org/mongo-driver/mongo"
)

type DriverRepository struct {
	dbCollection *mongo.Collection
}

func NewDriverRepository() *DriverRepository {
	context.TODO()
	return nil
}

//func NewRepository(URI string) (*Repository, error) {
//	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.Connect(context.TODO())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.Ping(context.TODO(), nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	mongoDb := client.Database("driver")
//	collectionTrip := mongoDb.Collection("trip")
//
//	return &Repository{
//		dbCollection: collectionTrip,
//	}, nil
//}

func (storage *DriverRepository) GetListTrip(user_id string, trip_id string) (*[]modals.Trip, error) {
	fmt.Print("not implement")
	return nil, nil
}

func (storage *DriverRepository) GetTripById(user_id string, trip_id string) (*modals.Trip, error) {
	var answerTrip modals.Trip

	filter := bson.M{"id": trip_id}

	err := storage.dbCollection.FindOne(context.TODO(), filter).Decode(&answerTrip)

	if err != nil {
		log.Fatal(err)
	}

	return &answerTrip, nil
}

func (storage *DriverRepository) CancelTrip(user_id string, trip_id string, reason string) error {
	filter := bson.M{"id": trip_id}
	update := bson.M{"$set": bson.M{"status": "Canceled"}}

	idTripCorrect := true // TODO()
	idUserCorrect := true // TODO()

	if !idUserCorrect {
		return errors.UserNotFoundInBD
	}

	if !idTripCorrect {
		return errors.TripNotFound
	}

	_, err := storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return errors.FailedToCancelTrip
	}

	return err
}

func (storage *DriverRepository) EndTrip(user_id string, trip_id string) error {
	filter := bson.M{"id": trip_id}
	update := bson.M{"$set": bson.M{"status": "Ended"}}

	idTripCorrect := true // TODO()
	idUserCorrect := true // TODO()

	if !idUserCorrect {
		return errors.UserNotFoundInBD
	}

	if !idTripCorrect {
		return errors.TripNotFound
	}

	_, err := storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return errors.FailedToEndTrip
	}

	return err
}

func (storage *DriverRepository) AcceptTrip(user_id string, trip_id string) error {
	filter := bson.M{"id": trip_id}
	update := bson.M{"$set": bson.M{"status": "Accepted"}}

	idTripCorrect := true // TODO()
	idUserCorrect := true // TODO()

	if !idUserCorrect {
		return errors.UserNotFoundInBD
	}

	if !idTripCorrect {
		return errors.TripNotFound
	}

	_, err := storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return errors.FailedToAcceptTrip
		log.Fatal(err)
	}

	return err
}

func (storage *DriverRepository) StartTrip(user_id string, trip_id string) error {
	filter := bson.M{"id": trip_id}
	update := bson.M{"$set": bson.M{"status": "Started"}}

	idTripCorrect := true // TODO()
	idUserCorrect := true // TODO()

	if !idUserCorrect {
		return errors.UserNotFoundInBD
	}

	if !idTripCorrect {
		return errors.TripNotFound
	}

	_, err := storage.dbCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return errors.FailedToStartTrip
		log.Fatal(err)
	}

	return err
}
