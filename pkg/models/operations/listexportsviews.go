package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type ListExportsViewsRequest struct {
	Retries *utils.RetryConfig
}

type ListExportsViewsResponse struct {
	ContentType                  string
	ListVideoViewExportsResponse *shared.ListVideoViewExportsResponse
	StatusCode                   int
}
