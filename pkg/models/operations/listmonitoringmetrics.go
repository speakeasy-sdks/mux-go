package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListMonitoringMetricsResponse struct {
	ContentType                   string
	ListMonitoringMetricsResponse *shared.ListMonitoringMetricsResponse
	StatusCode                    int
	RawResponse                   *http.Response
}
