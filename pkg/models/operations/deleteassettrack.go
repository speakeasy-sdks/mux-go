package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type DeleteAssetTrackPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
	TrackID string `pathParam:"style=simple,explode=false,name=TRACK_ID"`
}

type DeleteAssetTrackRequest struct {
	PathParams DeleteAssetTrackPathParams
	Retries    *utils.RetryConfig
}

type DeleteAssetTrackResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
