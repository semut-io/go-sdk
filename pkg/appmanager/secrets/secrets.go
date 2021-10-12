package secrets

import (
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// SecretSpec is the spec containing secret key and optional Deployment ID
type SecretSpec struct {
	// Optional ID of the deployment
	DeploymentID string `json:"deployment_id,omitempty"`
	// SecretName is the identifier of the secret
	SecretKey string `json:"secret_key"`
}

// CredentialFormat is used to define the format of the credential to be generated.
type CredentialFormat struct {
	// Type is used to denote the type of credential to be generated.
	Type common.CredentialType `json:"type"`
	// Length is total size of the credential to be generated if Type is AlphaNumericType
	// or HexadecimalType. A predefined Length is used if Type is SHA256Type. For RegexType,
	// Length denotes the maximum number of times star, range or plus should be repeated.
	// If Length is not provided, then it is randomly generated.
	Length int `json:"length,omitempty"`
	// RegexFormat is used when the Type is RegexType.
	RegexFormat string `json:"regex_format,omitempty"`
}

// StoreSecretRequest is used to store a secret
type StoreSecretRequest struct {
	SecretSpec
	// Secret is the conent to store.
	SecretValue string `json:"secret_key"`
	// Overwrite when set to true, will replace any existing secret
	Overwrite bool `json:"overwrite"`
}

// StoreSecretResponse is the response from the store request
type StoreSecretResponse struct {
	common.APIResponse
}

// RetrieveSecretRequest retrive stored secret
type RetrieveSecretRequest struct {
	SecretSpec
}

// RetrieveSecretResponse is the response from the retrieval request
type RetrieveSecretResponse struct {
	common.APIResponse
	// Secret is the conent to store.
	Secret string `json:"secret_value,omitempty"`
}

// DeleteSecretRequest retrive stored secret
type DeleteSecretRequest struct {
	SecretSpec
}

// DeleteSecretResponse is the response of the deletion request
type DeleteSecretResponse struct {
	common.APIResponse
}

// GenerateCredentialsRequest generate a secret credential
type GenerateCredentialsRequest struct {
	DeploymentID string `json:"deployment_id,omitempty"`
	// CredentialFormat is the format used for generating the secret.
	CredentialFormat
}

// GenerateCredentialsResponse is the response of the generate request
type GenerateCredentialsResponse struct {
	common.APIResponse
	// Secret is the generated secret
	Secret string `json:"secret_value,omitempty"`
}

// GenerateStoreCredentialsRequest generate a secret credential
// and store it to secret store
type GenerateStoreCredentialsRequest struct {
	SecretSpec
	// CredentialFormat is the format used for generating the secret.
	CredentialFormat
}

// GenerateStoreCredentialsResponse is the response of the generate-store request
type GenerateStoreCredentialsResponse struct {
	common.APIResponse
}

// StoreSecret stores secret in the secrets store
func (storeSecretRequest *StoreSecretRequest) StoreSecret() (success bool, apiErr *common.Error) {

	storeSecretResponse := StoreSecretResponse{}
	err := common.Execute("SecretsStore", storeSecretRequest, &storeSecretResponse)

	if err != nil {
		return false, err
	}

	if storeSecretResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        storeSecretResponse.StatusCode,
			ErrorDescription: storeSecretResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// RetrieveSecret retrieves secret that was previously stored in the secrets store
func (retrieveSecretRequest *RetrieveSecretRequest) RetrieveSecret() (secret string, apiErr *common.Error) {

	retrieveSecretResponse := RetrieveSecretResponse{}
	err := common.Execute("SecretsRetrieve", retrieveSecretRequest, &retrieveSecretResponse)

	if err != nil {
		return "", err
	}

	if retrieveSecretResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        retrieveSecretResponse.StatusCode,
			ErrorDescription: retrieveSecretResponse.Description,
		}

		return "", apiErr
	}

	return retrieveSecretResponse.Secret, nil
}

// DeleteSecret deletes stored secret
func (deleteSecretRequest *DeleteSecretRequest) DeleteSecret() (success bool, apiErr *common.Error) {

	deleteSecretResponse := DeleteSecretResponse{}
	err := common.Execute("SecretsDelete", deleteSecretRequest, &deleteSecretResponse)

	if err != nil {
		return false, err
	}

	if deleteSecretResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        deleteSecretResponse.StatusCode,
			ErrorDescription: deleteSecretResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// GenerateCredentials generates random credentials of specified format
func (generateRequest *GenerateCredentialsRequest) GenerateCredentials() (secret string, apiErr *common.Error) {

	generateResponse := GenerateCredentialsResponse{}
	err := common.Execute("SecretsGenerateCredentials", generateRequest, &generateResponse)

	if err != nil {
		return "", err
	}

	if generateResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        generateResponse.StatusCode,
			ErrorDescription: generateResponse.Description,
		}

		return "", apiErr
	}

	return generateResponse.Secret, nil
}

// GenerateStoreCredentials generates random credentials and stores them in ecrypted store
func (generateStoreRequest *GenerateStoreCredentialsRequest) GenerateStoreCredentials() (success bool, apiErr *common.Error) {

	generateStoreResponse := GenerateStoreCredentialsResponse{}
	err := common.Execute("SecretsGenerateStoreCredentials", generateStoreRequest, &generateStoreResponse)

	if err != nil {
		return false, err
	}

	if generateStoreResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        generateStoreResponse.StatusCode,
			ErrorDescription: generateStoreResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}
