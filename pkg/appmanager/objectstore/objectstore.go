package objectstore

import (
	"time"

	"go.semut.io/sdk/go-sdk/pkg/common"
)

// Object is used to describe an object that is stored in object store
type Object struct {
	// Deployment ID
	DeploymentID string `json:"deployment_id"`
	// Worker group ID
	WorkerGroupID string `json:"worker_group_id,omitempty"`
	// Worker ID
	WorkerID string `json:"worker_id"`
	// Unique ID of object
	ObjectID string `json:"object_id"`
	// Full name of the object
	FullName string `json:"full_name"`
	// Size of the object
	Size uint64 `json:"size"`
	// Time at which object was created
	CreatedAt time.Time `json:"created_at"`
	// Time at which object was modified
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// If object is encrypted or not
	Encrypted bool `json:"encrypted"`
}

// ListRequest list all objects where path or name is or begins with Name
type ListRequest struct {
	// ID of the deployment whose objects are to be listed
	DeploymentID string `json:"deployment_id"`
	// Find objects where path or name begins with or is equal to
	Pattern string `json:"pattern,omitempty"`
}

// ListResponse list of all objects
type ListResponse struct {
	common.APIResponse
	// List of objects
	Objects []Object `json:"objects,omitempty"`
}

// DescribeRequest get details of a specific object
type DescribeRequest struct {
	DeploymentID string `json:"deployment_id"`
	// ID of the object
	ObjectID string `json:"object_id,omitempty"`
	// Name of the object including its path
	FullName string `json:"full_name,omitempty"`
}

// DescribeResponse is object meta returned from Platform API
type DescribeResponse struct {
	common.APIResponse
	Object
}

// ObjectRequest object to upload
type ObjectRequest struct {
	// Directory of the source object
	SourceDir string `json:"source_dir"`
	// Name of the source object
	SourceName string `json:"source_name"`
	// Destination directory. Omit to use same path as SourcePath. Use "/" to upload to root folder
	DestinationDir string `json:"destination_dir"`
	// New name of the object. Omit to use same name as SourceName
	DestinationName string `json:"destination_name"`
	// Encrypted whether object will be stored with inbuilt encryption or not
	Encrypted bool `json:"encrypted"`
}

// DeleteRequest is used to request deletion of an object
type DeleteRequest struct {
	ObjectID string `json:"object_id"`
}

// DeleteResponse response to object deletion request
type DeleteResponse struct {
	common.APIResponse
}

// List returns a specific object, or list of objects that match the name pattern or all objects in a deployment
func (listRequest *ListRequest) List() (objects []Object, apiErr *common.Error) {

	listResponse := ListResponse{}
	err := common.Execute("ObjectStoreList", listRequest, &listResponse)

	if err != nil {
		return nil, err
	}

	if listResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        listResponse.StatusCode,
			ErrorDescription: listResponse.Description,
		}

		return nil, apiErr
	}

	return listResponse.Objects, nil
}

// Describe a specific object that matches either the ObjectID or the full name of the object
func (describeRequest *DescribeRequest) Describe() (object Object, apiErr *common.Error) {

	describeResponse := DescribeResponse{}
	err := common.Execute("ObjectStoreDescribe", describeRequest, &describeResponse)

	if err != nil {
		return Object{}, err
	}

	if describeResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        describeResponse.StatusCode,
			ErrorDescription: describeResponse.Description,
		}

		return Object{}, apiErr
	}

	return describeResponse.Object, nil
}

// Delete is used to delete an already stored object, this action is not reverseable
func (deleteRequest *DeleteRequest) Delete() (success bool, apiErr *common.Error) {

	deleteResponse := DeleteResponse{}
	err := common.Execute("ObjectStoreDelete", deleteRequest, &deleteResponse)

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
