package common

import "github.com/google/uuid"

// GetPlatformURL Returns the URL of the platform API server
// that can be reachable from the Application Manager
func GetPlatformURL() string {
	return "http://localhost:53377"
}

// GenerateRequestToken generates a request token to be used for asynchronous calls,
// it generates a random UUID string that can be passed as request token. This is
// primarily used to track async endpoints that use callbacks.
func GenerateRequestToken() string {
	return uuid.New().String()
}
