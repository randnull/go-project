package service

import (
	"project/modals"
)

type LocationRepositoryIn interface {
	GetAllDrivers() ([]modals.Driver, error)
	GetDriverLocationById(driverID string) (*modals.Driver, error)
	UpdateDriverPosition(driverID string, newLat, newLng float64) (*modals.Driver, error)
}

type Location struct {
	repo LocationRepositoryIn
}

func NewDriverService(repo LocationRepositoryIn) *Location {
	return &Location{
		repo: repo,
	}
}

func (l *Location) GetAllDrivers() ([]modals.Driver, error) {
	drivers, err := l.repo.GetAllDrivers()
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (l *Location) UpdateDriverPosition(driverID string, newLat, newLng float64) (*modals.Driver, error) {
	driver, err := l.repo.UpdateDriverPosition(driverID, newLat, newLng)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
