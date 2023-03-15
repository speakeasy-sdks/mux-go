package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateLiveStreamPlaybackIDRequest struct {
	CreatePlaybackIDRequest shared.CreatePlaybackIDRequest `request:"mediaType=application/json"`
	LiveStreamID            string                         `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type CreateLiveStreamPlaybackIDResponse struct {
	ContentType              string
	CreatePlaybackIDResponse *shared.CreatePlaybackIDResponse
	StatusCode               int
	RawResponse              *http.Response
}
