package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListDimensionValuesPathParams struct {
	DimensionID string `pathParam:"style=simple,explode=false,name=DIMENSION_ID"`
}

type ListDimensionValuesQueryParams struct {
	Filters   []string `queryParam:"style=form,explode=true,name=filters[]"`
	Limit     *int     `queryParam:"style=form,explode=true,name=limit"`
	Page      *int     `queryParam:"style=form,explode=true,name=page"`
	Timeframe []string `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListDimensionValuesRequest struct {
	PathParams  ListDimensionValuesPathParams
	QueryParams ListDimensionValuesQueryParams
}

type ListDimensionValuesResponse struct {
	ContentType                 string
	ListDimensionValuesResponse *shared.ListDimensionValuesResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
