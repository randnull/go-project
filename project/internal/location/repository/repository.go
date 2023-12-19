package repository

import (
	"context"
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

	query := `
		CREATE TABLE IF NOT EXISTS drivers (
			id SERIAL PRIMARY KEY,
			lat DOUBLE PRECISION,
			lng DOUBLE PRECISION,
			name VARCHAR(255),
			auto VARCHAR(255)
		)
	`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	return &LocationRepository{
		db: db,
	}
}

func (storage *LocationRepository) GetDrivers() ([]modals.Driver, error) {
	var drivers []modals.Driver
	err := storage.db.Select(&drivers, "SELECT * FROM drivers")
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (storage *LocationRepository) GetDriverLocationById(driverID int) (*modals.Driver, error) {
	driver := &modals.Driver{}
	err := storage.db.Get(driver, "SELECT * FROM driver WHERE id = $1", driverID)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
