package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetURLSigningKeyPathParams struct {
	SigningKeyID string `pathParam:"style=simple,explode=false,name=SIGNING_KEY_ID"`
}

type GetURLSigningKeyRequest struct {
	PathParams GetURLSigningKeyPathParams
}

type GetURLSigningKeyResponse struct {
	ContentType        string
	SigningKeyResponse *shared.SigningKeyResponse
	StatusCode         int
}
