package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type ListURLSigningKeysQueryParams struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListURLSigningKeysRequest struct {
	QueryParams ListURLSigningKeysQueryParams
	Retries     *utils.RetryConfig
}

type ListURLSigningKeysResponse struct {
	ContentType             string
	ListSigningKeysResponse *shared.ListSigningKeysResponse
	StatusCode              int
}
