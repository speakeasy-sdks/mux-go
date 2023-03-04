package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type CancelDirectUploadPathParams struct {
	UploadID string `pathParam:"style=simple,explode=false,name=UPLOAD_ID"`
}

type CancelDirectUploadRequest struct {
	PathParams CancelDirectUploadPathParams
	Retries    *utils.RetryConfig
}

type CancelDirectUploadResponse struct {
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
	UploadResponse *shared.UploadResponse
}
