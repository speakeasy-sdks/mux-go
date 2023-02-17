package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type ListPlaybackRestrictionsQueryParams struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListPlaybackRestrictionsRequest struct {
	QueryParams ListPlaybackRestrictionsQueryParams
}

type ListPlaybackRestrictionsResponse struct {
	ContentType                      string
	ListPlaybackRestrictionsResponse *shared.ListPlaybackRestrictionsResponse
	StatusCode                       int
}
