package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListLiveStreamsRequest struct {
	Limit     *int                         `queryParam:"style=form,explode=true,name=limit"`
	Page      *int                         `queryParam:"style=form,explode=true,name=page"`
	Status    *shared.LiveStreamStatusEnum `queryParam:"style=form,explode=true,name=status"`
	StreamKey *string                      `queryParam:"style=form,explode=true,name=stream_key"`
}

type ListLiveStreamsResponse struct {
	ContentType             string
	ListLiveStreamsResponse *shared.ListLiveStreamsResponse
	StatusCode              int
	RawResponse             *http.Response
}
