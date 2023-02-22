package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type ListRealtimeDimensionsRequest struct {
	Retries *utils.RetryConfig
}

type ListRealtimeDimensionsResponse struct {
	ContentType                    string
	ListRealTimeDimensionsResponse *shared.ListRealTimeDimensionsResponse
	StatusCode                     int
}
