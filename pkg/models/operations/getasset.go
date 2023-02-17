package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetAssetPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type GetAssetRequest struct {
	PathParams GetAssetPathParams
}

type GetAssetResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
}
