package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type ListExportsViewsRequest struct {
	Retries *utils.RetryConfig
}

type ListExportsViewsResponse struct {
	ContentType                  string
	ListVideoViewExportsResponse *shared.ListVideoViewExportsResponse
	StatusCode                   int
	RawResponse                  *http.Response
}
