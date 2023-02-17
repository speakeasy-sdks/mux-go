package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type UpdateAssetPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type UpdateAssetRequest struct {
	PathParams UpdateAssetPathParams
	Request    shared.UpdateAssetRequest `request:"mediaType=application/json"`
}

type UpdateAssetResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
}
