package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type ListExportsViewsResponse struct {
	ContentType                  string
	ListVideoViewExportsResponse *shared.ListVideoViewExportsResponse
	StatusCode                   int
}
