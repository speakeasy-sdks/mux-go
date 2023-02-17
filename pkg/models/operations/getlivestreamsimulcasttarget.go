package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetLiveStreamSimulcastTargetPathParams struct {
	LiveStreamID      string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	SimulcastTargetID string `pathParam:"style=simple,explode=false,name=SIMULCAST_TARGET_ID"`
}

type GetLiveStreamSimulcastTargetRequest struct {
	PathParams GetLiveStreamSimulcastTargetPathParams
}

type GetLiveStreamSimulcastTargetResponse struct {
	ContentType             string
	SimulcastTargetResponse *shared.SimulcastTargetResponse
	StatusCode              int
}
