package operations

type DeleteURLSigningKeyPathParams struct {
	SigningKeyID string `pathParam:"style=simple,explode=false,name=SIGNING_KEY_ID"`
}

type DeleteURLSigningKeyRequest struct {
	PathParams DeleteURLSigningKeyPathParams
}

type DeleteURLSigningKeyResponse struct {
	ContentType string
	StatusCode  int
}
