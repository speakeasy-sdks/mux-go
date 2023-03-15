package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateLiveStreamSimulcastTargetRequest struct {
	CreateSimulcastTargetRequest shared.CreateSimulcastTargetRequest `request:"mediaType=application/json"`
	LiveStreamID                 string                              `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type CreateLiveStreamSimulcastTargetResponse struct {
	ContentType             string
	SimulcastTargetResponse *shared.SimulcastTargetResponse
	StatusCode              int
	RawResponse             *http.Response
}
