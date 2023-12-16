package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"project/modals"
)

type LocationRepository struct {
	db *sqlx.DB
}

func NewLocationRepository() *LocationRepository {
	db, err := sqlx.Open("postgres", "postgres://postgres:admin@localhost:5432/driver?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	err = db.PingContext(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	// + goose для миграций

	return &LocationRepository{
		db: db,
	}
}

func (storage *LocationRepository) GetDrivers() ([]*modals.Driver, error) {
	fmt.Print("not implement")
	return nil, nil
}

func (storage *LocationRepository) GetDriverLocationById(driverID int) (*modals.Driver, error) {
	driver := &modals.Driver{}
	err := storage.db.Get(driver, "SELECT * FROM driver WHERE id = $1", driverID)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
