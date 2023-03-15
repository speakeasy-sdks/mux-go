package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type UpdateAssetRequest struct {
	AssetID            string                    `pathParam:"style=simple,explode=false,name=ASSET_ID"`
	UpdateAssetRequest shared.UpdateAssetRequest `request:"mediaType=application/json"`
}

type UpdateAssetResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}
