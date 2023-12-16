package service

import (
	"log"
	"project/modals"
)

type DriverRepository interface {
	GetListTrip() (*[]modals.Trip, error)
	GetTripById(id string) (*modals.Trip, error)
	CancelTrip(id string) error
	AcceptTrip(id string) error
	StartTrip(id string) error
}

type Driver struct {
	repo DriverRepository
}

func NewDriverService(repo DriverRepository) *Driver {
	return &Driver{
		repo: repo,
	}
}

func (d *Driver) GetAllTrips() (*[]modals.Trip, error) {
	trips, err := d.repo.GetListTrip()

	if err != nil {
		log.Fatal(err)
	}

	return trips, nil
}

func (d *Driver) GetIdTrip(id string) (*modals.Trip, error) {
	trip, err := d.repo.GetTripById(id)

	if err != nil {
		log.Fatal(err)
	}

	return trip, nil
}

func (d *Driver) Cancel(id string) error {
	err := d.repo.CancelTrip(id)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (d *Driver) Accept(id string) error {
	err := d.repo.AcceptTrip(id)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (d *Driver) Start(id string) error {
	err := d.repo.StartTrip(id)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
