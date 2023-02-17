package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type DeleteSpacePathParams struct {
	SpaceID string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type DeleteSpaceRequest struct {
	PathParams DeleteSpacePathParams
	Retries    *utils.RetryConfig
}

type DeleteSpaceResponse struct {
	ContentType string
	StatusCode  int
}
