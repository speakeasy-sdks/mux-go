package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type ListMonitoringDimensionsRequest struct {
	Retries *utils.RetryConfig
}

type ListMonitoringDimensionsResponse struct {
	ContentType                      string
	ListMonitoringDimensionsResponse *shared.ListMonitoringDimensionsResponse
	StatusCode                       int
}
