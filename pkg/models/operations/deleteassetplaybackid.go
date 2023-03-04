package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type DeleteAssetPlaybackIDPathParams struct {
	AssetID    string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
	PlaybackID string `pathParam:"style=simple,explode=false,name=PLAYBACK_ID"`
}

type DeleteAssetPlaybackIDRequest struct {
	PathParams DeleteAssetPlaybackIDPathParams
	Retries    *utils.RetryConfig
}

type DeleteAssetPlaybackIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
