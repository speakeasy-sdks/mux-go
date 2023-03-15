package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetOverallValuesRequest struct {
	MetricID    shared.MetricIDEnum     `pathParam:"style=simple,explode=false,name=METRIC_ID"`
	Filters     []string                `queryParam:"style=form,explode=true,name=filters[]"`
	Measurement *shared.MeasurementEnum `queryParam:"style=form,explode=true,name=measurement"`
	Timeframe   []string                `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type GetOverallValuesResponse struct {
	ContentType              string
	GetOverallValuesResponse *shared.GetOverallValuesResponse
	StatusCode               int
	RawResponse              *http.Response
}
