package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateDirectUploadRequest struct {
	Request shared.CreateUploadRequest `request:"mediaType=application/json"`
}

type CreateDirectUploadResponse struct {
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
	UploadResponse *shared.UploadResponse
}
