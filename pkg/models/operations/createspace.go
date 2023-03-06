package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type CreateSpaceRequest struct {
	Request shared.CreateSpaceRequest `request:"mediaType=application/json"`
	Retries *utils.RetryConfig
}

type CreateSpaceResponse struct {
	ContentType   string
	SpaceResponse *shared.SpaceResponse
	StatusCode    int
	RawResponse   *http.Response
}
