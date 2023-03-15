package operations

import (
	"net/http"
)

type DeleteAssetRequest struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type DeleteAssetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
