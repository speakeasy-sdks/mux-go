package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type GetRealtimeHistogramTimeseriesPathParams struct {
	RealtimeHistogramMetricID shared.RealtimeHistogramMetricIDEnum `pathParam:"style=simple,explode=false,name=REALTIME_HISTOGRAM_METRIC_ID"`
}

type GetRealtimeHistogramTimeseriesQueryParams struct {
	Filters []string `queryParam:"style=form,explode=true,name=filters[]"`
}

type GetRealtimeHistogramTimeseriesRequest struct {
	PathParams  GetRealtimeHistogramTimeseriesPathParams
	QueryParams GetRealtimeHistogramTimeseriesQueryParams
	Retries     *utils.RetryConfig
}

type GetRealtimeHistogramTimeseriesResponse struct {
	ContentType                            string
	GetRealTimeHistogramTimeseriesResponse *shared.GetRealTimeHistogramTimeseriesResponse
	StatusCode                             int
}
