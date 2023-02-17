package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type CreateAssetRequest struct {
	Request shared.CreateAssetRequest `request:"mediaType=application/json"`
}

type CreateAssetResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
}
