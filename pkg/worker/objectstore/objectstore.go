package objectstore

import (
	appmgrObjectstore "go.semut.io/sdk/go-sdk/pkg/appmanager/objectstore"
	"go.semut.io/sdk/go-sdk/pkg/common"
)

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

// PostRequest uploads an object prsent in the worker
type PostRequest struct {
	ObjectRequest
}

// PostResponse is the response of the post request
type PostResponse struct {
	// ID of the object
	ObjectID string `json:"object_id"`
	common.APIResponse
}

// GetRequest downloads an object to a given path on the worker
type GetRequest struct {
	ObjectRequest
}

// GetResponse is the response to get object request
type GetResponse struct {
	common.APIResponse
}

// Post is used to upload an object from the worker
func (postRequest *PostRequest) Post() (success bool, apiErr *common.Error) {

	postResponse := PostResponse{}
	err := common.Execute("ObjectStorePost", postRequest, &postResponse)

	if err != nil {
		return false, err
	}

	if postResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        postResponse.StatusCode,
			ErrorDescription: postResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// Get is used to download an object to the worker
func (getRequest *GetRequest) Get() (success bool, apiErr *common.Error) {

	getResponse := GetResponse{}
	err := common.Execute("ObjectStoreGet", getRequest, &getResponse)

	if err != nil {
		return false, err
	}

	if getResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        getResponse.StatusCode,
			ErrorDescription: getResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// DeleteRequest is used to delete an existing object from worker end
type DeleteRequest appmgrObjectstore.DeleteRequest

// Delete is an alias of app manager delete object
func (deleteRequest *DeleteRequest) Delete() (success bool, apiErr *common.Error) {
	dr := (*appmgrObjectstore.DeleteRequest)(deleteRequest)
	success, apiErr = dr.Delete()
	return
}
