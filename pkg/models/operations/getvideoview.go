package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetVideoViewPathParams struct {
	VideoViewID string `pathParam:"style=simple,explode=false,name=VIDEO_VIEW_ID"`
}

type GetVideoViewRequest struct {
	PathParams GetVideoViewPathParams
}

type GetVideoViewResponse struct {
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
	VideoViewResponse *shared.VideoViewResponse
}
