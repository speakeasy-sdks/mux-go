package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type ListFiltersRequest struct {
	Retries *utils.RetryConfig
}

type ListFiltersResponse struct {
	ContentType         string
	ListFiltersResponse *shared.ListFiltersResponse
	StatusCode          int
	RawResponse         *http.Response
}
