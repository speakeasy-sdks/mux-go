package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type UpdateLiveStreamPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type UpdateLiveStreamRequest struct {
	PathParams UpdateLiveStreamPathParams
	Request    shared.UpdateLiveStreamRequest `request:"mediaType=application/json"`
}

type UpdateLiveStreamResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
	RawResponse        *http.Response
}
