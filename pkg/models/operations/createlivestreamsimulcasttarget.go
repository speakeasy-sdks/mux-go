package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type CreateLiveStreamSimulcastTargetPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type CreateLiveStreamSimulcastTargetRequest struct {
	PathParams CreateLiveStreamSimulcastTargetPathParams
	Request    shared.CreateSimulcastTargetRequest `request:"mediaType=application/json"`
	Retries    *utils.RetryConfig
}

type CreateLiveStreamSimulcastTargetResponse struct {
	ContentType             string
	SimulcastTargetResponse *shared.SimulcastTargetResponse
	StatusCode              int
}
