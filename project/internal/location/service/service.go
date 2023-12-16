package service

import (
	"fmt"
	"project/modals"
)

type LocationRepositoryIn interface {
	GetDrivers() ([]*modals.Driver, error)
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

func (l *Location) GetAllDrivers() ([]*modals.Driver, error) {
	fmt.Print("wait implement of repositore")
	return nil, nil
}

func (l *Location) ChangeDriverById(driver_id string) (*modals.Driver, error) {
	fmt.Print("wait implement of repositore")
	return nil, nil
}
