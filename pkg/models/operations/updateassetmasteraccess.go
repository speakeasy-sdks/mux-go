package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type UpdateAssetMasterAccessRequest struct {
	AssetID                        string                                `pathParam:"style=simple,explode=false,name=ASSET_ID"`
	UpdateAssetMasterAccessRequest shared.UpdateAssetMasterAccessRequest `request:"mediaType=application/json"`
}

type UpdateAssetMasterAccessResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}
