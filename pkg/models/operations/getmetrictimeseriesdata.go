package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetMetricTimeseriesDataPathParams struct {
	MetricID shared.MetricIDEnum `pathParam:"style=simple,explode=false,name=METRIC_ID"`
}

type GetMetricTimeseriesDataQueryParams struct {
	Filters        []string                      `queryParam:"style=form,explode=true,name=filters[]"`
	GroupBy        *shared.TimeseriesGroupByEnum `queryParam:"style=form,explode=true,name=group_by"`
	Measurement    *shared.MeasurementEnum       `queryParam:"style=form,explode=true,name=measurement"`
	OrderDirection *shared.OrderDirectionEnum    `queryParam:"style=form,explode=true,name=order_direction"`
	Timeframe      []string                      `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type GetMetricTimeseriesDataRequest struct {
	PathParams  GetMetricTimeseriesDataPathParams
	QueryParams GetMetricTimeseriesDataQueryParams
}

type GetMetricTimeseriesDataResponse struct {
	ContentType                     string
	GetMetricTimeseriesDataResponse *shared.GetMetricTimeseriesDataResponse
	StatusCode                      int
	RawResponse                     *http.Response
}
