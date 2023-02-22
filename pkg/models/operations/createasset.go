package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type CreateAssetRequest struct {
	Request shared.CreateAssetRequest `request:"mediaType=application/json"`
	Retries *utils.RetryConfig
}

type CreateAssetResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
}
