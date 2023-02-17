package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type ListRealtimeMetricsResponse struct {
	ContentType                 string
	ListRealTimeMetricsResponse *shared.ListRealTimeMetricsResponse
	StatusCode                  int
}
