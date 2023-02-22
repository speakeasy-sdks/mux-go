package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type UpdateAssetMp4SupportPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type UpdateAssetMp4SupportRequest struct {
	PathParams UpdateAssetMp4SupportPathParams
	Request    shared.UpdateAssetMp4SupportRequest `request:"mediaType=application/json"`
	Retries    *utils.RetryConfig
}

type UpdateAssetMp4SupportResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
}
