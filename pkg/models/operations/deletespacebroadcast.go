package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type DeleteSpaceBroadcastPathParams struct {
	BroadcastID string `pathParam:"style=simple,explode=false,name=BROADCAST_ID"`
	SpaceID     string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type DeleteSpaceBroadcastRequest struct {
	PathParams DeleteSpaceBroadcastPathParams
	Retries    *utils.RetryConfig
}

type DeleteSpaceBroadcastResponse struct {
	ContentType string
	StatusCode  int
}
