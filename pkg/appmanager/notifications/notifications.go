package notifications

import (
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// NotificationTemplate contains email template
type NotificationTemplate struct {
	// To receiver of the notification email
	ToEmail string `json:"to_email"`
	// Subject of the email
	Subject string `json:"subject"`
	// Body of the email (may contain placeholder keys)
	Body string `json:"body"`
}

// Notification contains notification ID and template
type Notification struct {
	NotificationTemplate
	// Unqiue ID of notification
	NotificationID string `json:"notification_id"`
}

// EmailTemplateContent content to be replaced by placeholders
type EmailTemplateContent struct {
	// Key to look for in email message
	PlaceholderKey string `json:"placeholder_key"`
	// Value to replace key with during sending
	ReplacementValue string `json:"replacement_value"`
}

// EmailSpec used for converting email template into email content
type EmailSpec struct {
	EmailTemplateContent []EmailTemplateContent `json:"email_template_content"`
}

// AddEmailRequest is used to add a new email notification template
type AddEmailRequest struct {
	NotificationTemplate
}

// AddEmailResponse is response to add email request
type AddEmailResponse struct {
	common.APIResponse
}

// SendEmailRequest is used to send an email with notification template
type SendEmailRequest struct {
	NotificationID string `json:"notification_id"`
	// Specification
	EmailSpec
}

// SendEmailResponse is response to send request
type SendEmailResponse struct {
	common.APIResponse
}

// DeleteEmailRequest requests deletion of email template
type DeleteEmailRequest struct {
	// Unique ID of notification to delete
	NotificationID string `json:"notification_id"`
}

// DeleteEmailResponse is response to delete request
type DeleteEmailResponse struct {
	common.APIResponse
}

// UpdateEmailRequest request email template update
type UpdateEmailRequest struct {
	// Unique ID of notification to update
	NotificationID string `json:"notification_id"`
	NotificationTemplate
}

// UpdateEmailResponse response to update email request
type UpdateEmailResponse struct {
	common.APIResponse
}

// ListEmailRequest is used to request a list of email templates
type ListEmailRequest struct{}

// ListEmailResponse is used to response to list request
type ListEmailResponse struct {
	common.APIResponse
	// NotificationTemplate is the template of the Notification Email that can be sent using the service.
	NotificationList []Notification `json:"notification_list,omitempty"`
}

// VerifyEmailIDRequest is used to request verification of an email ID
type VerifyEmailIDRequest struct {
	EmailID string `json:"email_id"`
}

// VerifyEmailResponse is response to request
type VerifyEmailIDResponse struct {
	common.APIResponse
}

// ListVerifiedEmailIDsRequest is the request to listing of verified email IDs
type ListVerifiedEmailIDsRequest struct{}

// ListVerifiedEmailIDsResponse is response to request
type ListVerifiedEmailIDsResponse struct {
	common.APIResponse
	// List of IDs
	VerifiedEmailIDs []string `json:"verified_email_ids"`
}

// AddEmail is used to add email notification template
func (addEmailRequest *AddEmailRequest) AddEmail() (success bool, apiErr *common.Error) {

	addEmailResponse := AddEmailResponse{}
	err := common.Execute("NotificationsAddEmail", addEmailRequest, &addEmailResponse)

	if err != nil {
		return false, err
	}

	if addEmailResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        addEmailResponse.StatusCode,
			ErrorDescription: addEmailResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// SendEmail is used to send email based on email notification template
func (sendEmailRequest *SendEmailRequest) SendEmail() (success bool, apiErr *common.Error) {

	sendEmailResponse := SendEmailResponse{}
	err := common.Execute("NotificationsSendEmail", sendEmailRequest, &sendEmailResponse)

	if err != nil {
		return false, err
	}

	if sendEmailResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        sendEmailResponse.StatusCode,
			ErrorDescription: sendEmailResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// DeleteEmail is used to delete an existing email template
func (deleteEmailRequest *DeleteEmailRequest) DeleteEmail() (success bool, apiErr *common.Error) {

	deleteEmailResponse := DeleteEmailResponse{}
	err := common.Execute("NotificationsDeleteEmail", deleteEmailRequest, &deleteEmailResponse)

	if err != nil {
		return false, err
	}

	if deleteEmailResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        deleteEmailResponse.StatusCode,
			ErrorDescription: deleteEmailResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// UpdateEmail is used to update all fields of an existing email template
func (updateEmailRequest *UpdateEmailRequest) UpdateEmail() (success bool, apiErr *common.Error) {

	updateEmailResponse := UpdateEmailResponse{}
	err := common.Execute("NotificationsUpdateEmail", updateEmailRequest, &updateEmailResponse)

	if err != nil {
		return false, err
	}

	if updateEmailResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        updateEmailResponse.StatusCode,
			ErrorDescription: updateEmailResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// List is used to list all email notification templates
func (listEmailRequest *ListEmailRequest) ListEmail() (notificationList []Notification, apiErr *common.Error) {

	listEmailResponse := ListEmailResponse{}
	err := common.Execute("NotificationsListEmail", listEmailRequest, &listEmailResponse)

	if err != nil {
		return nil, err
	}

	if listEmailResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        listEmailResponse.StatusCode,
			ErrorDescription: listEmailResponse.Description,
		}

		return nil, apiErr
	}

	return listEmailResponse.NotificationList, nil
}

// VerifyEmailID is used to verify a specific email ID
func (verifyEmailIDRequest *VerifyEmailIDRequest) VerifyEmailID() (apiErr *common.Error) {

	verifyEmailIDResponse := VerifyEmailIDResponse{}
	err := common.Execute("NotificationsVerifyEmailID", verifyEmailIDRequest, &verifyEmailIDResponse)

	if err != nil {
		return err
	}

	if verifyEmailIDResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        verifyEmailIDResponse.StatusCode,
			ErrorDescription: verifyEmailIDResponse.Description,
		}

		return apiErr
	}

	return nil
}

// ListVerifiedEmailIDs is used to list already verified email IDs
func (listVerifiedEmailIDsRequest *ListVerifiedEmailIDsRequest) ListVerifiedEmailIDs() (idsList []string, apiErr *common.Error) {

	listVerifiedEmailIDsResponse := ListVerifiedEmailIDsResponse{}
	err := common.Execute("NotificationsListVerifiedEmailIDs", listVerifiedEmailIDsRequest, &listVerifiedEmailIDsResponse)

	if err != nil {
		return nil, err
	}

	if listVerifiedEmailIDsResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        listVerifiedEmailIDsResponse.StatusCode,
			ErrorDescription: listVerifiedEmailIDsResponse.Description,
		}

		return nil, apiErr
	}

	return listVerifiedEmailIDsResponse.VerifiedEmailIDs, nil
}
