package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetDirectUploadPathParams struct {
	UploadID string `pathParam:"style=simple,explode=false,name=UPLOAD_ID"`
}

type GetDirectUploadRequest struct {
	PathParams GetDirectUploadPathParams
}

type GetDirectUploadResponse struct {
	ContentType    string
	StatusCode     int
	UploadResponse *shared.UploadResponse
}
