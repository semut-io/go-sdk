package workergroup

import (
	"go.semut.io/sdk/go-sdk/pkg/appmanager/worker"
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// WorkerGroup is details of set of workers
type WorkerGroup struct {
	// ID of the deployment under which worker group exists
	DeploymentID string `json:"deployment_id,omitempty"`
	// ID of the worker group
	WorkerGroupID string `json:"worker_group_id"`
	// Details of worker group
	LaunchSpec
	// DedicatedIPAddress indicates if ip address attached is dedicated or shared
	DedicatedExternalIPAddress bool `json:"dedicated_external_ip_address,omitempty"`
	// ExternalIPAddress allocated to the worker group
	ExternalIPAddress string `json:"external_ip_address,omitempty"`
	// InternalIPAddress allocated to the worker group
	InternalIPAddress string `json:"internal_ip_address,omitempty"`
	// ExternalDNSName allocated to the worker group
	ExternalDNSName string `json:"external_dns_name,omitempty"`
	// InternalDNSName allocated to the worker group
	InternalDNSName string `json:"internal_dns_name,omitempty"`
	// Health indicates if the workergroup is healthy or unhealthy.
	common.Health `json:"health,omitempty"`
	// State indicates state of worker group
	common.Status `json:"status,omitempty"`
	// All workers in the worker group
	Workers []worker.Worker `json:"worker_instances,omitempty"`
}

// WorkerGroupRequestFields common request fields in worker group
type WorkerGroupRequestFields struct {
	// ID of the deploymnent
	DeploymentID string `json:"deployment_id,omitempty"`
	// ID of the worker group
	WorkerGroupID string `json:"worker_group_id,omitempty"`
}

// WorkerLaunchSpec defines all fields needed for launching new worker for the worker group
type LaunchSpec struct {
	worker.WorkerLaunchSpec
	// UpdateStrategy is the strategy to be used to update the Workers when the WorkerGroup is updated.
	UpdateStrategy common.UpdateStrategy `json:"update_strategy"`
	// NumWorkers is the count of Workers belonging to the WorkerGroup.
	NumWorkers int `json:"num_workers"`
}

// LaunchRequest is used request new worker group
type LaunchRequest struct {
	common.AsyncRequest
	// ID of the Deployment
	DeploymentID string `json:"deployment_id"`
	// Specification of WorkerGroup
	LaunchSpec
}

// LaunchResponse is response to launch request
type LaunchResponse struct {
	common.AsyncResponse
	// Worker group ID
	WorkerGroupID string `json:"worker_group_id"`
}

// LaunchCallback is callback response to launch request
type LaunchCallback struct {
	common.AsyncResponse
	// List of worker groups that were launched
	WorkerGroups []WorkerGroup `json:"worker_groups,omitempty"`
}

// DescribeRequest is used to describe multiple worker groups
type DescribeRequest struct {
	DeploymentID string `json:"deployment_id"`
	// WorkerGroupIDs is the list of unique UUIDs to be described.
	WorkerGroupIDs []string `json:"worker_group_ids,omitempty"`
}

// DescribeResponse is response to describe request
type DescribeResponse struct {
	common.APIResponse
	// List of worker group with respective information
	WorkerGroups []WorkerGroup `json:"worker_groups,omitempty"`
}

// TerminateRequest is request used to terminate a workergroup
type TerminateRequest struct {
	common.AsyncRequest
	// Request fields with IDs
	WorkerGroupRequestFields
}

// TerminateResponse is response to terminate request
type TerminateResponse struct {
	common.AsyncResponse
}

// TerminateCallback is callback response to terminate request
type TerminateCallback struct {
	common.AsyncResponse
	// WorkerID is returned if deletion is successful
	WorkerGroupID string `json:"worker_group_id,omitempty"`
}

// StopRequest is used to request stop of worker group
type StopRequest struct {
	common.AsyncRequest
	// Request fields with IDs
	WorkerGroupRequestFields
}

// StopResponse is response to stop request
type StopResponse struct {
	common.AsyncResponse
}

// StopCallback is callback response to stop request
type StopCallback struct {
	common.AsyncResponse
	// Worker group details
	WorkerGroup
}

// StartRequest is used to request start of worker group
type StartRequest struct {
	common.AsyncRequest
	// Request fields with IDs
	WorkerGroupRequestFields
}

// StartResponse is response to start request
type StartResponse struct {
	common.AsyncResponse
}

// StartCallback is callback response to start request
type StartCallback struct {
	common.AsyncResponse
	// Worker group details
	WorkerGroup
}

// MarkHealthyRequest is used to mark worker group as healthy
type MarkHealthyRequest struct {
	common.AsyncRequest
	// Request fields with IDs
	WorkerGroupRequestFields
}

// MarkHealthyResponse is response to mark healthy request
type MarkHealthyResponse struct {
	common.AsyncResponse
}

// MarkHealthyCallback is callback response to mark healthy request
type MarkHealthyCallback struct {
	common.AsyncResponse
	// Detailed information about the worker group
	WorkerGroup
}

// MarkUnhealthyRequest is used to mark worker group as unhealthy
type MarkUnhealthyRequest struct {
	common.AsyncRequest
	// Request fields with IDs
	WorkerGroupRequestFields
}

// MarkUnhealthyResponse is response to mark unhealthy request
type MarkUnhealthyResponse struct {
	common.AsyncResponse
}

// MarkUnhealthyCallback is callback response to mark unhealthy request
type MarkUnhealthyCallback struct {
	common.AsyncResponse
	// Details of worker group
	WorkerGroup
}

// HealthStatusRequest is used to request health status of worker group
type HealthStatusRequest struct {
	WorkerGroupRequestFields
}

// HealthStatusResponse is response to health status request
type HealthStatusResponse struct {
	common.APIResponse
	// Health status of worker group
	common.Health
}

// UpdateRequest is used to update worker group fields
type UpdateRequest struct {
	common.AsyncRequest
	// ID of deployment
	DeploymentID string `json:"deployment_id,omitempty"`
	// ID of worker group
	WorkerGroupID string `json:"worker_group_id"`
	// Worker group spec
	LaunchSpec
}

// UpdateResponse is response to update request
type UpdateResponse struct {
	common.AsyncResponse
}

// UpdateCallback is callback response to update request
type UpdateCallback struct {
	common.AsyncResponse
	// Details of worker group
	WorkerGroup
}

// ChangeUpdateStrategyRequest is used to request change of worker group update strategy
type ChangeUpdateStrategyRequest struct {
	common.AsyncRequest
	// ID of deployment
	DeploymentID string
	// ID of worker group
	WorkerGroupID string
	// New update strategy
	UpdateStrategy string
}

// ChangeUpdateStrategyResponse is the response to the change update strategy request
type ChangeUpdateStrategyResponse struct {
	common.AsyncResponse
}

// ChangeUpdateStrategyCallback is the callback response to the change update strategy request
type ChangeUpdateStrategyCallback struct {
	common.AsyncResponse
	// Worker group details
	WorkerGroup
}

// UpdateResourcesRequest is used to request change of worker group resources
type UpdateResourcesRequest struct {
	common.AsyncRequest
	// ID of worker group
	WorkerGroupID string
	// Resource request
	ResourceRequestWorker common.ResourceRequestRange
}

// UpdateResourcesResponse is the response to the change update resources request
type UpdateResourcesResponse struct {
	common.AsyncResponse
}

// UpdateResourcesCallback is the callback response to the update resources request
type UpdateResourcesCallback struct {
	common.AsyncResponse
	// Requested fields with IDs
	WorkerGroupRequestFields
}

// ScaleRequest is used to request change in number of workers in worek group
type ScaleRequest struct {
	common.AsyncRequest
	// Request fields with IDs
	WorkerGroupRequestFields
	// Number of workers have in the worker group
	NumWorkers int `json:"num_workers"`
}

// ScaleResponse is the response to the scale request
type ScaleResponse struct {
	common.AsyncResponse
}

// ScaleCallback is the callback response to the scale request
type ScaleCallback struct {
	common.AsyncResponse
	// Worker group details
	WorkerGroup
}

// Launch new worker group
func (launchRequest *LaunchRequest) Launch() (workerGroupID, requestID string, apiErr *common.Error) {

	launchResponse := LaunchResponse{}
	err := common.Execute("WorkerGroupLaunch", launchRequest, &launchResponse)

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

	return launchResponse.WorkerGroupID, launchResponse.RequestToken, nil
}

// Describe a specific worker group or all worker groups in a deployment
func (describeRequest *DescribeRequest) Describe() (workerGroups []WorkerGroup, apiErr *common.Error) {

	describeResponse := DescribeResponse{}
	err := common.Execute("WorkerGroupDescribe", describeRequest, &describeResponse)

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

	return describeResponse.WorkerGroups, nil

}

// Update one or more properties of worker group
func (updateRequest *UpdateRequest) Update() (requestID string, apiErr *common.Error) {

	updateResponse := UpdateResponse{}
	err := common.Execute("WorkerGroupUpdate", updateRequest, &updateResponse)

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

// Terminate is used to delete a worker group
func (terminateRequest *TerminateRequest) Terminate() (requestID string, apiErr *common.Error) {

	terminateResponse := TerminateResponse{}
	err := common.Execute("WorkerGroupTerminate", terminateRequest, &terminateResponse)

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

// Stop a worker group by removing all workers but retaining all other resources like allocated ip addresses
func (stopRequest *StopRequest) Stop() (requestID string, apiErr *common.Error) {

	stopResponse := StopResponse{}
	err := common.Execute("WorkerGroupStop", stopRequest, &stopResponse)

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

// Start an already existing worker group
func (startRequest *StartRequest) Start() (requestID string, apiErr *common.Error) {

	startResponse := StartResponse{}
	err := common.Execute("WorkerGroupStart", startRequest, &startResponse)

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
	err := common.Execute("WorkerGroupMarkHealthy", markHealthyRequest, &markHealthyResponse)

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

// MarkUnhealthy ensures worker group is marked as unhealthy
func (markUnhealthyRequest *MarkUnhealthyRequest) MarkUnhealthy() (requestID string, apiErr *common.Error) {

	markUnhealthyResponse := MarkUnhealthyResponse{}
	err := common.Execute("WorkerGroupMarkUnhealthy", markUnhealthyRequest, &markUnhealthyResponse)

	if err != nil {
		return "", err
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

// HealthStatus returns health of the status of the worker group
func (healthStatusRequest *HealthStatusRequest) HealthStatus() (healthStatus common.Health, apiErr *common.Error) {

	healthStatusResponse := HealthStatusResponse{}
	err := common.Execute("WorkerGroupHealthStatus", healthStatusRequest, &healthStatusResponse)

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

// ChangeUpdateStrategy changes update strategy of the worker group
func (changeUpdateStrategyRequest *ChangeUpdateStrategyRequest) ChangeUpdateStrategy() (requestID string, apiErr *common.Error) {

	changeUpdateStrategyResponse := ChangeUpdateStrategyResponse{}
	err := common.Execute("WorkerGroupChangeStrategy", changeUpdateStrategyRequest, &changeUpdateStrategyResponse)

	if err != nil {
		return "", err
	}

	if changeUpdateStrategyResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        changeUpdateStrategyResponse.StatusCode,
			ErrorDescription: changeUpdateStrategyResponse.Description,
		}

		return "", apiErr
	}

	return changeUpdateStrategyResponse.RequestToken, nil
}

// UpdateResourceLimits changes or adds (if not specified earlier) resource limits
func (updateResourcesRequest *UpdateResourcesRequest) UpdateResourceLimits() (requestID string, apiErr *common.Error) {

	updateResourcesResponse := UpdateResourcesResponse{}
	err := common.Execute("WorkerGroupUpdateResourceLimits", updateResourcesRequest, &updateResourcesResponse)

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

// Scale changes or adds or removes workers to worker group
func (scaleRequest *ScaleRequest) Scale() (requestID string, apiErr *common.Error) {

	scaleResponse := ScaleResponse{}
	err := common.Execute("WorkerGroupScale", scaleRequest, &scaleResponse)

	if err != nil {
		return "", err
	}

	if scaleResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        scaleResponse.StatusCode,
			ErrorDescription: scaleResponse.Description,
		}

		return "", apiErr
	}

	return scaleResponse.RequestToken, nil
}
