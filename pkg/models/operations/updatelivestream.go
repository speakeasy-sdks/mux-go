package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type UpdateLiveStreamPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type UpdateLiveStreamRequest struct {
	PathParams UpdateLiveStreamPathParams
	Request    shared.UpdateLiveStreamRequest `request:"mediaType=application/json"`
	Retries    *utils.RetryConfig
}

type UpdateLiveStreamResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
}
