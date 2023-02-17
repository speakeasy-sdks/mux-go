package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetMonitoringHistogramTimeseriesPathParams struct {
	MonitoringHistogramMetricID shared.MonitoringHistogramMetricIDEnum `pathParam:"style=simple,explode=false,name=MONITORING_HISTOGRAM_METRIC_ID"`
}

type GetMonitoringHistogramTimeseriesQueryParams struct {
	Filters []string `queryParam:"style=form,explode=true,name=filters[]"`
}

type GetMonitoringHistogramTimeseriesRequest struct {
	PathParams  GetMonitoringHistogramTimeseriesPathParams
	QueryParams GetMonitoringHistogramTimeseriesQueryParams
}

type GetMonitoringHistogramTimeseriesResponse struct {
	ContentType                              string
	GetMonitoringHistogramTimeseriesResponse *shared.GetMonitoringHistogramTimeseriesResponse
	StatusCode                               int
}
