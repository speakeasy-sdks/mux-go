package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type ListFiltersResponse struct {
	ContentType         string
	ListFiltersResponse *shared.ListFiltersResponse
	StatusCode          int
}
