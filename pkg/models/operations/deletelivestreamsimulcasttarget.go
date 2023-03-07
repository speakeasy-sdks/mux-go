package operations

import (
	"net/http"
)

type DeleteLiveStreamSimulcastTargetPathParams struct {
	LiveStreamID      string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
	SimulcastTargetID string `pathParam:"style=simple,explode=false,name=SIMULCAST_TARGET_ID"`
}

type DeleteLiveStreamSimulcastTargetRequest struct {
	PathParams DeleteLiveStreamSimulcastTargetPathParams
}

type DeleteLiveStreamSimulcastTargetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
