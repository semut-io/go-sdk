package applications

import (
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// InvokeRequest invokes another application
type InvokeRequest struct {
	common.AsyncRequest
	// ID of the application being invoked. Note:This is id of the public or proviate application. This id is unique across all regions of the platform
	ApplicationID string `json:"application_id"`
	// Endpoint of the application to be invoked
	Path string `json:"path"`
	// Request payload
	Payload interface{} `json:"payload"`
}

// InvokeResponse response from the platform and the application. If application was not invoked response would be empty
type InvokeResponse struct {
	common.APIResponse
}

// InvokeCallback response received via callback from the invoked application manager
type InvokeCallback struct {
	common.AsyncResponse
	// Response is the response given by the invoked application manager.
	Response []byte `json:"response"`
}

// Invoke endpoint of another application
func (invokeRequest *InvokeRequest) Invoke() (apiErr *common.Error) {
	invokeResponse := InvokeResponse{}
	err := common.Execute("ApplicationInvoke", invokeRequest, &invokeResponse)

	if err != nil {
		return &common.ErrInvalidResponseFromAPI
	}

	if invokeResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        invokeResponse.StatusCode,
			ErrorDescription: invokeResponse.Description,
		}

		return apiErr
	}

	return nil
}
