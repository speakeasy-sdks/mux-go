package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type DisableLiveStreamRequest struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type DisableLiveStreamResponse struct {
	ContentType               string
	DisableLiveStreamResponse *shared.DisableLiveStreamResponse
	StatusCode                int
	RawResponse               *http.Response
}
