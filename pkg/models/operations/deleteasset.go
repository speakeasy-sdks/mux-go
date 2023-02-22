package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type DeleteAssetPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type DeleteAssetRequest struct {
	PathParams DeleteAssetPathParams
	Retries    *utils.RetryConfig
}

type DeleteAssetResponse struct {
	ContentType string
	StatusCode  int
}
