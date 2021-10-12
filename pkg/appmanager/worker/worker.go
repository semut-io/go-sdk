package worker

import (
	"go.semut.io/sdk/go-sdk/pkg/appmanager/volumes"
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// Worker represent a deployed worker
type Worker struct {
	// ID of the deployment under which worker group exists
	DeploymentID string `json:"deployment_id,omitempty"`
	// ID of the worker group
	WorkerGroupID string `json:"worker_group_id,omitempty"`
	// ID of the worker
	WorkerID string `json:"worker_id"`
	// Details of worker
	WorkerLaunchSpec
	// ExternalIPAddres allocated to the worker
	ExternalIPAddress string `json:"external_ip_address,omitempty"`
	// InternalIPAddres allocated to the worker
	InternalIPAddress string `json:"internal_ip_address,omitempty"`
	// ExternalDNSName allocated to the worker
	ExternalDNSName string `json:"external_dns_name,omitempty"`
	// InternalDNSName allocated to the worker
	InternalDNSName string `json:"internal_dns_name,omitempty"`
	// Status indicates if the worker is stopped or running. This is not the health status of the worker
	Status string `json:"status,omitempty"`
	// Health indicates if the worker is healthy or unhealthy.
	common.Health `json:"health,omitempty"`
}

// WorkerRequestFields common request fields in worker group
type WorkerRequestFields struct {
	// ID of the deployment under which worker group exists
	DeploymentID string `json:"deployment_id,omitempty"`
	// Worker group ID
	WorkerGroupID string `json:"worker_group_id,omitempty"`
	// ID of the worker
	WorkerID string `json:"worker_id,omitempty"`
}

// LaunchRequest is to be used for launching new workers
type LaunchRequest struct {
	common.AsyncRequest
	// ID of the deployment
	DeploymentID string `json:"deployment_id"`
	// ID of worker group
	WorkerGroupID string `json:"worker_group_id,omitempty"`
	// Spec for new worker creation
	WorkerLaunchSpec
}

// WorkerSpec is the worker spec received from the worker
type WorkerLaunchSpec struct {
	// Name of the Worker.
	Name string `json:"name"`
	// NumWorkers is the number of Workers to be launched.
	NumWorkers int `json:"num_workers"`
	// Image is the container image to be used for the Worker. It is an optional field and the value is inherited
	// from the parent WorkerGroup if none is specified.
	Image string `json:"image,omitempty"`
	// Volumes is the list of volumes that are to be attached to the Worker. It is an optional field and the
	// value is inherited from the parent WorkerGroup if none is specified.
	Volumes []volumes.Volume `json:"volumes,omitempty"`
	// Ports is the list of exposed container ports on the Worker. It is an optional field and the value is
	// inherited from the parent WorkerGroup if none is specified.
	Ports []common.Port `json:"ports,omitempty"`
	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment.
	Entrypoint []string `json:"entrypoint,omitempty"`
	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment.
	// If a variable cannot be resolved, the reference in the input string will be unchanged.
	Cmd []string `json:"cmd,omitempty"`
	// MetricsEndpoint is the embedded struct that contains the details of the metrics endpoint exposed on the workers
	common.MetricsEndpoint
	// EnvVariables is the list of environment variables set in the Worker. It is an optional field and the
	// value is inherited from the parent WorkerGroup if none is specified.
	EnvVariables map[string]string `json:"env_variables,omitempty"`
	// Tags are the set of labels that are added to the Worker. It is an optional field and the value is
	// inherited from the parent WorkerGroup if none is specified.
	Tags map[string]string `json:"tags,omitempty"`
	// Logs is list of paths available for the Worker.
	Logs []common.Log `json:"logs"`
	// ResourceRequest is the resource requests and ranges that are to be applied to the Worker. It is an
	// optional field and the value is inherited from the parent WorkerGroup if none is specified.
	ResourceRequest common.ResourceRequestRange `json:"resource_request,omitempty"`
}

// LaunchCallback is the callback response for the launch request
type LaunchCallback struct {
	common.AsyncResponse
	// Set of workers to launch
	Workers []Worker `json:"workers,omitempty"`
}

// LaunchResponse is the response to the launch request
type LaunchResponse struct {
	common.AsyncResponse
	// IDs of launched workers
	WorkerIDs []string `json:"worker_ids,omitempty"`
}

// DescribeRequest is used to describe workers
type DescribeRequest struct {
	// ID of the deployment
	DeploymentID string `json:"deployment_id"`
	// ID of worker group
	WorkerGroupID string `json:"worker_group_id,omitempty"`
	// IDs of workers
	WorkerIDs []string `json:"worker_ids"`
}

// DescribeResponse is response to the describe request
type DescribeResponse struct {
	common.APIResponse
	Workers []Worker `json:"workers,omitempty"`
}

// TerminateRequest is used to request worker deletion
type TerminateRequest struct {
	WorkerRequestFields
	common.AsyncRequest
}

// TerminateResponse is the response to worker terminate request
type TerminateResponse struct {
	common.AsyncResponse
}

// TerminateCallback is the callback response for the terminate request
type TerminateCallback struct {
	common.AsyncResponse
	// WorkerID is returned if deletion is successful
	WorkerID string `json:"worker_id"`
}

// StopRequest is used to stop a worker
type StopRequest struct {
	WorkerRequestFields
	common.AsyncRequest
}

// StopResponse is the response to the stop request
type StopResponse struct {
	common.AsyncResponse
}

// StopCallback is the callback response for the stop request
type StopCallback struct {
	common.AsyncResponse
	// Details of the worker stopped
	Worker
}

// StartRequest is used to start a worker
type StartRequest struct {
	common.AsyncRequest
	// Request fields
	WorkerRequestFields
}

// StartResponse is the response to the start request
type StartResponse struct {
	common.AsyncResponse
}

// StartCallback is the callback response for the start request
type StartCallback struct {
	common.AsyncResponse
	// Details of the worker started
	Worker
}

// UpdateRequest is used to update a worker
type UpdateRequest struct {
	common.AsyncRequest
	// Request field with IDs
	WorkerRequestFields
	// Image is the container image to be used for the Worker. It is an optional field and the value is inherited
	// from the parent WorkerGroup if none is specified.
	Image string `json:"image,omitempty"`
	// Volumes is the list of volumes that are to be attached to the Worker. It is an optional field and the
	// value is inherited from the parent WorkerGroup if none is specified.
	Volumes []volumes.Volume `json:"volumes,omitempty"`
	// Ports is the list of exposed container ports on the Worker. It is an optional field and the value is
	// inherited from the parent WorkerGroup if none is specified.
	Ports []common.Port `json:"ports,omitempty"`
	// EnvVariables is the list of environment variables set in the Worker. It is an optional field and the
	// value is inherited from the parent WorkerGroup if none is specified.
	EnvVariables map[string]string `json:"env_variables,omitempty"`
	// Tags are the set of labels that are added to the Worker. It is an optional field and the value is
	// inherited from the parent WorkerGroup if none is specified.
	Tags map[string]string `json:"tags,omitempty"`
	// ResourceRequest is the resource requests and ranges that are to be applied to the Worker. It is an
	// optional field and the value is inherited from the parent WorkerGroup if none is specified.
	ResourceRequest common.ResourceRequestRange `json:"resource_request,omitempty"`
}

// UpdateResponse is response to the update request
type UpdateResponse struct {
	common.AsyncResponse
}

// UpdateCallback is callback response for the update request
type UpdateCallback struct {
	common.AsyncResponse
	Worker
}

// MarkHealthyRequest is used to mark a worker as healthy
type MarkHealthyRequest struct {
	common.AsyncRequest
	// Request fields with IDs
	WorkerRequestFields
}

// MarkHealthyResponse is the response to the mark healthy request
type MarkHealthyResponse struct {
	common.AsyncResponse
}

// MarkHealthyCallback is the callback response for the mark healthy request
type MarkHealthyCallback struct {
	common.AsyncResponse
	// MarkHealthySuccessful bool
	Worker
}

// MarkUnhealthyRequest is used to mark a worker as unhealthy
type MarkUnhealthyRequest struct {
	WorkerRequestFields
	common.AsyncRequest
}

// MarkUnhealthyResponse is the response to the mark unhealthy request
type MarkUnhealthyResponse struct {
	common.AsyncResponse
}

// MarkUnhealthyCallback is the callback response for the mark unhealthy request
type MarkUnhealthyCallback struct {
	common.AsyncResponse
	Worker
}

// HealthStatusRequest is the request to get the health status of a worker
type HealthStatusRequest struct {
	WorkerRequestFields
}

// HealthStatusResponse is the response to the health status request
type HealthStatusResponse struct {
	common.APIResponse
	// Health of the worker
	common.Health
}

// UpdateResourcesRequest is the request to update resources of a worker
type UpdateResourcesRequest struct {
	common.AsyncRequest
	// Request fields with IDs
	WorkerRequestFields
	// Request for resources
	ResourceRequestWorker common.ResourceRequestRange
}

// UpdateResourcesResponse is response to the update resources request
type UpdateResourcesResponse struct {
	common.AsyncResponse
}

// UpdateResourcesCallback is callback response for the update resources request
type UpdateResourcesCallback struct {
	common.AsyncResponse
	// Worker details
	Worker
}

// Describe a specific worker group or all worker groups in a deployment
func (describeRequest *DescribeRequest) Describe() (workers []Worker, apiErr *common.Error) {

	describeResponse := DescribeResponse{}
	err := common.Execute("WorkerDescribe", describeRequest, &describeResponse)

	if err != nil {
		return nil, err
	}

	if describeResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        describeResponse.StatusCode,
			ErrorDescription: describeResponse.Description,
		}

		return nil, apiErr
	}

	return describeResponse.Workers, nil

}

// Launch new worker group
func (launchRequest *LaunchRequest) Launch() (workerIDs []string, requestID string, apiErr *common.Error) {

	launchResponse := LaunchResponse{}
	err := common.Execute("WorkerLaunch", launchRequest, &launchResponse)

	if err != nil {
		return []string{}, "", &common.ErrInvalidResponseFromAPI
	}

	if launchResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        launchResponse.StatusCode,
			ErrorDescription: launchResponse.Description,
		}

		return []string{}, "", apiErr
	}

	return launchResponse.WorkerIDs, launchResponse.RequestToken, nil
}

// Terminate is used to delete a worker
func (terminateRequest *TerminateRequest) Terminate() (requestID string, apiErr *common.Error) {

	terminateResponse := TerminateResponse{}
	err := common.Execute("WorkerTerminate", terminateRequest, &terminateResponse)

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

// Stop worker group by removing all workers but retaining all other resources like allocated ip addresses
func (stopRequest *StopRequest) Stop() (requestID string, apiErr *common.Error) {

	stopResponse := StopResponse{}
	err := common.Execute("WorkerStop", stopRequest, &stopResponse)

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

// Start worker group by launching
func (startRequest *StartRequest) Start() (requestID string, apiErr *common.Error) {

	startResponse := StartResponse{}
	err := common.Execute("WorkerStart", startRequest, &startResponse)

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

// MarkHealthy ensures worker group is marked as healthy
func (markHealthyRequest *MarkHealthyRequest) MarkHealthy() (requestID string, apiErr *common.Error) {

	markHealthyResponse := MarkHealthyResponse{}
	err := common.Execute("WorkerMarkHealthy", markHealthyRequest, &markHealthyResponse)

	if err != nil {
		return "", err
	}

	if markHealthyResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        markHealthyResponse.StatusCode,
			ErrorDescription: markHealthyResponse.Description,
		}

		return "", apiErr
	}

	return markHealthyResponse.RequestToken, nil
}

// MarkUnhealthy ensures worker group is marked as healthy
func (markUnhealthyRequest *MarkUnhealthyRequest) MarkUnhealthy() (requestID string, apiErr *common.Error) {

	markUnhealthyResponse := MarkUnhealthyResponse{}
	err := common.Execute("WorkerMarkUnhealthy", markUnhealthyRequest, &markUnhealthyResponse)

	if err != nil {
		return "", &common.ErrInvalidResponseFromAPI
	}

	if markUnhealthyResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        markUnhealthyResponse.StatusCode,
			ErrorDescription: markUnhealthyResponse.Description,
		}

		return "", apiErr
	}

	return markUnhealthyResponse.RequestToken, nil
}

// HealthStatus return health of the status
func (healthStatusRequest *HealthStatusRequest) HealthStatus() (healthStatus common.Health, apiErr *common.Error) {

	healthStatusResponse := HealthStatusResponse{}
	err := common.Execute("WorkerHealthStatus", healthStatusRequest, &healthStatusResponse)

	if err != nil {
		return "", err
	}

	if healthStatusResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        healthStatusResponse.StatusCode,
			ErrorDescription: healthStatusResponse.Description,
		}

		return "", apiErr
	}

	return healthStatusResponse.Health, nil
}

// UpdateResourceLimits changes or adds (if not specified earlier) resource limits
func (updateResourcesRequest *UpdateResourcesRequest) UpdateResourceLimits() (requestID string, apiErr *common.Error) {

	updateResourcesResponse := UpdateResourcesResponse{}
	err := common.Execute("WorkerUpdateResourceLimits", updateResourcesRequest, &updateResourcesResponse)

	if err != nil {
		return "", err
	}

	if updateResourcesResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        updateResourcesResponse.StatusCode,
			ErrorDescription: updateResourcesResponse.Description,
		}

		return "", apiErr
	}

	return updateResourcesResponse.RequestToken, nil
}

// Update worker specification
func (updateRequest *UpdateRequest) Update() (requestID string, apiErr *common.Error) {

	updateResponse := UpdateResponse{}
	err := common.Execute("WorkerUpdate", updateRequest, &updateResponse)

	if err != nil {
		return "", err
	}

	if updateResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        updateResponse.StatusCode,
			ErrorDescription: updateResponse.Description,
		}

		return "", apiErr
	}

	return updateResponse.RequestToken, nil
}
