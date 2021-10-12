package cron

import (
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// Spec is the complete specification of a cron job
type Spec struct {
	// Schedule of the cron job as per linux crontab format
	Schedule string `json:"schedule"`
	// URI where to pass callback
	CallbackURL string `json:"callback_url"`
	// Content to pass to the callback URI
	CallbackMessageContent string `json:"callback_message_content"`
}

// Cron is a cron job with an ID
type Cron struct {
	// Unique ID of the cron job
	CronID string `json:"cron_id"`
	Spec
}

// CreateRequest create new cron
type CreateRequest struct {
	Spec
}

// CreateResponse response to the cron creation request
type CreateResponse struct {
	common.APIResponse
	// Unique ID of the cron
	CronID string `json:"cron_id"`
}

// ListRequest request list of all cron jobs
type ListRequest struct{}

// ListResponse is the response of list request
type ListResponse struct {
	common.APIResponse
	// List of crons
	CronList []Cron `json:"cron_list,omitempty"`
}

// UpdateRequest update cron
type UpdateRequest struct {
	// Unique ID of cron job to update
	CronID string `json:"cron_id"`
	Spec
}

// UpdateResponse response of update request
type UpdateResponse struct {
	common.APIResponse
}

// DeleteRequest delete a cron job
type DeleteRequest struct {
	// ID of the cron to delete
	CronID string `json:"cron_id"`
}

// DeleteResponse response of delete request
type DeleteResponse struct {
	common.APIResponse
}

// Create is used to create a new cron job
func (createRequest *CreateRequest) Create() (cronID string, apiErr *common.Error) {

	createResponse := CreateResponse{}
	err := common.Execute("CronCreate", createRequest, &createResponse)

	if err != nil {
		return "", &common.ErrInvalidResponseFromAPI
	}

	if createResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        createResponse.StatusCode,
			ErrorDescription: createResponse.Description,
		}

		return "", apiErr
	}

	return createResponse.CronID, nil
}

// List is to used to fetch a listing of all existing cron jobs
func (listRequest *ListRequest) List() (cronList []Cron, apiErr *common.Error) {

	listResponse := ListResponse{}
	err := common.Execute("CronList", listRequest, &listResponse)

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

	return listResponse.CronList, nil

}

// Update is used to update all fields of an existing cron job
func (updateRequest *UpdateRequest) Update() (sucess bool, apiErr *common.Error) {

	updateResponse := UpdateResponse{}
	err := common.Execute("CronUpdate", updateRequest, &updateResponse)

	if err != nil {
		return false, err
	}

	if updateResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        updateResponse.StatusCode,
			ErrorDescription: updateResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// Delete is used to delete an existing cron job, allowing it to stop executing anymore
func (deleteRequest *DeleteRequest) Delete() (sucess bool, apiErr *common.Error) {

	deleteResponse := DeleteResponse{}
	err := common.Execute("CronDelete", deleteRequest, &deleteResponse)

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
