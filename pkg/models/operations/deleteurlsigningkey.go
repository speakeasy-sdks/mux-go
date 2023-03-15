package operations

import (
	"net/http"
)

type DeleteURLSigningKeyRequest struct {
	SigningKeyID string `pathParam:"style=simple,explode=false,name=SIGNING_KEY_ID"`
}

type DeleteURLSigningKeyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
