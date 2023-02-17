package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type CreateDirectUploadRequest struct {
	Request shared.CreateUploadRequest `request:"mediaType=application/json"`
	Retries *utils.RetryConfig
}

type CreateDirectUploadResponse struct {
	ContentType    string
	StatusCode     int
	UploadResponse *shared.UploadResponse
}
