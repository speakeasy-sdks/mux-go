package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type CancelDirectUploadPathParams struct {
	UploadID string `pathParam:"style=simple,explode=false,name=UPLOAD_ID"`
}

type CancelDirectUploadRequest struct {
	PathParams CancelDirectUploadPathParams
}

type CancelDirectUploadResponse struct {
	ContentType    string
	StatusCode     int
	UploadResponse *shared.UploadResponse
}
