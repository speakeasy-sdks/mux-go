package operations

import (
	"net/http"
)

type DeleteAssetPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type DeleteAssetRequest struct {
	PathParams DeleteAssetPathParams
}

type DeleteAssetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
