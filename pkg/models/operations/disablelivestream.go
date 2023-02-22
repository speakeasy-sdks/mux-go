package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type DisableLiveStreamPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type DisableLiveStreamRequest struct {
	PathParams DisableLiveStreamPathParams
	Retries    *utils.RetryConfig
}

type DisableLiveStreamResponse struct {
	ContentType               string
	DisableLiveStreamResponse *shared.DisableLiveStreamResponse
	StatusCode                int
}
