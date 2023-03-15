package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetPlaybackRestrictionRequest struct {
	PlaybackRestrictionID string `pathParam:"style=simple,explode=false,name=PLAYBACK_RESTRICTION_ID"`
}

type GetPlaybackRestrictionResponse struct {
	ContentType                 string
	PlaybackRestrictionResponse *shared.PlaybackRestrictionResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
