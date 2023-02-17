package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type GetAssetPlaybackIDPathParams struct {
	AssetID    string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
	PlaybackID string `pathParam:"style=simple,explode=false,name=PLAYBACK_ID"`
}

type GetAssetPlaybackIDRequest struct {
	PathParams GetAssetPlaybackIDPathParams
	Retries    *utils.RetryConfig
}

type GetAssetPlaybackIDResponse struct {
	ContentType                string
	GetAssetPlaybackIDResponse *shared.GetAssetPlaybackIDResponse
	StatusCode                 int
}
