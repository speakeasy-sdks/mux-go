package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type CreateLiveStreamPlaybackIDPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type CreateLiveStreamPlaybackIDRequest struct {
	PathParams CreateLiveStreamPlaybackIDPathParams
	Request    shared.CreatePlaybackIDRequest `request:"mediaType=application/json"`
}

type CreateLiveStreamPlaybackIDResponse struct {
	ContentType              string
	CreatePlaybackIDResponse *shared.CreatePlaybackIDResponse
	StatusCode               int
}
