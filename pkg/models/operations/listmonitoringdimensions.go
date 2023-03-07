package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListMonitoringDimensionsResponse struct {
	ContentType                      string
	ListMonitoringDimensionsResponse *shared.ListMonitoringDimensionsResponse
	StatusCode                       int
	RawResponse                      *http.Response
}
