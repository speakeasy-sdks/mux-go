package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreatePlaybackRestrictionRequest struct {
	Request shared.CreatePlaybackRestrictionRequest `request:"mediaType=application/json"`
}

type CreatePlaybackRestrictionResponse struct {
	ContentType                 string
	PlaybackRestrictionResponse *shared.PlaybackRestrictionResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
