package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetRealtimeHistogramTimeseriesRequest struct {
	RealtimeHistogramMetricID shared.RealtimeHistogramMetricIDEnum `pathParam:"style=simple,explode=false,name=REALTIME_HISTOGRAM_METRIC_ID"`
	Filters                   []string                             `queryParam:"style=form,explode=true,name=filters[]"`
}

type GetRealtimeHistogramTimeseriesResponse struct {
	ContentType                            string
	GetRealTimeHistogramTimeseriesResponse *shared.GetRealTimeHistogramTimeseriesResponse
	StatusCode                             int
	RawResponse                            *http.Response
}
