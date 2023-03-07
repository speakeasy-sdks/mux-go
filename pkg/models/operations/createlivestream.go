package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateLiveStreamRequest struct {
	Request shared.CreateLiveStreamRequest `request:"mediaType=application/json"`
}

type CreateLiveStreamResponse struct {
	ContentType        string
	LiveStreamResponse *shared.LiveStreamResponse
	StatusCode         int
	RawResponse        *http.Response
}
