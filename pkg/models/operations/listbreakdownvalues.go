package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListBreakdownValuesPathParams struct {
	MetricID shared.MetricIDEnum `pathParam:"style=simple,explode=false,name=METRIC_ID"`
}

type ListBreakdownValuesQueryParams struct {
	Filters        []string                   `queryParam:"style=form,explode=true,name=filters[]"`
	GroupBy        *shared.GroupByEnum        `queryParam:"style=form,explode=true,name=group_by"`
	Limit          *int                       `queryParam:"style=form,explode=true,name=limit"`
	Measurement    *shared.MeasurementEnum    `queryParam:"style=form,explode=true,name=measurement"`
	OrderBy        *shared.OrderByEnum        `queryParam:"style=form,explode=true,name=order_by"`
	OrderDirection *shared.OrderDirectionEnum `queryParam:"style=form,explode=true,name=order_direction"`
	Page           *int                       `queryParam:"style=form,explode=true,name=page"`
	Timeframe      []string                   `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListBreakdownValuesRequest struct {
	PathParams  ListBreakdownValuesPathParams
	QueryParams ListBreakdownValuesQueryParams
}

type ListBreakdownValuesResponse struct {
	ContentType                 string
	ListBreakdownValuesResponse *shared.ListBreakdownValuesResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
