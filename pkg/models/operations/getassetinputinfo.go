package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetAssetInputInfoPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type GetAssetInputInfoRequest struct {
	PathParams GetAssetInputInfoPathParams
}

type GetAssetInputInfoResponse struct {
	ContentType               string
	GetAssetInputInfoResponse *shared.GetAssetInputInfoResponse
	StatusCode                int
}
