package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type ListAssetsQueryParams struct {
	Limit        *int    `queryParam:"style=form,explode=true,name=limit"`
	LiveStreamID *string `queryParam:"style=form,explode=true,name=live_stream_id"`
	Page         *int    `queryParam:"style=form,explode=true,name=page"`
	UploadID     *string `queryParam:"style=form,explode=true,name=upload_id"`
}

type ListAssetsRequest struct {
	QueryParams ListAssetsQueryParams
	Retries     *utils.RetryConfig
}

type ListAssetsResponse struct {
	ContentType        string
	ListAssetsResponse *shared.ListAssetsResponse
	StatusCode         int
}
