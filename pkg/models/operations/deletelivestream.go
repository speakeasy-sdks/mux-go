package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type DeleteLiveStreamPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type DeleteLiveStreamRequest struct {
	PathParams DeleteLiveStreamPathParams
	Retries    *utils.RetryConfig
}

type DeleteLiveStreamResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
