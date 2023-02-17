package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type DeleteURLSigningKeyPathParams struct {
	SigningKeyID string `pathParam:"style=simple,explode=false,name=SIGNING_KEY_ID"`
}

type DeleteURLSigningKeyRequest struct {
	PathParams DeleteURLSigningKeyPathParams
	Retries    *utils.RetryConfig
}

type DeleteURLSigningKeyResponse struct {
	ContentType string
	StatusCode  int
}
