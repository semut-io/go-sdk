package encryption

import (
	appmgrEncryption "go.semut.io/sdk/go-sdk/pkg/appmanager/encryption"

	"go.semut.io/sdk/go-sdk/pkg/common"
)

// EncryptFileRequest is used to request encryption of a file present in the worker
type EncryptFileRequest struct {
	// FileOrDirectory is the path of the file or directory that is to be encrypted
	FileOrDirectory string `json:"file_or_directory"`
}

// EncryptFileResponse is the response containing the encryption key after file is encrypted
type EncryptFileResponse struct {
	common.APIResponse
	// EncryptionKeyIdentifier is the unique key that can be used to decrypt the file that was encrypted
	EncryptionKeyIdentifier string `json:"encryption_key_identifier,omitempty"`
}

// DecryptFileRequest is used to request decryption of an already encrypted file present on the worker
type DecryptFileRequest struct {
	// EncryptedFile is the path to the encrypted file that is to be decrypted
	EncryptedFile string `json:"encrypted_file"`
	// EncryptionKeyIdentifier is the key used to decrypt the file that was generated
	// at the time of encryption
	EncryptionKeyIdentifier string `json:"encryption_key_identifier,omitempty"`
}

// DecryptFileResponse is the response received after file is decrypted
type DecryptFileResponse struct {
	common.APIResponse
}

// EncryptFile encrypts the file on the worker and returns the generated key that can be used for decrypting it
func (encryptFileRequest *EncryptFileRequest) EncryptFile() (encryptionKeyIndentifier string, apiErr *common.Error) {

	encryptFileResponse := EncryptFileResponse{}
	err := common.Execute("EncryptionEncryptFile", encryptFileRequest, &encryptFileResponse)

	if err != nil {
		return "", err
	}

	if encryptFileResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        encryptFileResponse.StatusCode,
			ErrorDescription: encryptFileResponse.Description,
		}

		return "", apiErr
	}

	return encryptFileResponse.EncryptionKeyIdentifier, nil
}

// DecryptFile decrypts specified file on the worker using encryption key that was previously encrypted
func (decryptFileRequest *DecryptFileRequest) DecryptFile() (success bool, apiErr *common.Error) {

	decryptFileResponse := DecryptFileResponse{}
	err := common.Execute("EncryptionEncryptFile", decryptFileRequest, &decryptFileResponse)

	if err != nil {
		return false, err
	}

	if decryptFileResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        decryptFileResponse.StatusCode,
			ErrorDescription: decryptFileResponse.Description,
		}

		return false, apiErr
	}

	return true, nil
}

// DeleteRequest is used to delete an existing encryption key
type DeleteEncryptionKeyIdentifierRequest appmgrEncryption.DeleteEncryptionKeyIdentifierRequest

// Delete is an alias of app manager delete encryption key
func (deleteRequest *DeleteEncryptionKeyIdentifierRequest) Delete() (success bool, apiErr *common.Error) {
	dr := (*appmgrEncryption.DeleteEncryptionKeyIdentifierRequest)(deleteRequest)
	success, apiErr = dr.DeleteEncryptionKeyIdentifier()
	return
}
