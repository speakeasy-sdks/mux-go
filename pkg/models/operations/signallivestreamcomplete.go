package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type SignalLiveStreamCompletePathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type SignalLiveStreamCompleteRequest struct {
	PathParams SignalLiveStreamCompletePathParams
	Retries    *utils.RetryConfig
}

type SignalLiveStreamCompleteResponse struct {
	ContentType                      string
	SignalLiveStreamCompleteResponse *shared.SignalLiveStreamCompleteResponse
	StatusCode                       int
	RawResponse                      *http.Response
}
