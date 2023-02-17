package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type CreateURLSigningKeyResponse struct {
	ContentType        string
	SigningKeyResponse *shared.SigningKeyResponse
	StatusCode         int
}
