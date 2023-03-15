package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListPlaybackRestrictionsRequest struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListPlaybackRestrictionsResponse struct {
	ContentType                      string
	ListPlaybackRestrictionsResponse *shared.ListPlaybackRestrictionsResponse
	StatusCode                       int
	RawResponse                      *http.Response
}
