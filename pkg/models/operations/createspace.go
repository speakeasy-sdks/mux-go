package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type CreateSpaceRequest struct {
	Request shared.CreateSpaceRequest `request:"mediaType=application/json"`
}

type CreateSpaceResponse struct {
	ContentType   string
	SpaceResponse *shared.SpaceResponse
	StatusCode    int
}
