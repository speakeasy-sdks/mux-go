package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetAssetInputInfoRequest struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type GetAssetInputInfoResponse struct {
	ContentType               string
	GetAssetInputInfoResponse *shared.GetAssetInputInfoResponse
	StatusCode                int
	RawResponse               *http.Response
}
