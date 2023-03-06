package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type CreateDirectUploadRequest struct {
	Request shared.CreateUploadRequest `request:"mediaType=application/json"`
	Retries *utils.RetryConfig
}

type CreateDirectUploadResponse struct {
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
	UploadResponse *shared.UploadResponse
}
