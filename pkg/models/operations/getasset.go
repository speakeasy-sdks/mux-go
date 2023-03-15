package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetAssetRequest struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type GetAssetResponse struct {
	AssetResponse *shared.AssetResponse
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}
