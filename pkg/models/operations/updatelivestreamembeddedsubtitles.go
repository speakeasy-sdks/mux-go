package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type UpdateLiveStreamEmbeddedSubtitlesRequest struct {
	LiveStreamID                             string                                          `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	UpdateLiveStreamEmbeddedSubtitlesRequest shared.UpdateLiveStreamEmbeddedSubtitlesRequest `request:"mediaType=application/json"`
}

type UpdateLiveStreamEmbeddedSubtitlesResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
	RawResponse        *http.Response
}
