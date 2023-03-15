package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetMonitoringHistogramTimeseriesRequest struct {
	MonitoringHistogramMetricID shared.MonitoringHistogramMetricIDEnum `pathParam:"style=simple,explode=false,name=MONITORING_HISTOGRAM_METRIC_ID"`
	Filters                     []string                               `queryParam:"style=form,explode=true,name=filters[]"`
}

type GetMonitoringHistogramTimeseriesResponse struct {
	ContentType                              string
	GetMonitoringHistogramTimeseriesResponse *shared.GetMonitoringHistogramTimeseriesResponse
	StatusCode                               int
	RawResponse                              *http.Response
}
