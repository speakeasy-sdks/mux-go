package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type GetRealtimeTimeseriesPathParams struct {
	RealtimeMetricID shared.RealtimeMetricIDEnum `pathParam:"style=simple,explode=false,name=REALTIME_METRIC_ID"`
}

type GetRealtimeTimeseriesQueryParams struct {
	Filters []string `queryParam:"style=form,explode=true,name=filters[]"`
}

type GetRealtimeTimeseriesRequest struct {
	PathParams  GetRealtimeTimeseriesPathParams
	QueryParams GetRealtimeTimeseriesQueryParams
	Retries     *utils.RetryConfig
}

type GetRealtimeTimeseriesResponse struct {
	ContentType                   string
	GetRealTimeTimeseriesResponse *shared.GetRealTimeTimeseriesResponse
	StatusCode                    int
	RawResponse                   *http.Response
}
