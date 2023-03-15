package operations

import (
	"net/http"
)

type DeletePlaybackRestrictionRequest struct {
	PlaybackRestrictionID string `pathParam:"style=simple,explode=false,name=PLAYBACK_RESTRICTION_ID"`
}

type DeletePlaybackRestrictionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
