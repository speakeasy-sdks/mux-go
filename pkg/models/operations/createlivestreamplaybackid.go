package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type CreateLiveStreamPlaybackIDPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type CreateLiveStreamPlaybackIDRequest struct {
	PathParams CreateLiveStreamPlaybackIDPathParams
	Request    shared.CreatePlaybackIDRequest `request:"mediaType=application/json"`
	Retries    *utils.RetryConfig
}

type CreateLiveStreamPlaybackIDResponse struct {
	ContentType              string
	CreatePlaybackIDResponse *shared.CreatePlaybackIDResponse
	StatusCode               int
}
