package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetLiveStreamPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type GetLiveStreamRequest struct {
	PathParams GetLiveStreamPathParams
}

type GetLiveStreamResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
}
