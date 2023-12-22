package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	insertDataQuery := `
		INSERT INTO drivers (lat, lng, name, auto) VALUES ($1, $2, $3, $4)
	`
	_, err = db.Exec(insertDataQuery, 123.23, 123.42, "ivan", "toyta")
	_, err = db.Exec(insertDataQuery, 125.23, 121.42, "misha", "bmw")
	_, err = db.Exec(insertDataQuery, 130.0, 135.0, "anna", "honda")
	_, err = db.Exec(insertDataQuery, 140.0, 145.0, "john", "audi")
	_, err = db.Exec(insertDataQuery, 150.0, 155.0, "emily", "mercedes")
	_, err = db.Exec(insertDataQuery, 160.0, 165.0, "alex", "volkswagen")
	_, err = db.Exec(insertDataQuery, 170.0, 175.0, "olivia", "subaru")
	_, err = db.Exec(insertDataQuery, 180.0, 185.0, "liam", "porsche")
	_, err = db.Exec(insertDataQuery, 40.7128, -74.0060, "misha", "toyota")
	_, err = db.Exec(insertDataQuery, 34.0522, -118.2437, "sasha", "bmw")
	_, err = db.Exec(insertDataQuery, 51.509865, -0.118092, "Ya", "honda")
	_, err = db.Exec(insertDataQuery, 48.8566, 2.3522, "Kirill", "audi")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted successfully")

	return &LocationRepository{
		db: db,
	}
}

func (storage *LocationRepository) GetAllDrivers() ([]modals.Driver, error) {
	var drivers []modals.Driver
	err := storage.db.Select(&drivers, "SELECT * FROM drivers")
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (storage *LocationRepository) GetDriverLocationById(driverID string) (*modals.Driver, error) {
	driver := &modals.Driver{}
	err := storage.db.Get(driver, "SELECT * FROM driver WHERE id = $1", driverID)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (storage *LocationRepository) UpdateDriverPosition(driverID string, newLat, newLng float64) (*modals.Driver, error) {
	query := `UPDATE drivers SET lat = $2, lng = $3 WHERE id = $1 RETURNING *`
	driver := &modals.Driver{}
	err := storage.db.Get(driver, query, driverID, newLat, newLng)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
