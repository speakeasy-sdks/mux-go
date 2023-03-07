package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type UpdateReferrerDomainRestrictionPathParams struct {
	PlaybackRestrictionID string `pathParam:"style=simple,explode=false,name=PLAYBACK_RESTRICTION_ID"`
}

type UpdateReferrerDomainRestrictionRequest struct {
	PathParams UpdateReferrerDomainRestrictionPathParams
	Request    interface{} `request:"mediaType=application/json"`
}

type UpdateReferrerDomainRestrictionResponse struct {
	ContentType                 string
	PlaybackRestrictionResponse *shared.PlaybackRestrictionResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
