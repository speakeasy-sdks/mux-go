package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetAssetOrLivestreamIDPathParams struct {
	PlaybackID string `pathParam:"style=simple,explode=false,name=PLAYBACK_ID"`
}

type GetAssetOrLivestreamIDRequest struct {
	PathParams GetAssetOrLivestreamIDPathParams
}

type GetAssetOrLivestreamIDResponse struct {
	ContentType                    string
	GetAssetOrLiveStreamIDResponse *shared.GetAssetOrLiveStreamIDResponse
	StatusCode                     int
	RawResponse                    *http.Response
}
