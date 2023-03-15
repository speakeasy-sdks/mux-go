package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type EnableLiveStreamRequest struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type EnableLiveStreamResponse struct {
	ContentType              string
	EnableLiveStreamResponse *shared.EnableLiveStreamResponse
	StatusCode               int
	RawResponse              *http.Response
}
