package images

import (
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// ImageSpec contains all the details needed to import a new image on the platform
type ImageSpec struct {
	// Registry is the URL to the registry from where the image is to be pulled. It is left empty if the
	// image is to be pulled from DockerHub (https:// hub.docker.com/). Examples include quay.io, gcr.io,
	// <AWS_ACCOUNT_ID>.dkr.ecr.(AWS_REGION).amazonaws.com and so on.
	RegistryHost string `json:"registry,omitempty"`
	// RegistryPort is the port of the specific registry. As container registries generally work
	// with https, so if it is left empty 443 is used by default.
	RegistryPort int `json:"registry_port,omitempty"`
	// Name of the image.
	Name string `json:"name"`
	// Tag of the image. `latest` tag is used if none is provided.
	Tag string `json:"tag,omitempty"`
}

// Image is the description of an Image
type Image struct {
	// ImageURI is the unique UUID of the Image
	ImageURI string `json:"image_uri"`
	// ImageSpec is the specification
	ImageSpec
}

// NewImageRequest is the spec for adding a new image to the Platform repository
type NewImageRequest struct {
	ImageSpec
	// Auth represents a request token to be used for pulling the image from a private registy,
	// it is essentially the base64 encoded string containing `username:token` of the registry,
	// this is not required if the image is public
	Auth string `json:"auth,omitempty"`
}

// NewImageResponse is the response to the new image request
type NewImageResponse struct {
	common.APIResponse
	// ID of the imported image
	ImageURI string `json:"image_uri,omitempty"`
}

// DeleteRequest is used to request existing image deletion
type DeleteRequest struct {
	// ID of the existing image
	ImageURI string `json:"image_uri,omitempty"`
}

// DeleteResponse is the response of the delete request
type DeleteResponse struct {
	common.APIResponse
}

// DescribeRequest is used to request the details of an existing image
type DescribeRequest struct {
	// ID of the existing image
	ImageURI string `json:"image_uri,omitempty"`
}

// DescribeResponse is the response to the get request
type DescribeResponse struct {
	common.APIResponse
	// Details of the image
	Image
}

// ListRequest is to request a listing of all the images
type ListRequest struct{}

// ListResponse is the response to the list request
type ListResponse struct {
	Images []Image `json:"images,omitempty"`
	common.APIResponse
}

// New image request is used to import an image into the platform from a specific registry
func (req *NewImageRequest) New() (imageURI string, apiErr *common.Error) {
	res := NewImageResponse{}
	err := common.Execute("ImagesNew", req, &res)

	if err != nil {
		return "", &common.ErrInvalidResponseFromAPI
	}
	if res.StatusCode != "200" {
		apiErr = &common.Error{
			ErrorCode:        res.StatusCode,
			ErrorDescription: res.Description,
		}
		return "", apiErr
	}
	return res.ImageURI, nil
}

// Describe image request is used to get the details of an existing image
func (req *DescribeRequest) Describe() (image Image, apiErr *common.Error) {
	res := DescribeResponse{}
	err := common.Execute("ImagesGet", req, &res)

	if err != nil {
		return image, &common.ErrInvalidResponseFromAPI
	}
	if res.StatusCode != "200" {
		apiErr = &common.Error{
			ErrorCode:        res.StatusCode,
			ErrorDescription: res.Description,
		}
		return image, apiErr
	}
	return res.Image, nil
}

// Delete image is used to delete an existing image
func (req *DeleteRequest) Delete() (success bool, apiErr *common.Error) {
	res := DeleteResponse{}
	err := common.Execute("ImagesDelete", req, &res)

	if err != nil {
		return false, &common.ErrInvalidResponseFromAPI
	}
	if res.StatusCode != "200" {
		apiErr = &common.Error{
			ErrorCode:        res.StatusCode,
			ErrorDescription: res.Description,
		}
		return false, apiErr
	}
	return true, nil
}

// List image is used to get a listing of all the existing images
func (req *ListRequest) List() (images []Image, apiErr *common.Error) {
	res := ListResponse{}
	err := common.Execute("ImagesList", req, &res)

	if err != nil {
		return images, &common.ErrInvalidResponseFromAPI
	}
	if res.StatusCode != "200" {
		apiErr = &common.Error{
			ErrorCode:        res.StatusCode,
			ErrorDescription: res.Description,
		}
		return images, apiErr
	}
	return res.Images, nil
}
