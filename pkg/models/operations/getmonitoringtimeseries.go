package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetMonitoringTimeseriesRequest struct {
	MonitoringMetricID shared.MonitoringMetricIDEnum `pathParam:"style=simple,explode=false,name=MONITORING_METRIC_ID"`
	Filters            []string                      `queryParam:"style=form,explode=true,name=filters[]"`
}

type GetMonitoringTimeseriesResponse struct {
	ContentType                     string
	GetMonitoringTimeseriesResponse *shared.GetMonitoringTimeseriesResponse
	StatusCode                      int
	RawResponse                     *http.Response
}
