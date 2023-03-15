package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListDirectUploadsRequest struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListDirectUploadsResponse struct {
	ContentType         string
	ListUploadsResponse *shared.ListUploadsResponse
	StatusCode          int
	RawResponse         *http.Response
}
