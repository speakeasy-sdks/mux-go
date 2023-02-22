package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type UpdateLiveStreamGeneratedSubtitlesPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type UpdateLiveStreamGeneratedSubtitlesRequest struct {
	PathParams UpdateLiveStreamGeneratedSubtitlesPathParams
	Request    shared.UpdateLiveStreamGeneratedSubtitlesRequest `request:"mediaType=application/json"`
	Retries    *utils.RetryConfig
}

type UpdateLiveStreamGeneratedSubtitlesResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
}
