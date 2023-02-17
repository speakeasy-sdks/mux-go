package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type GetMonitoringBreakdownPathParams struct {
	MonitoringMetricID shared.MonitoringMetricIDEnum `pathParam:"style=simple,explode=false,name=MONITORING_METRIC_ID"`
}

type GetMonitoringBreakdownQueryParams struct {
	Dimension      *shared.MonitoringDimensionEnum `queryParam:"style=form,explode=true,name=dimension"`
	Filters        []string                        `queryParam:"style=form,explode=true,name=filters[]"`
	OrderBy        *shared.OrderByEnum             `queryParam:"style=form,explode=true,name=order_by"`
	OrderDirection *shared.OrderDirectionEnum      `queryParam:"style=form,explode=true,name=order_direction"`
	Timestamp      *int                            `queryParam:"style=form,explode=true,name=timestamp"`
}

type GetMonitoringBreakdownRequest struct {
	PathParams  GetMonitoringBreakdownPathParams
	QueryParams GetMonitoringBreakdownQueryParams
	Retries     *utils.RetryConfig
}

type GetMonitoringBreakdownResponse struct {
	ContentType                    string
	GetMonitoringBreakdownResponse *shared.GetMonitoringBreakdownResponse
	StatusCode                     int
}
