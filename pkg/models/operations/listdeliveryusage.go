package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListDeliveryUsageQueryParams struct {
	AssetID      *string  `queryParam:"style=form,explode=true,name=asset_id"`
	Limit        *int     `queryParam:"style=form,explode=true,name=limit"`
	LiveStreamID *string  `queryParam:"style=form,explode=true,name=live_stream_id"`
	Page         *int     `queryParam:"style=form,explode=true,name=page"`
	Timeframe    []string `queryParam:"style=form,explode=true,name=timeframe[]"`
}

type ListDeliveryUsageRequest struct {
	QueryParams ListDeliveryUsageQueryParams
}

type ListDeliveryUsageResponse struct {
	ContentType               string
	ListDeliveryUsageResponse *shared.ListDeliveryUsageResponse
	StatusCode                int
	RawResponse               *http.Response
}
