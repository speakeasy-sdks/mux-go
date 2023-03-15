package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CancelDirectUploadRequest struct {
	UploadID string `pathParam:"style=simple,explode=false,name=UPLOAD_ID"`
}

type CancelDirectUploadResponse struct {
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
	UploadResponse *shared.UploadResponse
}
