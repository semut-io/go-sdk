package metrics

import (
	"time"

	"go.semut.io/sdk/go-sdk/pkg/common"
)

// QueryRequest is the request used to query either user metrics or usage metrics from platform
type QueryRequest struct {
	// DeploymentID of the metrics to be queried, for worker this is only applicable if query mode is FetchAggregate
	DeploymentID string `json:"deployment_id,omitempty"`

	// WorkerGroupID of the metrics to be queried, for worker this is only applicable if query mode is FetchAggregate
	WorkerGroupID string `json:"worker_group_id,omitempty"`

	// WorkerID of the metrics to be queried, for worker this is only applicable if query mode is FetchAggregate
	WorkerID string `json:"worker_id,omitempty"`

	// Metrics is the list of metric names which are to be queried
	Metrics []string `json:"metrics"`

	// StartTime marks the starting time for getting the metrics,
	// empty value of StartTime, EndTime is applicable when using FetchAggregate, should
	// be empty when using FetchLatest
	StartTime time.Time `json:"start_time,omitempty"`

	// EndTime marks the ending time for getting the metrics,
	// empty value of StartTime, EndTime is applicable when using FetchAggregate, should
	// be empty when using FetchLatest
	EndTime time.Time `json:"end_time,omitempty"`

	// QueryMode specifies whether metrics query is to FetchLatest or FetchAggregate
	QueryMode common.MetricsQueryMode `json:"query_mode"`

	// MetricsAggregationLevel is used with FetchAggregate mode to aggregate metrics at
	// either Deployment, WorkerGroup or Worker level
	common.MetricsAggregationLevel `json:"aggregation_level,omitempty"`

	// AggregationInterval is used to determine the time duration per aggregate
	AggregationInterval time.Duration `json:"aggregation_interval,omitempty"`
}

// QueryResult is the result of querying the metrics, a list
// of these make up the response
type QueryResult struct {
	// DeploymentID of the metrics
	DeploymentID string `json:"deployment_id,omitempty"`
	// WorkerGroupID of the metrics
	WorkerGroupID string `json:"worker_group_id,omitempty"`
	// WorkerID of the metrics
	WorkerID string `json:"worker_id"`
	// MetricName is the name of the metric
	MetricName string `json:"metric_name"`
	// Timestamp when the metric was collected
	TimeStamp time.Time `json:"timestamp"`
	// Value is the value of the metric
	Value float64 `json:"value"`
	// Min is the value of aggregated minimum
	Min float64 `json:"min,omitempty"`
	// Max is the value of aggregated maximum
	Max float64 `json:"max,omitempty"`
	// Sum is the value of aggregated summation
	Sum float64 `json:"sum,omitempty"`
	// Avg is the value of aggregated average
	Avg float64 `json:"avg,omitempty"`
	// Median is the value of aggregated median
	Median float64 `json:"median,omitempty"`
}

// QueryResponse is the response to the metrics query
type QueryResponse struct {
	common.APIResponse
	// QueryResults return a list of metrics with their values and IDs
	QueryResults []QueryResult `json:"query_results,omitempty"`
}

// Query get metrics from the platform, metrics that can be queried can either be application collected metrics
// or the usage metrics from the platform like CPU, RAM usage
func (queryRequest *QueryRequest) Query() (queryResults []QueryResult, apiErr *common.Error) {

	queryResponse := QueryResponse{}
	err := common.Execute("MetricsQuery", queryRequest, &queryResponse)

	if err != nil {
		return nil, err
	}

	if queryResponse.StatusCode != "200" {

		apiErr = &common.Error{
			ErrorCode:        queryResponse.StatusCode,
			ErrorDescription: queryResponse.Description,
		}

		return nil, apiErr
	}

	return queryResponse.QueryResults, nil
}
