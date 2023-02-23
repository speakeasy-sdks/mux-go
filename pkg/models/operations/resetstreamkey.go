package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type ResetStreamKeyPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type ResetStreamKeyRequest struct {
	PathParams ResetStreamKeyPathParams
	Retries    *utils.RetryConfig
}

type ResetStreamKeyResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
}
