package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListExportsResponse struct {
	ContentType         string
	ListExportsResponse *shared.ListExportsResponse
	StatusCode          int
	RawResponse         *http.Response
}
