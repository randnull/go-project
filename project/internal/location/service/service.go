package service

import (
	"project/modals"
)

type LocationRepositoryIn interface {
	GetDrivers() ([]modals.Driver, error)
	GetDriverLocationById(driver_id string) (*modals.Driver, error)
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

func (l *Location) ChangeDriverById(driver_id string) (*modals.Driver, error) {
	driver, err := l.repo.GetDriverLocationById(driver_id)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
