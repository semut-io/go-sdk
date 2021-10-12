package logexport

import (
	"time"

	"go.semut.io/sdk/go-sdk/pkg/common"
)

type LogExportSpec struct {
	// DeploymentID for this request
	DeploymentID string `json:"deployment_id"`
	// WorkerGroupID for this request
	WorkerGroupID string `json:"worker_group_id,omitempty"`
	// WorkerID for this request
	WorkerID string `json:"worker_id,omitempty"`
	// Filepath is the path for the log file for which log rotation and upload is being enabled
	Filepath string `json:"filepath"`
	// MaxFileSize is the threshold for maximum size of the log file, past which the rotation process
	// is rotated
	MaxFileSize int64 `json:"max_file_size,omitempty"`
	// MaxCreatedBefore is the threshold for file creation time past which the rotation process is started.
	MaxCreatedBefore time.Duration `json:"max_created_before,omitempty"`
	// RotationFormat is the format used to name the rotated file. Should contain a timestamp format.
	// Only to be provided if log rotation is implemented at the application level itself so that the
	// service can upload the rotated files to the object store
	RotationFormat string `json:"rotation_format,omitempty"`
	// FileCount is the number of log files that are kept on a pod at a time
	FileCount int `json:"file_count,omitempty"`
	// Upload denotes whether the rotated file needs to be uploaded or not
	Upload bool `json:"upload,omitempty"`
}

// LogExport contains export id and specification details
type LogExport struct {
	// LogExportID is the unique UUID of the LogExport.
	LogExportID string `json:"log_export_id"`
	// Specification of log export
	LogExportSpec
}

// LogExportEnable is the request used to enable exporting of logs from either workers or worker groups
type EnableLogExportRequest struct {
	LogExportSpec
}

// LogExportDisable is used to disable an already existing log exporter
type DisableLogExportRequest struct {
	// DeploymentID for this request
	DeploymentID string `json:"deployment_id"`
	// WorkerGroupID for this request
	WorkerGroupID string `json:"worker_group_id,omitempty"`
	// WorkerID for this request
	WorkerID string `json:"worker_id,omitempty"`
	// File path of the log
	Filepath string `json:"filepath"`
}

// LogExportEnable is the response of the LogExport Enable endpoint
type EnableLogExportResponse struct {
	common.APIResponse
}

// LogExportEnable is the response for the LogExport Disable endpoint
type DisableLogExportResponse struct {
	common.APIResponse
}

// ListLogExportRequest to used to request listing of existing log exporters
type ListLogExportRequest struct {
	// DeploymentID for this request
	DeploymentID string `json:"deployment_id"`
	// WorkerGroupID for this request
	WorkerGroupID string `json:"worker_group_id,omitempty"`
	// WorkerID for this request
	WorkerID string `json:"worker_id,omitempty"`
}

// ListLogExportResponse is response of the list request
type ListLogExportResponse struct {
	common.APIResponse
	// List of log exporters
	LogExports []LogExport `json:"log_exports,omitempty"`
}

// UpdateLogExportRequest is used to update the spec of an existing Log Export
type UpdateLogExportRequest struct {
	LogExportSpec
}

// UpdateLogExportResponse is response to the request
type UpdateLogExportResponse struct {
	common.APIResponse
}

// Enable is used to enable the log exporter to run on worker / all workers in worker group / whole deployment
func (req *EnableLogExportRequest) Enable() (apiErr *common.Error) {
	res := EnableLogExportResponse{}
	err := common.Execute("LogExportEnable", req, &res)

	if err != nil {
		return &common.ErrInvalidResponseFromAPI
	}
	if res.StatusCode != "200" {
		apiErr = &common.Error{
			ErrorCode:        res.StatusCode,
			ErrorDescription: res.Description,
		}
		return apiErr
	}
	return nil
}

// Disable to used to disable an existing log exporter
func (req *DisableLogExportRequest) Disable() (apiErr *common.Error) {
	res := DisableLogExportResponse{}
	err := common.Execute("LogExportDisable", req, &res)

	if err != nil {
		return &common.ErrInvalidResponseFromAPI
	}
	if res.StatusCode != "200" {
		apiErr = &common.Error{
			ErrorCode:        res.StatusCode,
			ErrorDescription: res.Description,
		}
		return apiErr
	}
	return nil
}

// List to used to get list of existing log exporters
func (req *ListLogExportRequest) List() (logExports []LogExport, apiErr *common.Error) {
	res := ListLogExportResponse{}
	err := common.Execute("LogExportList", req, &res)

	if err != nil {
		return nil, &common.ErrInvalidResponseFromAPI
	}
	if res.StatusCode != "200" {
		apiErr = &common.Error{
			ErrorCode:        res.StatusCode,
			ErrorDescription: res.Description,
		}
		return nil, apiErr
	}
	return res.LogExports, nil
}

// Update is used to update existing log export
func (req *UpdateLogExportRequest) Update() (apiErr *common.Error) {
	res := UpdateLogExportResponse{}
	err := common.Execute("LogExportUpdate", req, &res)

	if err != nil {
		return &common.ErrInvalidResponseFromAPI
	}
	if res.StatusCode != "200" {
		apiErr = &common.Error{
			ErrorCode:        res.StatusCode,
			ErrorDescription: res.Description,
		}
		return apiErr
	}
	return nil
}
