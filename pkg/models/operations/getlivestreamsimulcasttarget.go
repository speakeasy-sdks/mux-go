package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetLiveStreamSimulcastTargetRequest struct {
	LiveStreamID      string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	SimulcastTargetID string `pathParam:"style=simple,explode=false,name=SIMULCAST_TARGET_ID"`
}

type GetLiveStreamSimulcastTargetResponse struct {
	ContentType             string
	SimulcastTargetResponse *shared.SimulcastTargetResponse
	StatusCode              int
	RawResponse             *http.Response
}
