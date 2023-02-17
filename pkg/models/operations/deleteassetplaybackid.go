package operations

type DeleteAssetPlaybackIDPathParams struct {
	AssetID    string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
	PlaybackID string `pathParam:"style=simple,explode=false,name=PLAYBACK_ID"`
}

type DeleteAssetPlaybackIDRequest struct {
	PathParams DeleteAssetPlaybackIDPathParams
}

type DeleteAssetPlaybackIDResponse struct {
	ContentType string
	StatusCode  int
}
