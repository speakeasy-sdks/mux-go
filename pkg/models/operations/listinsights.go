package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type ListInsightsPathParams struct {
	MetricID shared.MetricIDEnum `pathParam:"style=simple,explode=false,name=METRIC_ID"`
}

type ListInsightsQueryParams struct {
	Measurement    *shared.MeasurementEnum              `queryParam:"style=form,explode=true,name=measurement"`
	OrderDirection *shared.OrderDirectionDeprecatedEnum `queryParam:"style=form,explode=true,name=order_direction"`
	Timeframe      []string                             `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListInsightsRequest struct {
	PathParams  ListInsightsPathParams
	QueryParams ListInsightsQueryParams
	Retries     *utils.RetryConfig
}

type ListInsightsResponse struct {
	ContentType          string
	ListInsightsResponse *shared.ListInsightsResponse
	StatusCode           int
}
