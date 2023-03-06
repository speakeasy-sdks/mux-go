package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type ListErrorsQueryParams struct {
	Filters   []string `queryParam:"style=form,explode=true,name=filters[]"`
	Timeframe []string `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListErrorsRequest struct {
	QueryParams ListErrorsQueryParams
	Retries     *utils.RetryConfig
}

type ListErrorsResponse struct {
	ContentType        string
	ListErrorsResponse *shared.ListErrorsResponse
	StatusCode         int
	RawResponse        *http.Response
}
