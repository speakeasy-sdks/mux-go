package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type ListRealtimeDimensionsRequest struct {
	Retries *utils.RetryConfig
}

type ListRealtimeDimensionsResponse struct {
	ContentType                    string
	ListRealTimeDimensionsResponse *shared.ListRealTimeDimensionsResponse
	StatusCode                     int
	RawResponse                    *http.Response
}
