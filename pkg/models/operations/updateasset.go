package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type UpdateAssetPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type UpdateAssetRequest struct {
	PathParams UpdateAssetPathParams
	Request    shared.UpdateAssetRequest `request:"mediaType=application/json"`
	Retries    *utils.RetryConfig
}

type UpdateAssetResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
}
