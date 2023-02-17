package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type UpdateAssetMp4SupportPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type UpdateAssetMp4SupportRequest struct {
	PathParams UpdateAssetMp4SupportPathParams
	Request    shared.UpdateAssetMp4SupportRequest `request:"mediaType=application/json"`
}

type UpdateAssetMp4SupportResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
}
