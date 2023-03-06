package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type ListLiveStreamsQueryParams struct {
	Limit     *int                         `queryParam:"style=form,explode=true,name=limit"`
	Page      *int                         `queryParam:"style=form,explode=true,name=page"`
	Status    *shared.LiveStreamStatusEnum `queryParam:"style=form,explode=true,name=status"`
	StreamKey *string                      `queryParam:"style=form,explode=true,name=stream_key"`
}

type ListLiveStreamsRequest struct {
	QueryParams ListLiveStreamsQueryParams
	Retries     *utils.RetryConfig
}

type ListLiveStreamsResponse struct {
	ContentType             string
	ListLiveStreamsResponse *shared.ListLiveStreamsResponse
	StatusCode              int
	RawResponse             *http.Response
}
