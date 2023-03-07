package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateURLSigningKeyResponse struct {
	ContentType        string
	SigningKeyResponse *shared.SigningKeyResponse
	StatusCode         int
	RawResponse        *http.Response
}
