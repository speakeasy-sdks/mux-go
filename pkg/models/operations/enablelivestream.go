package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type EnableLiveStreamPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type EnableLiveStreamRequest struct {
	PathParams EnableLiveStreamPathParams
	Retries    *utils.RetryConfig
}

type EnableLiveStreamResponse struct {
	ContentType              string
	EnableLiveStreamResponse *shared.EnableLiveStreamResponse
	StatusCode               int
	RawResponse              *http.Response
}
