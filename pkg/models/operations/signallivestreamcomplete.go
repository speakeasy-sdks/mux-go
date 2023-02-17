package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type SignalLiveStreamCompletePathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type SignalLiveStreamCompleteRequest struct {
	PathParams SignalLiveStreamCompletePathParams
}

type SignalLiveStreamCompleteResponse struct {
	ContentType                      string
	SignalLiveStreamCompleteResponse *shared.SignalLiveStreamCompleteResponse
	StatusCode                       int
}
