package encryption

import (
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// EncryptContentRequest is used to request encryption of some content
type EncryptContentRequest struct {
	// String to encrypt
	Content string `json:"content"`
}

// EncryptContentResponse is the response to the encrypt content request
type EncryptContentResponse struct {
	common.APIResponse
	// Encrypted content as string
	EncryptedContent string `json:"encrypted_content"`
	// Key identified used during encryption
	EncryptionKeyIdentifier string `json:"encryption_key_identifier"`
}

// DecryptContentRequest derypt content request
type DecryptContentRequest struct {
	// Encrypted content
	Content string `json:"encrypted_content"`
	// Key identified used during encryption
	EncryptionKeyIdentifier string `json:"encryption_key_identifier"`
}

// DecryptContentResponse decrypted content response from request
type DecryptContentResponse struct {
	common.APIResponse
	// Decrypted content
	Content string `json:"content"`
}

// DeleteEncryptionKeyRequest is used to request deletion of an existing encryption key
type DeleteEncryptionKeyIdentifierRequest struct {
	// EncryptionKeyIdentifier is the key used to decrypt the file that was generated
	// at the time of encryption
	EncryptionKeyIdentifier string `json:"encryption_key_identifier"`
}

// DeleteEncryptionKeyResponse is the response received after key is deleted
type DeleteEncryptionKeyIdentifierResponse struct {
	common.APIResponse
}

// EncryptContent ecrypts and returns content with encryption key identifier. This is useful for encrypting as well as storing secrets
func (encryptContentRequest *EncryptContentRequest) EncryptContent() (encryptedContent, encryptionKeyID string, apiErr *common.Error) {

	encryptContentResponse := EncryptContentResponse{}
	err := common.Execute("EncryptionEncryptContent", encryptContentRequest, &encryptContentResponse)

	if err != nil {
		return "", "", err
	}

	if encryptContentResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        encryptContentResponse.StatusCode,
			ErrorDescription: encryptContentResponse.Description,
		}

		return "", "", apiErr
	}

	return encryptContentResponse.EncryptedContent, encryptContentResponse.EncryptionKeyIdentifier, nil
}

// DecryptContent decrypts content and returns unencrypted content
func (decryptContentRequest *DecryptContentRequest) DecryptContent() (content string, apiErr *common.Error) {

	decryptContentResponse := DecryptContentResponse{}
	err := common.Execute("EncryptionEncryptContent", decryptContentRequest, &decryptContentResponse)

	if err != nil {
		return "", err
	}

	if decryptContentResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        decryptContentResponse.StatusCode,
			ErrorDescription: decryptContentResponse.Description,
		}

		return "", apiErr
	}

	return decryptContentResponse.Content, nil
}

// DeleteEncryptionKey is used to delete the encryption key that was previously generated during
// any file encryption
func (dekRequest *DeleteEncryptionKeyIdentifierRequest) DeleteEncryptionKeyIdentifier() (success bool, apiErr *common.Error) {

	dekResponse := DeleteEncryptionKeyIdentifierResponse{}
	err := common.Execute("EncryptionDeleteKey", dekRequest, &dekResponse)

	if err != nil {
		return false, err
	}

	if dekResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        dekResponse.StatusCode,
			ErrorDescription: dekResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}
