package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListFilterValuesRequest struct {
	FilterID  string   `pathParam:"style=simple,explode=false,name=FILTER_ID"`
	Filters   []string `queryParam:"style=form,explode=true,name=filters[]"`
	Limit     *int     `queryParam:"style=form,explode=true,name=limit"`
	Page      *int     `queryParam:"style=form,explode=true,name=page"`
	Timeframe []string `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListFilterValuesResponse struct {
	ContentType              string
	ListFilterValuesResponse *shared.ListFilterValuesResponse
	StatusCode               int
	RawResponse              *http.Response
}
