package operations

import (
	"net/http"
)

type DeleteAssetPlaybackIDRequest struct {
	AssetID    string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
	PlaybackID string `pathParam:"style=simple,explode=false,name=PLAYBACK_ID"`
}

type DeleteAssetPlaybackIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
