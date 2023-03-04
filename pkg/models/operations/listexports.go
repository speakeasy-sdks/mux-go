package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type ListExportsRequest struct {
	Retries *utils.RetryConfig
}

type ListExportsResponse struct {
	ContentType         string
	ListExportsResponse *shared.ListExportsResponse
	StatusCode          int
	RawResponse         *http.Response
}
