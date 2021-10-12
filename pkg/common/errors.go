package common

import "errors"

// Err* is used to throw errors while using the SDK
var (
	ErrVersionNotAvailable  = errors.New("invalid version")
	ErrEndpointNotAvailable = errors.New(
		"invalid endpoint, please verify if the endpoint exists in the specified version")
)

var (
	ErrInvalidInput = Error{
		ErrorCode:        "400",
		ErrorDescription: "invalid JSON input",
	}

	ErrInvalidResponseFromAPI = Error{
		ErrorCode:        "500",
		ErrorDescription: "invalid JSON response from platform API server",
	}
)
