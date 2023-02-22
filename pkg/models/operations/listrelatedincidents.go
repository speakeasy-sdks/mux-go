package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type ListRelatedIncidentsPathParams struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=INCIDENT_ID"`
}

type ListRelatedIncidentsQueryParams struct {
	Limit          *int                       `queryParam:"style=form,explode=true,name=limit"`
	OrderBy        *shared.OrderByEnum        `queryParam:"style=form,explode=true,name=order_by"`
	OrderDirection *shared.OrderDirectionEnum `queryParam:"style=form,explode=true,name=order_direction"`
	Page           *int                       `queryParam:"style=form,explode=true,name=page"`
}

type ListRelatedIncidentsRequest struct {
	PathParams  ListRelatedIncidentsPathParams
	QueryParams ListRelatedIncidentsQueryParams
	Retries     *utils.RetryConfig
}

type ListRelatedIncidentsResponse struct {
	ContentType                  string
	ListRelatedIncidentsResponse *shared.ListRelatedIncidentsResponse
	StatusCode                   int
}
