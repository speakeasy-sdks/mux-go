package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type SignalLiveStreamCompleteRequest struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type SignalLiveStreamCompleteResponse struct {
	ContentType                      string
	SignalLiveStreamCompleteResponse *shared.SignalLiveStreamCompleteResponse
	StatusCode                       int
	RawResponse                      *http.Response
}
