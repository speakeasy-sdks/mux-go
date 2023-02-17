package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type DeleteLiveStreamPlaybackIDPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	PlaybackID   string `pathParam:"style=simple,explode=false,name=PLAYBACK_ID"`
}

type DeleteLiveStreamPlaybackIDRequest struct {
	PathParams DeleteLiveStreamPlaybackIDPathParams
	Retries    *utils.RetryConfig
}

type DeleteLiveStreamPlaybackIDResponse struct {
	ContentType string
	StatusCode  int
}
