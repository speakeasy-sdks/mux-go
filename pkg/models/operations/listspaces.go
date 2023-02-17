package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type ListSpacesQueryParams struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListSpacesRequest struct {
	QueryParams ListSpacesQueryParams
	Retries     *utils.RetryConfig
}

type ListSpacesResponse struct {
	ContentType        string
	ListSpacesResponse *shared.ListSpacesResponse
	StatusCode         int
}
