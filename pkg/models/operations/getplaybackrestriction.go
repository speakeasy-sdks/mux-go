package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetPlaybackRestrictionPathParams struct {
	PlaybackRestrictionID string `pathParam:"style=simple,explode=false,name=PLAYBACK_RESTRICTION_ID"`
}

type GetPlaybackRestrictionRequest struct {
	PathParams GetPlaybackRestrictionPathParams
}

type GetPlaybackRestrictionResponse struct {
	ContentType                 string
	PlaybackRestrictionResponse *shared.PlaybackRestrictionResponse
	StatusCode                  int
}
