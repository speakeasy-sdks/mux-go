package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type UpdateAssetMasterAccessPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type UpdateAssetMasterAccessRequest struct {
	PathParams UpdateAssetMasterAccessPathParams
	Request    shared.UpdateAssetMasterAccessRequest `request:"mediaType=application/json"`
	Retries    *utils.RetryConfig
}

type UpdateAssetMasterAccessResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
}
