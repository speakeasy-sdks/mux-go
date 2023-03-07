package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListFiltersResponse struct {
	ContentType         string
	ListFiltersResponse *shared.ListFiltersResponse
	StatusCode          int
	RawResponse         *http.Response
}
