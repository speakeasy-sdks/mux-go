package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetSpaceRequest struct {
	SpaceID string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type GetSpaceResponse struct {
	ContentType   string
	SpaceResponse *shared.SpaceResponse
	StatusCode    int
	RawResponse   *http.Response
}
