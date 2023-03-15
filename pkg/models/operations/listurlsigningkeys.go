package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListURLSigningKeysRequest struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListURLSigningKeysResponse struct {
	ContentType             string
	ListSigningKeysResponse *shared.ListSigningKeysResponse
	StatusCode              int
	RawResponse             *http.Response
}
