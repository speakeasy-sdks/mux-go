package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type DeleteLiveStreamSimulcastTargetPathParams struct {
	LiveStreamID      string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	SimulcastTargetID string `pathParam:"style=simple,explode=false,name=SIMULCAST_TARGET_ID"`
}

type DeleteLiveStreamSimulcastTargetRequest struct {
	PathParams DeleteLiveStreamSimulcastTargetPathParams
	Retries    *utils.RetryConfig
}

type DeleteLiveStreamSimulcastTargetResponse struct {
	ContentType string
	StatusCode  int
}
