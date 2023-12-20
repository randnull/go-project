package service

import (
	"log"
	"project/modals"
)

type DriverRepository interface {
	GetListTrip(user_id string) (*[]modals.Trip, error)
	GetTripById(user_id string, trip_id string) (*modals.Trip, error)
	CancelTrip(user_id string, trip_id string, reason string) error
	AcceptTrip(user_id string, trip_id string) error
	StartTrip(user_id string, trip_id string) error
	EndTrip(user_id string, trip_id string) error
}

type Driver struct {
	repo DriverRepository
}

func NewDriverService(repo DriverRepository) *Driver {
	return &Driver{
		repo: repo,
	}
}

func (d *Driver) GetAllTrips(user_id string) (*[]modals.Trip, error) {
	trips, err := d.repo.GetListTrip(user_id)

	if err != nil {
		log.Fatal(err)
	}

	return trips, nil
}

func (d *Driver) GetIdTrip(user_id string, trip_id string) (*modals.Trip, error) {
	trip, err := d.repo.GetTripById(user_id, trip_id)

	if err != nil {
		log.Fatal(err)
	}

	return trip, nil
}

func (d *Driver) Cancel(user_id string, trip_id string, reason string) error {
	err := d.repo.CancelTrip(user_id, trip_id, reason)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (d *Driver) End(user_id string, trip_id string) error {
	err := d.repo.EndTrip(user_id, trip_id)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (d *Driver) Accept(user_id string, trip_id string) error {
	err := d.repo.AcceptTrip(user_id, trip_id)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (d *Driver) Start(user_id string, trip_id string) error {
	err := d.repo.StartTrip(user_id, trip_id)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
