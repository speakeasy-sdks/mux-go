package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type UpdateLiveStreamRequest struct {
	LiveStreamID            string                         `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	UpdateLiveStreamRequest shared.UpdateLiveStreamRequest `request:"mediaType=application/json"`
}

type UpdateLiveStreamResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
	RawResponse        *http.Response
}
