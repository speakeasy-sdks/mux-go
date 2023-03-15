package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetVideoViewRequest struct {
	VideoViewID string `pathParam:"style=simple,explode=false,name=VIDEO_VIEW_ID"`
}

type GetVideoViewResponse struct {
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
	VideoViewResponse *shared.VideoViewResponse
}
