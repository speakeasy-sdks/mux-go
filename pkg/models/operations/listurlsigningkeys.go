package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type ListURLSigningKeysQueryParams struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListURLSigningKeysRequest struct {
	QueryParams ListURLSigningKeysQueryParams
}

type ListURLSigningKeysResponse struct {
	ContentType             string
	ListSigningKeysResponse *shared.ListSigningKeysResponse
	StatusCode              int
}
