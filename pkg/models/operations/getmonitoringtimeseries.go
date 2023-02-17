package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetMonitoringTimeseriesPathParams struct {
	MonitoringMetricID shared.MonitoringMetricIDEnum `pathParam:"style=simple,explode=false,name=MONITORING_METRIC_ID"`
}

type GetMonitoringTimeseriesQueryParams struct {
	Filters []string `queryParam:"style=form,explode=true,name=filters[]"`
}

type GetMonitoringTimeseriesRequest struct {
	PathParams  GetMonitoringTimeseriesPathParams
	QueryParams GetMonitoringTimeseriesQueryParams
}

type GetMonitoringTimeseriesResponse struct {
	ContentType                     string
	GetMonitoringTimeseriesResponse *shared.GetMonitoringTimeseriesResponse
	StatusCode                      int
}
