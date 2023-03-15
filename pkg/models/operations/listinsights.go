package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListInsightsRequest struct {
	MetricID       shared.MetricIDEnum                  `pathParam:"style=simple,explode=false,name=METRIC_ID"`
	Measurement    *shared.MeasurementEnum              `queryParam:"style=form,explode=true,name=measurement"`
	OrderDirection *shared.OrderDirectionDeprecatedEnum `queryParam:"style=form,explode=true,name=order_direction"`
	Timeframe      []string                             `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListInsightsResponse struct {
	ContentType          string
	ListInsightsResponse *shared.ListInsightsResponse
	StatusCode           int
	RawResponse          *http.Response
}
