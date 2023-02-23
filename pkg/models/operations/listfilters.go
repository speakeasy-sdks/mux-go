package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type ListFiltersRequest struct {
	Retries *utils.RetryConfig
}

type ListFiltersResponse struct {
	ContentType         string
	ListFiltersResponse *shared.ListFiltersResponse
	StatusCode          int
}
