package operations

import (
	"net/http"
)

type DeleteLiveStreamPlaybackIDPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	PlaybackID   string `pathParam:"style=simple,explode=false,name=PLAYBACK_ID"`
}

type DeleteLiveStreamPlaybackIDRequest struct {
	PathParams DeleteLiveStreamPlaybackIDPathParams
}

type DeleteLiveStreamPlaybackIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
