package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListAssetsRequest struct {
	Limit        *int    `queryParam:"style=form,explode=true,name=limit"`
	LiveStreamID *string `queryParam:"style=form,explode=true,name=live_stream_id"`
	Page         *int    `queryParam:"style=form,explode=true,name=page"`
	UploadID     *string `queryParam:"style=form,explode=true,name=upload_id"`
}

type ListAssetsResponse struct {
	ContentType        string
	ListAssetsResponse *shared.ListAssetsResponse
	StatusCode         int
	RawResponse        *http.Response
}
