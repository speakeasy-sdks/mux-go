package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type UpdateAssetMp4SupportRequest struct {
	AssetID                      string                              `pathParam:"style=simple,explode=false,name=ASSET_ID"`
	UpdateAssetMp4SupportRequest shared.UpdateAssetMp4SupportRequest `request:"mediaType=application/json"`
}

type UpdateAssetMp4SupportResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}
