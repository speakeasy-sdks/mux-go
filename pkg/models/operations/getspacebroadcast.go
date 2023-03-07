package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetSpaceBroadcastPathParams struct {
	BroadcastID string `pathParam:"style=simple,explode=false,name=BROADCAST_ID"`
	SpaceID     string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type GetSpaceBroadcastRequest struct {
	PathParams GetSpaceBroadcastPathParams
}

type GetSpaceBroadcastResponse struct {
	BroadcastResponse *shared.BroadcastResponse
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
}
