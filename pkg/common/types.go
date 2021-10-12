package common

import "fmt"

// PlatformURL URL of Platform API
var PlatformURL string

// Version of the Platform API to use
type Version string

// APIVersion is the version of the API to use
var APIVersion Version

// Versions all supported Platform API versions
type Versions []Version

// APIResponse is generic response type implemented by all endpoints (sync/async)
type APIResponse struct {
	// Status code from platform API
	StatusCode string `json:"status_code"`
	// Description from API
	Description string `json:"description"`
}

// AsyncResponse is received for all async responses
type AsyncResponse struct {
	// API response is part of async response
	APIResponse
	// Token is the token during async request
	RequestToken string `json:"request_token"`
}

// AsyncRequest is set of fields that should be included in all async requests
type AsyncRequest struct {
	// Callback url is URI of callback on app manager
	CallbackURL string `json:"callback_url"`
	// Unique token during making request
	RequestToken string `json:"request_token"`
}

// ResourceRequestRange specifies minimum and maximum resource limits
type ResourceRequestRange struct {
	// Lower limit of CPU usage as fraction of vCPU eg. `1.5`
	CPULowerLimit string `json:"cpu_lower_limit"`
	// Upper limit of CPU usage as fraction of vCPU eg. `2.5`
	CPUUpperLimit string `json:"cpu_upper_limit"`
	// Lower limit of memory usage eg. `0.5GB`
	MemoryLowerLimit string `json:"memory_lower_limit"`
	// Upper limit of memory usage eg. `1.4GB`
	MemoryUpperLimit string `json:"memory_upper_limit"`
}

// ResourceRequestLimit specifies resource limit
type ResourceRequestLimit struct {
	// Upper limit of CPU as fraction of vCPU eg. `2.5`
	CPUUpperLimit string `json:"cpu_upper_limit"`
	// Upper limit of memory eg. `2.5GB`
	MemoryUpperLimit string `json:"memory_upper_limit"`
}

// Protocol is the one of the primary layer-4 communication protocols.
type Protocol string

// This block defines the possible values for Protocol.
const (
	// TCP network protocol
	ProtocolTCP Protocol = "TCP"
	// UDP network protocol
	ProtocolUDP Protocol = "UDP"
	// SCTP network protcol
	ProtocolSCTP Protocol = "SCTP"
)

// Port is used to define a service port
type Port struct {
	// Name is the unique name for port of a given Group.
	Name string `json:"name"`
	// PortNumber is the port in container to be exposed by the port.
	PortNumber int32 `json:"portnumber"`
	// Protocol defines one of the network communication protocols to use for the
	// given port.
	Protocol `json:"protocol,omitempty"`
}

// Log identifies a log file
type Log struct {
	// Path is path to log file
	Path string `json:"path"`
	// Name is name
	Name string `json:"name"`
}

// UpdateStrategy is the strategy used to update workers in workergroups
type UpdateStrategy string

const (
	// RollingUpdate strategy is used to add only new workers after updations
	RollingUpdate UpdateStrategy = "RollingUpdate"
	// Recreate strategy will recreate all workers after updations
	Recreate UpdateStrategy = "Recreate"
)

// Health describes all health statuses
type Health string

const (
	// Healthy state
	Healthy Health = "Healthy"
	// Unhealthy state
	Unhealthy Health = "Unhealthy"
)

// Status desribes current state of an entity
type Status string

const (
	// Launching state
	Launching Status = "Launching"

	// LaunchFailed state
	LaunchFailed Status = "LaunchFailed"

	// Running state
	Running Status = "Running"

	// Starting indicates entity is in the process of starting
	Starting Status = "Starting"

	// Updating indicates entity is in the process of being updated
	Updating Status = "Updating"

	// Stopping indicates entity is in the process of stopping
	Stopping Status = "Stopping"

	// Stopped indicates entity is stopped
	Stopped Status = "Stopped"

	// Terminating indicates entity is in the process of terminating
	Terminating Status = "Terminating"

	// Terminated indicates entity is terminated
	Terminated Status = "Terminated"
)

// MetricsEndpoint defines the endpoint on the workers for fetching the metrics
type MetricsEndpoint struct {
	// MetricsPath is the path at which the metrics are exposed
	MetricsPath string `json:"metrics_path,omitempty"`
	// MetricsPort is the port number on the worker(main-container) on which metrics are exposed
	MetricsPort int `json:"metrics_port,omitempty"`
}

// EventType is used to describe the type of event
type EventType string

const (
	// MetricEvent are related to monitoring of a metrics expression. When the expression condition becomes
	// true, an event is emitted.
	MetricEvent EventType = "Metric"
	// LoggingEvent are related to monitoring of logs of an application. These events look for the
	// occurrence of a regex in any log files of an application and emits an event whenever it is found.
	LoggingEvent EventType = "Logging"
)

// VolumeType specifies type of the volume - dedicated or shared
type VolumeType string

// Block volume. Can only be attached to one volume at a time.
const BlockVolume VolumeType = "Block"

// FileSystem volume. Can only be attached to one volume at a time.
const FileSystemVolume VolumeType = "FileSystem"

type MetricsQueryMode string

const (
	// MetricsAggregateQuery specifies whether the metrics values are to be aggregated (summarized) or not.
	MetricsAggregate MetricsQueryMode = "FetchAggregate"
	// MetricsFetchLatestQuery sends the last collection value of the metrics.
	MetricsFetchLatest MetricsQueryMode = "FetchLatest"
)

type MetricsAggregationLevel string

const (
	// MetricsAggregateDeploymentLevel aggregates metrics at deployment level
	MetricsAggregateDeploymentLevel MetricsAggregationLevel = "Deployment"
	// MetricsAggregateWorkerGroupLevel aggregates metrics at worker group level
	MetricsAggregateWorkerGroupLevel MetricsAggregationLevel = "WorkerGroup"
	// MetricsAggregateWorkerGroupLevel aggregates metrics at worker level
	MetricsAggregateWorkerLevel MetricsAggregationLevel = "Worker"
)

type CredentialType string

const (
	// AlphaNumericType is used to generate AplhaNumeric credential.
	AlphaNumericType CredentialType = "AlphaNumeric"
	// AlphaNumericType is used to generate SHA256 credential.
	SHA256Type CredentialType = "SHA256"
	// AlphaNumericType is used to generate Hexadecimal credential.
	HexadecimalType CredentialType = "Hexadecimal"
	// AlphaNumericType is used to generate credential using the Regex format provided by user.
	RegexType CredentialType = "Regex"
)

// Error errors returned from API server
type Error struct {
	// Error code
	ErrorCode string
	// Error description
	ErrorDescription string
}

// Error is used to stringify the error into it's message
func (err Error) Error() string {
	return fmt.Sprintf("%s: %s", err.ErrorCode, err.ErrorDescription)
}
