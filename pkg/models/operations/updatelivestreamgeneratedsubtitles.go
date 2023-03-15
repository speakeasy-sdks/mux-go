package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type UpdateLiveStreamGeneratedSubtitlesRequest struct {
	LiveStreamID                              string                                           `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	UpdateLiveStreamGeneratedSubtitlesRequest shared.UpdateLiveStreamGeneratedSubtitlesRequest `request:"mediaType=application/json"`
}

type UpdateLiveStreamGeneratedSubtitlesResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
	RawResponse        *http.Response
}
