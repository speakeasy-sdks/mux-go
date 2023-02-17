package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type DeletePlaybackRestrictionPathParams struct {
	PlaybackRestrictionID string `pathParam:"style=simple,explode=false,name=PLAYBACK_RESTRICTION_ID"`
}

type DeletePlaybackRestrictionRequest struct {
	PathParams DeletePlaybackRestrictionPathParams
	Retries    *utils.RetryConfig
}

type DeletePlaybackRestrictionResponse struct {
	ContentType string
	StatusCode  int
}
