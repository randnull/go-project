package errors

import "errors"

var FailedToCancelTrip = errors.New("Failed to cancel trip")
var FailedToEndTrip = errors.New("Failed to end trip")
var FailedToAcceptTrip = errors.New("Failed to accept trip")
var FailedToStartTrip = errors.New("Failed to start trip")
var FailedToGetTripsList = errors.New("Failed to get trips")
var FailedToPutNewTrip = errors.New("Failed to get trip by ID")

var InvalidTripID = errors.New("Invalid trip id")
