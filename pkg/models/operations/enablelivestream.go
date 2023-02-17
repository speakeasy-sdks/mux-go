package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type EnableLiveStreamPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type EnableLiveStreamRequest struct {
	PathParams EnableLiveStreamPathParams
}

type EnableLiveStreamResponse struct {
	ContentType              string
	EnableLiveStreamResponse *shared.EnableLiveStreamResponse
	StatusCode               int
}
