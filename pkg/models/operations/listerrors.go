package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListErrorsQueryParams struct {
	Filters   []string `queryParam:"style=form,explode=true,name=filters[]"`
	Timeframe []string `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListErrorsRequest struct {
	QueryParams ListErrorsQueryParams
}

type ListErrorsResponse struct {
	ContentType        string
	ListErrorsResponse *shared.ListErrorsResponse
	StatusCode         int
	RawResponse        *http.Response
}
