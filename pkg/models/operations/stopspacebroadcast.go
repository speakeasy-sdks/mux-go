package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type StopSpaceBroadcastPathParams struct {
	BroadcastID string `pathParam:"style=simple,explode=false,name=BROADCAST_ID"`
	SpaceID     string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type StopSpaceBroadcastRequest struct {
	PathParams StopSpaceBroadcastPathParams
	Retries    *utils.RetryConfig
}

type StopSpaceBroadcastResponse struct {
	ContentType                string
	StatusCode                 int
	RawResponse                *http.Response
	StopSpaceBroadcastResponse *shared.StopSpaceBroadcastResponse
}
