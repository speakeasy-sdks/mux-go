package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type ListMonitoringMetricsRequest struct {
	Retries *utils.RetryConfig
}

type ListMonitoringMetricsResponse struct {
	ContentType                   string
	ListMonitoringMetricsResponse *shared.ListMonitoringMetricsResponse
	StatusCode                    int
}
