package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListAllMetricValuesRequest struct {
	Dimension *shared.DimensionEnum `queryParam:"style=form,explode=true,name=dimension"`
	Filters   []string              `queryParam:"style=form,explode=true,name=filters[]"`
	Timeframe []string              `queryParam:"style=form,explode=true,name=timeframe[]"`
	Value     *string               `queryParam:"style=form,explode=true,name=value"`
}

type ListAllMetricValuesResponse struct {
	ContentType                 string
	ListAllMetricValuesResponse *shared.ListAllMetricValuesResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
