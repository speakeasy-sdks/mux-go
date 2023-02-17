package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type GetRealtimeBreakdownPathParams struct {
	RealtimeMetricID shared.RealtimeMetricIDEnum `pathParam:"style=simple,explode=false,name=REALTIME_METRIC_ID"`
}

type GetRealtimeBreakdownQueryParams struct {
	Dimension      *shared.RealtimeDimensionEnum `queryParam:"style=form,explode=true,name=dimension"`
	Filters        []string                      `queryParam:"style=form,explode=true,name=filters[]"`
	OrderBy        *shared.OrderByEnum           `queryParam:"style=form,explode=true,name=order_by"`
	OrderDirection *shared.OrderDirectionEnum    `queryParam:"style=form,explode=true,name=order_direction"`
	Timestamp      *int                          `queryParam:"style=form,explode=true,name=timestamp"`
}

type GetRealtimeBreakdownRequest struct {
	PathParams  GetRealtimeBreakdownPathParams
	QueryParams GetRealtimeBreakdownQueryParams
	Retries     *utils.RetryConfig
}

type GetRealtimeBreakdownResponse struct {
	ContentType                  string
	GetRealTimeBreakdownResponse *shared.GetRealTimeBreakdownResponse
	StatusCode                   int
}
