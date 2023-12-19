package errors

import "errors"

var InvalidLocation = errors.New("Invalid location")
var InvalidDriverId = errors.New("Invalid driver ID")
var DriversNotFound = errors.New("Drivers not found")
var FailedToUpdatePosition = errors.New("Failed to update position")
