package operations

type DeletePlaybackRestrictionPathParams struct {
	PlaybackRestrictionID string `pathParam:"style=simple,explode=false,name=PLAYBACK_RESTRICTION_ID"`
}

type DeletePlaybackRestrictionRequest struct {
	PathParams DeletePlaybackRestrictionPathParams
}

type DeletePlaybackRestrictionResponse struct {
	ContentType string
	StatusCode  int
}
