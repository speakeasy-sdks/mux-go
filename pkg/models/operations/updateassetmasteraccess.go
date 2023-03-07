package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type UpdateAssetMasterAccessPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type UpdateAssetMasterAccessRequest struct {
	PathParams UpdateAssetMasterAccessPathParams
	Request    shared.UpdateAssetMasterAccessRequest `request:"mediaType=application/json"`
}

type UpdateAssetMasterAccessResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}
