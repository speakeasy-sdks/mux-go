package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetRealtimeTimeseriesRequest struct {
	RealtimeMetricID shared.RealtimeMetricIDEnum `pathParam:"style=simple,explode=false,name=REALTIME_METRIC_ID"`
	Filters          []string                    `queryParam:"style=form,explode=true,name=filters[]"`
}

type GetRealtimeTimeseriesResponse struct {
	ContentType                   string
	GetRealTimeTimeseriesResponse *shared.GetRealTimeTimeseriesResponse
	StatusCode                    int
	RawResponse                   *http.Response
}
