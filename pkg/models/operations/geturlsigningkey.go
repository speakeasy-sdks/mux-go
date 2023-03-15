package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetURLSigningKeyRequest struct {
	SigningKeyID string `pathParam:"style=simple,explode=false,name=SIGNING_KEY_ID"`
}

type GetURLSigningKeyResponse struct {
	ContentType        string
	SigningKeyResponse *shared.SigningKeyResponse
	StatusCode         int
	RawResponse        *http.Response
}
