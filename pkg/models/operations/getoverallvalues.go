package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type GetOverallValuesPathParams struct {
	MetricID shared.MetricIDEnum `pathParam:"style=simple,explode=false,name=METRIC_ID"`
}

type GetOverallValuesQueryParams struct {
	Filters     []string                `queryParam:"style=form,explode=true,name=filters[]"`
	Measurement *shared.MeasurementEnum `queryParam:"style=form,explode=true,name=measurement"`
	Timeframe   []string                `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type GetOverallValuesRequest struct {
	PathParams  GetOverallValuesPathParams
	QueryParams GetOverallValuesQueryParams
	Retries     *utils.RetryConfig
}

type GetOverallValuesResponse struct {
	ContentType              string
	GetOverallValuesResponse *shared.GetOverallValuesResponse
	StatusCode               int
	RawResponse              *http.Response
}
