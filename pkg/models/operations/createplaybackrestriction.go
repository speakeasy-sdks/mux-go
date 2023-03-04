package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type CreatePlaybackRestrictionRequest struct {
	Request shared.CreatePlaybackRestrictionRequest `request:"mediaType=application/json"`
	Retries *utils.RetryConfig
}

type CreatePlaybackRestrictionResponse struct {
	ContentType                 string
	PlaybackRestrictionResponse *shared.PlaybackRestrictionResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
