package secrets

import (
	appmgrSecrets "go.semut.io/sdk/go-sdk/pkg/appmanager/secrets"
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// StoreSecretRequest is used to request storing a secret string from the worker end
type StoreSecretRequest appmgrSecrets.StoreSecretRequest

// RetrieveSecretRequest is used to request retrieval of an already stored secret string from the worker end
type RetrieveSecretRequest appmgrSecrets.RetrieveSecretRequest

// DeleteSecretRequest is used to request secret deletion from worker end
type DeleteSecretRequest appmgrSecrets.DeleteSecretRequest

// GenerateCredentialsRequest is used to request random secret generation from worker end
type GenerateCredentialsRequest appmgrSecrets.GenerateCredentialsRequest

// GenerateStoreCredentialsRequest is used to request ranndom secret generation and then
// also store it from worker end
type GenerateStoreCredentialsRequest appmgrSecrets.GenerateStoreCredentialsRequest

// StoreSecret is an alias of app manager store secret
func (storeSecretRequest *StoreSecretRequest) StoreSecret() (success bool, apiErr *common.Error) {
	ssr := (*appmgrSecrets.StoreSecretRequest)(storeSecretRequest)
	success, apiErr = ssr.StoreSecret()
	return
}

// RetrieveSecret is an alias of app manager retrieve secret

func (retrieveSecretRequest *RetrieveSecretRequest) RetrieveSecret() (secret string, apiErr *common.Error) {
	rsr := (*appmgrSecrets.RetrieveSecretRequest)(retrieveSecretRequest)
	secret, apiErr = rsr.RetrieveSecret()
	return
}

// DeleteSecret is an alias of app manager delete secret

func (deleteSecretRequest *DeleteSecretRequest) DeleteSecret() (success bool, apiErr *common.Error) {
	dsr := (*appmgrSecrets.DeleteSecretRequest)(deleteSecretRequest)
	success, apiErr = dsr.DeleteSecret()
	return
}

// GenerateCredentials is an alias of app manager secret generate credentials
func (generateRequest *GenerateCredentialsRequest) GenerateCredentials() (secret string, apiErr *common.Error) {
	gr := (*appmgrSecrets.GenerateCredentialsRequest)(generateRequest)
	secret, apiErr = gr.GenerateCredentials()
	return
}

// GenerateStoreCredentials is an alias of app manager secret generate-store credentials
func (generateStoreRequest *GenerateStoreCredentialsRequest) GenerateStoreCredentials() (success bool, apiErr *common.Error) {
	gsr := (*appmgrSecrets.GenerateStoreCredentialsRequest)(generateStoreRequest)
	success, apiErr = gsr.GenerateStoreCredentials()
	return
}
