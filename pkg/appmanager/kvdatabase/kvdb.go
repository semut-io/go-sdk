package kvdatabase

import (
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// Record represents a single record in the database
type Record struct {
	// Key of the key value pair
	Key string `json:"key"`
	// Value of the key value pair
	Value string `json:"value"`
}

// GetRequest keys to get from DB
type GetRequest struct {
	// Multiple keys
	// values of all keys will be retrieved in case of passing empty slice
	Keys []string `json:"keys"`
}

// GetResponse records returned from DB
type GetResponse struct {
	common.APIResponse
	// Multiple key value pair records
	Records []Record `json:"records,omitempty"`
}

// SetRequest records to set in database
type SetRequest struct {
	Records []Record `json:"records"`
}

// SetResponse Set response received from Platform API
type SetResponse struct {
	common.APIResponse
}

// DeleteRequest records to set in database
type DeleteRequest struct {
	// List of keys to delete from database
	Keys []string `json:"keys"`
}

// DeleteResponse records to set in database
type DeleteResponse struct {
	common.APIResponse
	// Number of keys that were deleted from database
	NumKeysDeleted int `json:"num_keys_deleted,omitempty"`
}

// RenameKey key to rename
type RenameKey struct {
	// Existing name of the key
	ExistingName string `json:"existing_name"`
	// New name of the key
	NewName string `json:"new_name"`
	// Overwrite will overwrite the key with the new name,
	// even if a key with the new name exists
	Overwrite bool `json:"overwrite"`
}

// RenameKeyResult result of rename keys
type RenameKeyResult struct {
	// Existing name of the key
	ExistingName string `json:"existing_name"`
	// New name of the key
	NewName string `json:"new_name"`
	// Key value pair
	Result bool `json:"result"`
}

// RenameRequest rename keys
type RenameRequest struct {
	// List of key value pair
	RenameKeys []RenameKey `json:"rename_keys"`
}

// RenameResponse is the response of the rename request
type RenameResponse struct {
	common.APIResponse
	// List of renamed keys result
	RenameKeys []RenameKeyResult `json:"rename_keys"`
}

// ListRequest list keys that match the pattern
type ListRequest struct {
	KeyPrefix string `json:"key_prefix"`
}

// ListResponse return keys that match the prefix in request
type ListResponse struct {
	common.APIResponse
	// List of key value pairs
	Records []Record `json:"records"`
}

// Get one or more keys from database
func (getRequest *GetRequest) Get() (records []Record, apiErr *common.Error) {

	getResponse := GetResponse{}

	err := common.Execute("DatabaseGet", getRequest, &getResponse)

	if err != nil {
		return nil, err
	}

	if getResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        getResponse.StatusCode,
			ErrorDescription: getResponse.Description,
		}

		return nil, apiErr
	}

	return getResponse.Records, nil
}

// Set one or more keys in the Database. This operation is atomic and will overwrite if keys exist
func (setRequest *SetRequest) Set() (apiErr *common.Error) {

	setResponse := SetResponse{}

	err := common.Execute("DatabaseSet", setRequest, &setResponse)

	if err != nil {
		return err
	}

	if setResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        setResponse.StatusCode,
			ErrorDescription: setResponse.Description,
		}

		return apiErr
	}

	return nil
}

// Delete existing records from database
func (deleteRequest *DeleteRequest) Delete() (numKeysDeleted int, apiErr *common.Error) {

	deleteResponse := DeleteResponse{}

	err := common.Execute("DatabaseDelete", deleteRequest, &deleteResponse)

	if err != nil {
		return 0, err
	}

	if deleteResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        deleteResponse.StatusCode,
			ErrorDescription: deleteResponse.Description,
		}

		return 0, apiErr
	}

	return deleteResponse.NumKeysDeleted, nil
}

// Rename records in the database for specific keys
func (renameRequest *RenameRequest) Rename() (renameKeyResult []RenameKeyResult, apiErr *common.Error) {

	renameResponse := RenameResponse{}
	err := common.Execute("DatabaseRename", renameRequest, &renameResponse)

	if err != nil {
		return nil, err
	}

	if renameResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        renameResponse.StatusCode,
			ErrorDescription: renameResponse.Description,
		}

		return nil, apiErr
	}

	return renameResponse.RenameKeys, nil
}

// List records from the database
func (listRequest *ListRequest) List() (records []Record, apiErr *common.Error) {

	listResponse := ListResponse{}
	err := common.Execute("DatabaseList", listRequest, &listResponse)

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

	return listResponse.Records, err
}
