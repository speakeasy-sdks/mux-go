package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListDimensionsResponse struct {
	ContentType            string
	ListDimensionsResponse *shared.ListDimensionsResponse
	StatusCode             int
	RawResponse            *http.Response
}
