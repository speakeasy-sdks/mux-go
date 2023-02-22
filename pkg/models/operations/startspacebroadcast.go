package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type StartSpaceBroadcastPathParams struct {
	BroadcastID string `pathParam:"style=simple,explode=false,name=BROADCAST_ID"`
	SpaceID     string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type StartSpaceBroadcastRequest struct {
	PathParams StartSpaceBroadcastPathParams
	Retries    *utils.RetryConfig
}

type StartSpaceBroadcastResponse struct {
	ContentType                 string
	StartSpaceBroadcastResponse *shared.StartSpaceBroadcastResponse
	StatusCode                  int
}
