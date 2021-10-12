package deployment

import (
	"go.semut.io/sdk/go-sdk/pkg/appmanager/workergroup"
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// Deployment contains complete information about a deployment
type Deployment struct {
	// ID of the deployment
	DeploymentID string `json:"deployment_id"`
	// Name of the deployment
	Name string `json:"name,omitempty"`
	// Worker groups associated with the deployment
	WorkerGroups []workergroup.WorkerGroup `json:"worker_groups,omitempty"`
	// Current status of the deployment
	Status common.Status `json:"status,omitempty"`
}

// DescribeRequest to describe a single deployment
type DescribeRequest struct {
	// Unique ID
	DeploymentID string `json:"deployment_id,omitempty"`
}

// DescribeResponse is the response to the describe request
type DescribeResponse struct {
	common.APIResponse
	// List of deployments
	Deployment Deployment `json:"deployment,omitempty"`
}

// LaunchRequest is used to request a new deployment
type LaunchRequest struct {
	// Name of the deployment to be created
	Name string `json:"name,omitempty"`
	// List of worker group spceifications for the new deployment
	WorkerGroups []workergroup.LaunchSpec `json:"worker_groups"`
	common.AsyncRequest
}

// LaunchResponse response to launch deployment request
type LaunchResponse struct {
	common.AsyncResponse
	// Unique ID of the new deployment
	DeploymentID string `json:"deployment_id"`
}

// LaunchCallback is the response callback from launch rqeust
type LaunchCallback struct {
	common.AsyncResponse
	// Deployment details
	Deployment
}

// TerminateRequest is used to request deletion of a deployment
type TerminateRequest struct {
	common.AsyncRequest
	// ID of the deployment to be deleted
	DeploymentID string `json:"deployment_id"`
}

// TerminateResponse is the response to delete request
type TerminateResponse struct {
	common.AsyncResponse
}

// TerminateCallback is the response callback from delete request
type TerminateCallback struct {
	common.AsyncResponse
	// DeploymentID is returned if deletion is successful
	DeploymentID string `json:"deployment_id,omitempty"`
}

// StopRequest is used to request stopping of a running deployment
type StopRequest struct {
	common.AsyncRequest
	// ID of the deployment to be stopped
	DeploymentID string `json:"deployment_id"`
}

// StopResponse is response to stop request
type StopResponse struct {
	common.AsyncResponse
	// ID of deployment
	DeploymentID string `json:"deployment_id,omitempty"`
}

// StopCallback is the response callback from stop request
type StopCallback struct {
	common.AsyncResponse
	// Deployment details
	Deployment
}

// StartRequest is used to start the deployment
type StartRequest struct {
	common.AsyncRequest
	// ID of deployment to be started
	DeploymentID string `json:"deployment_id"`
}

// StartResponse is response to start request
type StartResponse struct {
	common.AsyncResponse
	// ID of deployment
	DeploymentID string `json:"deployment_id"`
}

// StartCallback is the response callback from start request
type StartCallback struct {
	common.AsyncResponse
	// Deployment details
	Deployment
}

// Describe a deployment
func (describeRequest *DescribeRequest) Describe() (deployment Deployment, apiErr *common.Error) {

	describeResponse := DescribeResponse{}
	err := common.Execute("DeploymentDescribe", describeRequest, &describeResponse)

	if err != nil {
		return deployment, err
	}

	if describeResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        describeResponse.StatusCode,
			ErrorDescription: describeResponse.Description,
		}

		return deployment, apiErr
	}

	return describeResponse.Deployment, nil
}

// Launch a new deployment, returns id of the new deployment and request ID
func (launchRequest *LaunchRequest) Launch() (deploymentID, requestID string, apiErr *common.Error) {

	launchResponse := LaunchResponse{}
	err := common.Execute("DeploymentLaunch", launchRequest, &launchResponse)

	if err != nil {
		return "", "", err
	}

	if launchResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        launchResponse.StatusCode,
			ErrorDescription: launchResponse.Description,
		}

		return "", "", apiErr
	}

	return launchResponse.DeploymentID, launchResponse.RequestToken, nil
}

// Terminate an existing deployment, returns a request ID
func (terminateRequest *TerminateRequest) Terminate() (requestID string, apiErr *common.Error) {

	terminateResponse := TerminateResponse{}
	err := common.Execute("DeploymentTerminate", terminateRequest, &terminateResponse)

	if err != nil {
		return "", err
	}

	if terminateResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        terminateResponse.StatusCode,
			ErrorDescription: terminateResponse.Description,
		}

		return "", apiErr
	}

	return terminateResponse.RequestToken, nil
}

// Start is used to start an existing deployment, request ID is returned
func (startRequest *StartRequest) Start() (requestID string, apiErr *common.Error) {

	startResponse := StartResponse{}
	err := common.Execute("DeploymentStart", startRequest, &startResponse)

	if err != nil {
		return "", err
	}

	if startResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        startResponse.StatusCode,
			ErrorDescription: startResponse.Description,
		}

		return "", apiErr
	}

	return startResponse.RequestToken, nil
}

// Stop is used to stop an existing deployment, request ID is returned
func (stopRequest *StopRequest) Stop() (requestID string, apiErr *common.Error) {
	stopResponse := StopResponse{}
	err := common.Execute("DeploymentStop", stopRequest, &stopResponse)

	if err != nil {
		return "", err
	}

	if stopResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        stopResponse.StatusCode,
			ErrorDescription: stopResponse.Description,
		}

		return "", apiErr
	}

	return stopResponse.RequestToken, nil
}
