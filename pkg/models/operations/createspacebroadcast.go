package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type CreateSpaceBroadcastPathParams struct {
	SpaceID string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type CreateSpaceBroadcastRequest struct {
	PathParams CreateSpaceBroadcastPathParams
	Request    shared.CreateBroadcastRequest `request:"mediaType=application/json"`
	Retries    *utils.RetryConfig
}

type CreateSpaceBroadcastResponse struct {
	BroadcastResponse *shared.BroadcastResponse
	ContentType       string
	StatusCode        int
}
