package volumes

import (
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// VolumeSpec describes volume
type VolumeSpec struct {
	// Name of the volume as it exists on the cluster.
	Name string `json:"name,omitempty"`
	// Type whether to use Block or FS type volume
	Type common.VolumeType `json:"type,omitempty"`
	// Dedicated volume otherwise
	// shared  volume, when dedicated is false
	// PS: Block volumes can only be dedicated
	Dedicated bool `json:"dedicated,omitempty"`
	// MountPath is the absolute path at which the volume will be mounted inside
	// the container.
	MountPath string `json:"mount_path,omitempty"`
}

// Volume an instance of volume
type Volume struct {
	// Either ID of volume can be used
	VolumeID string `json:"volume_id,omitempty"`
	// Or spec to create a new volume
	VolumeSpec
}

// DescribeRequest get volume meta by id
type DescribeRequest struct {
	// ID of deplyment
	DeploymentID string `json:"deployment_id"`
	// Unique ID of the volume
	VolumeID string `json:"volume_id"`
}

// DescribeResponse response of describe request
type DescribeResponse struct {
	common.APIResponse
	// Details of the volume
	Volume
}

// CreateRequest spec for creating new volume
type CreateRequest struct {
	// ID of the deployment under which volume is being created
	DeploymentID string `json:"deployment_id"`
	// Create a volume from existing snapshot
	SnapshotID string `json:"snapshot_id,omitempty"`
	// Volume specification
	VolumeSpec `json:"volume_spec,omitempty"`
	common.AsyncRequest
}

// CreateResponse spec for creating new volume
type CreateResponse struct {
	common.AsyncResponse
	// ID of volume
	VolumeID string `json:"volume_id,omitempty"`
}

// CreateCallback is callback response of the create volume request
type CreateCallback struct {
	common.AsyncResponse
	// Details of volume
	Volume
}

// AttachRequest attaches a volume to a worker
type AttachRequest struct {
	common.AsyncRequest
	// Deployment ID
	DeploymentID string `json:"deployment_id"`
	// Worker ID to attach
	WorkerID string `json:"snapshot_id,omitempty"`
	// Volume ID to attach
	VolumeID string `json:"volume_id"`
	// Path to mount volume at
	MounthPath string `json:"mount_path"`
}

// AttachResponse response from attach request
type AttachResponse struct {
	common.AsyncResponse
}

// AttachCallback is callback response of the attach request
type AttachCallback struct {
	common.AsyncResponse
}

// CreateAttachRequest request to create a new volume and attach it to a worker
type CreateAttachRequest struct {
	// ID of deployment
	DeploymentID string `json:"deployment_id"`
	// ID of worker to attach to
	WorkerID string `json:"worker_id"`
	// Create a volume from existing snapshot
	SnapshotID string `json:"snapshot_id,omitempty"`
	// Volume specification
	VolumeSpec `json:"volume_spec,omitempty"`
	common.AsyncRequest
}

// CreateAttachResponse is response from create attach request
type CreateAttachResponse struct {
	common.AsyncResponse
	// ID of volume
	VolumeID string `json:"volume_id"`
}

// CreateAttachCallback is callback response of the create attach request
type CreateAttachCallback struct {
	common.AsyncResponse
	// Details of volume
	Volume
}

// DetachRequest is a request to detach a volume from a worker
type DetachRequest struct {
	common.AsyncRequest
	// ID of deployment
	DeploymentID string `json:"deployment_id"`
	// ID of volume to detach
	VolumeID string `json:"volumes_id"`
}

// DetachResponse is response from detach request
type DetachResponse struct {
	common.AsyncResponse
}

// DetachCallback is callback response of the detach request
type DetachCallback struct {
	common.AsyncResponse
}

// DeleteRequest is used to delete a volume
type DeleteRequest struct {
	common.AsyncRequest
	// ID of deployment
	DeploymentID string `json:"deployment_id"`
	// IDs of volumes to delete
	VolumeIDs []string `json:"volume_id"`
	// Whether to force deletion or not
	Force bool `json:"force,omitempty"`
}

// DeleteResponse is response from delete request
type DeleteResponse struct {
	common.AsyncResponse
}

// DeleteResult is used in delete callback
type IDResult struct {
	// Used in delete callback
	ID string `json:"id,omitempty"`
	common.APIResponse
}

// DeleteCallback is the callback for delete request
type DeleteCallback struct {
	common.AsyncResponse
	// VolumeID is returned if deletion is successful
	VolumeIDs []IDResult `json:"volume_ids,omitempty"`
}

// CopyRequest is used to copy volumes
type CopyRequest struct {
	// ID of deployment
	DeploymentID string `json:"deployment_id"`
	// ID of volume
	VolumeID string `json:"volume_id"`
	// Target deployment ID to copy to
	TargetDeploymentID string `json:"target_deployment_id"`
	// Number of copies
	NumCopies int `json:"num_copies,omitempty"`
	common.AsyncRequest
}

// CopyResponse is response from copy request
type CopyResponse struct {
	common.AsyncResponse
	// IDs of volume
	VolumeIDs []string `json:"volume_ids,omitempty"`
}

// CopyCallback is callback response of the copy request
type CopyCallback struct {
	common.AsyncResponse
	// IDs of volume that were copied
	VolumeIDs []IDResult `json:"volume_ids,omitempty"`
}

// Describe gets information about a volume
func (describeRequest *DescribeRequest) Describe() (volume Volume, apiErr *common.Error) {

	describeResponse := DescribeResponse{}
	err := common.Execute("VolumesDescribe", describeRequest, &describeResponse)

	if err != nil {
		return Volume{}, err
	}

	if describeResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        describeResponse.StatusCode,
			ErrorDescription: describeResponse.Description,
		}

		return Volume{}, apiErr
	}

	return describeResponse.Volume, nil
}

// Create is used to createa new volume, this will be an empty volume that is not attached to workers
func (createRequest *CreateRequest) Create() (volumeID, requestID string, apiErr *common.Error) {

	createResponse := CreateResponse{}
	err := common.Execute("VolumesCreate", createRequest, &createResponse)

	if err != nil {
		return "", "", &common.ErrInvalidResponseFromAPI
	}

	if createResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        createResponse.StatusCode,
			ErrorDescription: createResponse.Description,
		}

		return "", "", apiErr
	}

	return createResponse.VolumeID, createResponse.RequestToken, nil
}

// Attach an available volume to a running or a stopped worker
func (attachRequest *AttachRequest) Attach() (requestID string, apiErr *common.Error) {

	attachResponse := AttachResponse{}
	err := common.Execute("VolumesAttach", attachRequest, &attachResponse)

	if err != nil {
		return "", err
	}

	if attachResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        attachResponse.StatusCode,
			ErrorDescription: attachResponse.Description,
		}

		return "", apiErr
	}

	return attachResponse.RequestToken, nil
}

// CreateAttach creates a new volume and attaches it to a worker
func (createAttachRequest *CreateAttachRequest) CreateAttach() (volumeID, requestID string, apiErr *common.Error) {

	createAttachResponse := CreateAttachResponse{}
	err := common.Execute("VolumesCreateAttach", createAttachRequest, &createAttachResponse)

	if err != nil {
		return "", "", &common.ErrInvalidResponseFromAPI
	}

	if createAttachResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        createAttachResponse.StatusCode,
			ErrorDescription: createAttachResponse.Description,
		}

		return "", "", apiErr
	}

	return volumeID, createAttachResponse.RequestToken, nil
}

// Detach detaches a volume from a worker
func (detachRequest *DetachRequest) Detach() (requestID string, apiErr *common.Error) {

	detachResponse := DetachResponse{}
	err := common.Execute("VolumesDetach", detachRequest, &detachResponse)

	if err != nil {
		return "", &common.ErrInvalidResponseFromAPI
	}

	if detachResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        detachResponse.StatusCode,
			ErrorDescription: detachResponse.Description,
		}

		return "", apiErr
	}

	return detachResponse.RequestToken, nil
}

// Delete permanently deletes a volume, this action is not reversible
func (deleteRequest *DeleteRequest) Delete() (requestID string, apiErr *common.Error) {

	deleteResponse := DeleteResponse{}
	err := common.Execute("VolumesDelete", deleteRequest, &deleteResponse)

	if err != nil {
		return "", err
	}

	if deleteResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        deleteResponse.StatusCode,
			ErrorDescription: deleteResponse.Description,
		}

		return "", apiErr
	}

	return deleteResponse.RequestToken, nil
}

// Copy creates specified number of copies of the volume either on same deployment or on
// different deployment specified by target deployment id
func (copyRequest *CopyRequest) Copy() (volumeIDs []string, requestID string, apiErr *common.Error) {

	copyResponse := CopyResponse{}
	err := common.Execute("VolumesCopy", copyRequest, &copyResponse)

	if err != nil {
		return nil, "", err
	}

	if copyResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        copyResponse.StatusCode,
			ErrorDescription: copyResponse.Description,
		}

		return nil, "", apiErr
	}

	return copyResponse.VolumeIDs, copyResponse.RequestToken, nil
}
