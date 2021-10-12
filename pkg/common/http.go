package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"

	retry "github.com/hashicorp/go-retryablehttp"
)

// Execute calls the platform API and stores the result in the response. This method uses an
// exponential back-off based retry loop and ensure that response passed is a pointer to the struct
// containing the response information.
func Execute(endpoint string, request interface{}, response interface{}) *Error {

	supportedVersions := GetSupportedVersions()

	// set version as v1
	currentVersion := supportedVersions[0]

	success, err := SetVersion(currentVersion)
	if !success {
		return &Error{ErrorCode: "100", ErrorDescription: err.Error()}
	}

	endpointPath, err := GetEndpoint(endpoint)
	if err != nil {
		return &Error{ErrorCode: "100", ErrorDescription: err.Error()}
	}

	endpointURL := getEndpointURL(endpointPath)
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return &ErrInvalidInput
	}

	jsonResponse, statusCode, err := makeCall(endpointURL, jsonRequest)

	if err != nil {
		return &Error{
			ErrorCode:        strconv.Itoa(statusCode),
			ErrorDescription: err.Error(),
		}
	}

	err = json.Unmarshal(jsonResponse, response)
	if err != nil {
		return &ErrInvalidResponseFromAPI
	}

	return nil
}

// getEndpointURL is used to append the platform API server hostname with the endpoint paths
func getEndpointURL(endpointPath string) string {
	return GetPlatformURL() + endpointPath
}

// makeCall implements the interface UpstreamRequester and behaves as the final tip which touches the platform
func makeCall(URI string, body []byte) (response []byte, statusCode int, err error) {

	retryClient := retry.NewClient()
	retryClient.RetryMax = 5

	resp, err := retryClient.Post(URI, "application/json", body)

	// 500 Internal Server Errors will be auto-retried by go-retryablehttp
	// at this point if any!

	if err != nil {
		return nil, 502, errors.New("platform API server unavailable")
	}

	resposeBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, 503, errors.New("error reading response from platform API server")
	}

	return resposeBody, resp.StatusCode, nil
}
