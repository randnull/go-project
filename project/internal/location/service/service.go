package service

import (
	"project/modals"
)

type LocationRepositoryIn interface {
	GetDrivers() ([]modals.Driver, error)
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
	drivers, err := l.repo.GetDrivers()
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (l *Location) ChangeDriverById(driverID string) (*modals.Driver, error) {
	driver, err := l.repo.GetDriverLocationById(driverID)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (l *Location) UpdateDriverPosition(driverID string, newLat, newLng float64) (*modals.Driver, error) {
	return nil, nil
}
