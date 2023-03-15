package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetDirectUploadRequest struct {
	UploadID string `pathParam:"style=simple,explode=false,name=UPLOAD_ID"`
}

type GetDirectUploadResponse struct {
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
	UploadResponse *shared.UploadResponse
}
