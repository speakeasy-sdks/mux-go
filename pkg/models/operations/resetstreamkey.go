package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ResetStreamKeyPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type ResetStreamKeyRequest struct {
	PathParams ResetStreamKeyPathParams
}

type ResetStreamKeyResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
	RawResponse        *http.Response
}
