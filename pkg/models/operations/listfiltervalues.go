package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type ListFilterValuesPathParams struct {
	FilterID string `pathParam:"style=simple,explode=false,name=FILTER_ID"`
}

type ListFilterValuesQueryParams struct {
	Filters   []string `queryParam:"style=form,explode=true,name=filters[]"`
	Limit     *int     `queryParam:"style=form,explode=true,name=limit"`
	Page      *int     `queryParam:"style=form,explode=true,name=page"`
	Timeframe []string `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListFilterValuesRequest struct {
	PathParams  ListFilterValuesPathParams
	QueryParams ListFilterValuesQueryParams
	Retries     *utils.RetryConfig
}

type ListFilterValuesResponse struct {
	ContentType              string
	ListFilterValuesResponse *shared.ListFilterValuesResponse
	StatusCode               int
	RawResponse              *http.Response
}
