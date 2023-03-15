package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListErrorsRequest struct {
	Filters   []string `queryParam:"style=form,explode=true,name=filters[]"`
	Timeframe []string `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListErrorsResponse struct {
	ContentType        string
	ListErrorsResponse *shared.ListErrorsResponse
	StatusCode         int
	RawResponse        *http.Response
}
