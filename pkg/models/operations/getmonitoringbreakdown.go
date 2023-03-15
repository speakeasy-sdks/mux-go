package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetMonitoringBreakdownRequest struct {
	MonitoringMetricID shared.MonitoringMetricIDEnum   `pathParam:"style=simple,explode=false,name=MONITORING_METRIC_ID"`
	Dimension          *shared.MonitoringDimensionEnum `queryParam:"style=form,explode=true,name=dimension"`
	Filters            []string                        `queryParam:"style=form,explode=true,name=filters[]"`
	OrderBy            *shared.OrderByEnum             `queryParam:"style=form,explode=true,name=order_by"`
	OrderDirection     *shared.OrderDirectionEnum      `queryParam:"style=form,explode=true,name=order_direction"`
	Timestamp          *int                            `queryParam:"style=form,explode=true,name=timestamp"`
}

type GetMonitoringBreakdownResponse struct {
	ContentType                    string
	GetMonitoringBreakdownResponse *shared.GetMonitoringBreakdownResponse
	StatusCode                     int
	RawResponse                    *http.Response
}
