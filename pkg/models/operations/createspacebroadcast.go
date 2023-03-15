package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateSpaceBroadcastRequest struct {
	CreateBroadcastRequest shared.CreateBroadcastRequest `request:"mediaType=application/json"`
	SpaceID                string                        `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type CreateSpaceBroadcastResponse struct {
	BroadcastResponse *shared.BroadcastResponse
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
}
