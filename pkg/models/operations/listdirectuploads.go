package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type ListDirectUploadsQueryParams struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListDirectUploadsRequest struct {
	QueryParams ListDirectUploadsQueryParams
	Retries     *utils.RetryConfig
}

type ListDirectUploadsResponse struct {
	ContentType         string
	ListUploadsResponse *shared.ListUploadsResponse
	StatusCode          int
}
