package metrics

import (
	appmgrMetrics "go.semut.io/sdk/go-sdk/pkg/appmanager/metrics"
	"go.semut.io/sdk/go-sdk/pkg/common"
)

// QueryRequest is the request used to query either user metrics or usage metrics from platform
type QueryRequest appmgrMetrics.QueryRequest

// QueryResult is the result generated from metrics query
type QueryResults []appmgrMetrics.QueryResult

// Query is an alias of app manager query request
func (queryRequest *QueryRequest) Query() (QueryResults, *common.Error) {
	qr := (*appmgrMetrics.QueryRequest)(queryRequest)
	results, apiErr := qr.Query()
	return QueryResults(results), apiErr
}
