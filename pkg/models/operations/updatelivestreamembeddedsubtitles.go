package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type UpdateLiveStreamEmbeddedSubtitlesPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type UpdateLiveStreamEmbeddedSubtitlesRequest struct {
	PathParams UpdateLiveStreamEmbeddedSubtitlesPathParams
	Request    shared.UpdateLiveStreamEmbeddedSubtitlesRequest `request:"mediaType=application/json"`
}

type UpdateLiveStreamEmbeddedSubtitlesResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
}
