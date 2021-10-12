package snapshots

import (
	"time"

	"go.semut.io/sdk/go-sdk/pkg/common"
)

// Snapshot meta
type Snapshot struct {
	// ID of snapshot
	SnapshotID string `json:"snapshot_id"`
	// Name of snapshot
	Name string `json:"name,omitempty"`
	// Time when snapshot was created
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// InitiateRequest is used to initiate a snaphot
type InitiateRequest struct {
	// ID of the deployment
	DeploymentID string `json:"deployment_id"`
	// Volume ID of which snapshot has to be taken
	VolumeID string `json:"volume_id"`
	// Name of the snapshot
	Name string `json:"name"`
	common.AsyncRequest
}

// InitiateResponse response to the initiate request
type InitiateResponse struct {
	// ID of the snaphot
	SnapshotID string `json:"snapshot_id"`
	common.AsyncResponse
}

// InitiateCallback callback response from platform
type InitiateCallback struct {
	common.AsyncResponse
	// Callback response containing snapshot info
	Snapshot
}

// DescribeRequest get information on a snapshot or all snapshots in a deployment
type DescribeRequest struct {
	DeploymentID string `json:"deployment_id"`
	// SnapshotIDs that are to be described.
	SnapshotIDs []string `json:"snapshot_ids,omitempty"`
}

// DescribeResponse information of specified snapshot or all snapshots in the
type DescribeResponse struct {
	common.APIResponse
	Snapshots []Snapshot `json:"snapshots,omitempty"`
}

// DeleteRequest snapshot to be deleted
type DeleteRequest struct {
	DeploymentID string `json:"deployment_id"`
	SnapshotID   string `json:"snapshot_id"`
}

// DeleteResponse snapshot to be deleted
type DeleteResponse struct {
	common.APIResponse
}

// Initiate initiates process for taking snapshot of the specified volume
func (initiateRequest *InitiateRequest) Initiate() (snapshotID, requestID string, apiErr *common.Error) {

	initiateResponse := InitiateResponse{}
	err := common.Execute("SnapshotsTake", initiateRequest, &initiateResponse)

	if err != nil {
		return "", "", &common.ErrInvalidResponseFromAPI
	}

	if initiateResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        initiateResponse.StatusCode,
			ErrorDescription: initiateResponse.Description,
		}

		return "", "", apiErr
	}

	return initiateResponse.SnapshotID, initiateResponse.RequestToken, nil
}

// Describe a snapshot or all snapshots in a deployment
func (describeRequest *DescribeRequest) Describe() (snapshots []Snapshot, apiErr *common.Error) {

	describeResponse := DescribeResponse{}
	err := common.Execute("SnapshotsDescribe", describeRequest, &describeResponse)

	if err != nil {
		return []Snapshot{}, err
	}

	if describeResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        describeResponse.StatusCode,
			ErrorDescription: describeResponse.Description,
		}

		return []Snapshot{}, apiErr
	}

	return describeResponse.Snapshots, nil
}

// Delete a snapshot
func (deleteRequest *DeleteRequest) Delete() (success bool, apiErr *common.Error) {

	deleteResponse := DeleteResponse{}
	err := common.Execute("SnapshotsDelete", deleteRequest, &deleteResponse)

	if err != nil {
		return false, err
	}

	if deleteResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        deleteResponse.StatusCode,
			ErrorDescription: deleteResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}
