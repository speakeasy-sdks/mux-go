package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type ListAllMetricValuesQueryParams struct {
	Dimension *shared.DimensionEnum `queryParam:"style=form,explode=true,name=dimension"`
	Filters   []string              `queryParam:"style=form,explode=true,name=filters[]"`
	Timeframe []string              `queryParam:"style=form,explode=true,name=timeframe[]"`
	Value     *string               `queryParam:"style=form,explode=true,name=value"`
}

type ListAllMetricValuesRequest struct {
	QueryParams ListAllMetricValuesQueryParams
	Retries     *utils.RetryConfig
}

type ListAllMetricValuesResponse struct {
	ContentType                 string
	ListAllMetricValuesResponse *shared.ListAllMetricValuesResponse
	StatusCode                  int
}
