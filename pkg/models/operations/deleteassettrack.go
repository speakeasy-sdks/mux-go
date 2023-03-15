package operations

import (
	"net/http"
)

type DeleteAssetTrackRequest struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
	TrackID string `pathParam:"style=simple,explode=false,name=TRACK_ID"`
}

type DeleteAssetTrackResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
