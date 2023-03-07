package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetLiveStreamPlaybackIDPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	PlaybackID   string `pathParam:"style=simple,explode=false,name=PLAYBACK_ID"`
}

type GetLiveStreamPlaybackIDRequest struct {
	PathParams GetLiveStreamPlaybackIDPathParams
}

type GetLiveStreamPlaybackIDResponse struct {
	ContentType                     string
	GetLiveStreamPlaybackIDResponse *shared.GetLiveStreamPlaybackIDResponse
	StatusCode                      int
	RawResponse                     *http.Response
}
