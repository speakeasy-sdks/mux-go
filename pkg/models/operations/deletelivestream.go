package operations

import (
	"net/http"
)

type DeleteLiveStreamRequest struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type DeleteLiveStreamResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
