package errors

import "errors"

var InvalidLocation = errors.New("Invalid location")
var DriversNotFound = errors.New("Drivers not found")

var BadRequest = errors.New("Bad request")
var InternalServerError = errors.New("Internal server error")
