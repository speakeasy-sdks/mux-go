package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateAssetPlaybackIDPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type CreateAssetPlaybackIDRequest struct {
	PathParams CreateAssetPlaybackIDPathParams
	Request    shared.CreatePlaybackIDRequest `request:"mediaType=application/json"`
}

type CreateAssetPlaybackIDResponse struct {
	ContentType              string
	CreatePlaybackIDResponse *shared.CreatePlaybackIDResponse
	StatusCode               int
	RawResponse              *http.Response
}
