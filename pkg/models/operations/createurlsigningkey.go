package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type CreateURLSigningKeyRequest struct {
	Retries *utils.RetryConfig
}

type CreateURLSigningKeyResponse struct {
	ContentType        string
	SigningKeyResponse *shared.SigningKeyResponse
	StatusCode         int
}
