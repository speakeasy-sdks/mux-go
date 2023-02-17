package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type ListVideoViewsQueryParams struct {
	ErrorID        *int                       `queryParam:"style=form,explode=true,name=error_id"`
	Filters        []string                   `queryParam:"style=form,explode=true,name=filters[]"`
	Limit          *int                       `queryParam:"style=form,explode=true,name=limit"`
	OrderDirection *shared.OrderDirectionEnum `queryParam:"style=form,explode=true,name=order_direction"`
	Page           *int                       `queryParam:"style=form,explode=true,name=page"`
	Timeframe      []string                   `queryParam:"style=form,explode=true,name=timeframe[]"`
	ViewerID       *string                    `queryParam:"style=form,explode=true,name=viewer_id"`
}

type ListVideoViewsRequest struct {
	QueryParams ListVideoViewsQueryParams
	Retries     *utils.RetryConfig
}

type ListVideoViewsResponse struct {
	ContentType            string
	ListVideoViewsResponse *shared.ListVideoViewsResponse
	StatusCode             int
}
