package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type GetVideoViewPathParams struct {
	VideoViewID string `pathParam:"style=simple,explode=false,name=VIDEO_VIEW_ID"`
}

type GetVideoViewRequest struct {
	PathParams GetVideoViewPathParams
	Retries    *utils.RetryConfig
}

type GetVideoViewResponse struct {
	ContentType       string
	StatusCode        int
	VideoViewResponse *shared.VideoViewResponse
}
