package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type ListRealtimeMetricsRequest struct {
	Retries *utils.RetryConfig
}

type ListRealtimeMetricsResponse struct {
	ContentType                 string
	ListRealTimeMetricsResponse *shared.ListRealTimeMetricsResponse
	StatusCode                  int
}
