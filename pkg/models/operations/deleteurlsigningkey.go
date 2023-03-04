package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
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
	RawResponse *http.Response
}
