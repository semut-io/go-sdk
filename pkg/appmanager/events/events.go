package events

import (
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// Criterion is a single criterion to be evaluated
type Criterion struct {
	// Unique name of the criteria
	Name string `json:"name"`
	// Field on which criteria needs to be applied. This can be field in a log or name of the metric being
	// evaluated.
	Field string `json:"field"`
	// Operand can be any number comparison operator, string evaluation function or regex.
	Operand string `json:"operand"`
	// A number, string or regex pattern. This Value is treated as default value. Users can modify value
	// during subscription if IsValueModifiable is true.
	Value string `json:"value"`
}

// EventCriteria overall criteria to determine if the event has been observed
type EventCriteria struct {
	// Criteria to evaluate
	Criteria []Criterion `json:"criteria"`
	// Equation is a boolean expression of all criteria. Eg CPUBreach OR RAMBreach. CPUBreach and RAMBreach are two Criteria. Empty equation will be treated as and condition
	Equation string `json:"equation"`
}

// EventSpec is the specification of an event
type EventSpec struct {
	// Name of the event
	Name string `json:"name"`
	// Description of the event
	Description string `json:"description"`
	// Type of the event
	Type common.EventType `json:"type"`
	EventCriteria
}

// Event is used to describe an event
type Event struct {
	// Unique ID of the event
	EventID string `json:"event_id,omitempty"`
	EventSpec
}

// Subscription A subscribed event
type SubscriptionSpec struct {
	EventID                string `json:"event_id"`
	CallbackURL            string `json:"callback_url"`
	CallbackMessageContent string `json:"callback_message_content"`
	DeploymentID           string `json:"deployment_id"`
	WorkerGroupID          string `json:"worker_group_id,omitempty"`
	WorkerID               string `json:"worker_id,omitempty"`
}

type Subscription struct {
	SubscriptionID string `json:"subscription_id"`
	SubscriptionSpec
}

// NewRequest new event request
type NewRequest struct {
	EventSpec
}

// NewResponse new event response
type NewResponse struct {
	common.APIResponse
	// ID of the event created
	EventID string `json:"event_id,omitempty"`
}

// UpdateRequest updates an existing event
type UpdateRequest struct {
	EventID string `json:"event_id"`
	EventSpec
}

// UpdateResponse response to the update request
type UpdateResponse struct {
	common.APIResponse
}

// DeleteRequest event to be deleted
type DeleteRequest struct {
	// ID of event to be deleted
	EventID string `json:"event_id,omitempty"`
}

// DeleteResponse event deletion result
type DeleteResponse struct {
	common.APIResponse
}

// ListRequest list events of certain type
type ListRequest struct {
	// Type of event - metric, logging,
	Type common.EventType `json:"event_type,omitempty"`
}

// ListResponse list of events
type ListResponse struct {
	common.APIResponse
	// List of events
	Events []Event `json:"events,omitempty"`
}

// NewSubscriptionRequest subscribe to an event
type NewSubscriptionRequest struct {
	SubscriptionSpec
}

// NewSubscriptionResponse subscription response
type NewSubscriptionResponse struct {
	common.APIResponse
	// ID of subscription
	SubscriptionID string `json:"subscription_id"`
}

// DeleteSubscriptionRequest unsubscribe to an event
type DeleteSubscriptionRequest struct {
	// ID of subscription
	SubscriptionID string `json:"subscription_id"`
}

// DeleteSubscriptionResponse Unsubscribe result
type DeleteSubscriptionResponse struct {
	common.APIResponse
}

// ListSubscriptionRequest list all subscribed events
type ListSubscriptionRequest struct{}

// ListSubscriptionResponse list of subscribed events
type ListSubscriptionResponse struct {
	common.APIResponse
	// List of subscription
	Subscriptions []Subscription `json:"subscriptions"`
}

// New creates a new user defined event
func (newRequest *NewRequest) New() (eventID string, apiErr *common.Error) {

	newResponse := NewResponse{}
	err := common.Execute("EventsNew", newRequest, &newResponse)

	if err != nil {
		return "", &common.ErrInvalidResponseFromAPI
	}

	if newResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        newResponse.StatusCode,
			ErrorDescription: newResponse.Description,
		}

		return "", apiErr
	}

	return newResponse.EventID, nil
}

// Update will update event
func (updateRequest *UpdateRequest) Update() (apiErr *common.Error) {
	updateResponse := UpdateResponse{}
	err := common.Execute("EventsUpdate", updateRequest, &updateResponse)

	if err != nil {
		return &common.ErrInvalidResponseFromAPI
	}

	if updateResponse.StatusCode != "200" {
		apiErr = &common.Error{
			ErrorCode:        updateResponse.StatusCode,
			ErrorDescription: updateResponse.Description,
		}
		return apiErr
	}

	return nil
}

// Delete an event
func (deleteRequest *DeleteRequest) Delete() (success bool, apiErr *common.Error) {

	deleteResponse := DeleteResponse{}
	err := common.Execute("EventsDelete", deleteRequest, &deleteResponse)

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

// List all events
func (listRequest *ListRequest) List() (events []Event, apiErr *common.Error) {

	listResponse := ListResponse{}
	err := common.Execute("EventsList", listRequest, &listResponse)

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

	return listResponse.Events, nil
}

// NewSubscription is used to subscribe to an event
func (subscribeRequest *NewSubscriptionRequest) NewSubscription() (subscriptionID string, apiErr *common.Error) {

	subscribeResponse := NewSubscriptionResponse{}
	err := common.Execute("EventSubscriptionsNew", subscribeRequest, &subscribeResponse)

	if err != nil {
		return "", &common.ErrInvalidResponseFromAPI
	}

	if subscribeResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        subscribeResponse.StatusCode,
			ErrorDescription: subscribeResponse.Description,
		}

		return "", apiErr
	}

	return subscribeResponse.SubscriptionID, nil
}

// DeleteSubscription is used to unsubscribe from an event
func (unsubscribeRequest *DeleteSubscriptionRequest) DeleteSubscription() (success bool, apiErr *common.Error) {

	unsubscribeResponse := DeleteSubscriptionResponse{}
	err := common.Execute("EventSubscriptionsDelete", unsubscribeRequest, &unsubscribeResponse)

	if err != nil {
		return false, err
	}

	if unsubscribeResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        unsubscribeResponse.StatusCode,
			ErrorDescription: unsubscribeResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// ListSubscription is used to list all event subscriptions
func (ListSubscriptionRequest *ListSubscriptionRequest) ListSubscriptions() (subscription []Subscription, apiErr *common.Error) {

	listSubscriptionResponse := ListSubscriptionResponse{}
	err := common.Execute("EventSubscriptionsList", ListSubscriptionRequest, &listSubscriptionResponse)

	if err != nil {
		return nil, err
	}

	if listSubscriptionResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        listSubscriptionResponse.StatusCode,
			ErrorDescription: listSubscriptionResponse.Description,
		}

		return nil, apiErr
	}

	return listSubscriptionResponse.Subscriptions, nil
}
