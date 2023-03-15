package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListRelatedIncidentsRequest struct {
	IncidentID     string                     `pathParam:"style=simple,explode=false,name=INCIDENT_ID"`
	Limit          *int                       `queryParam:"style=form,explode=true,name=limit"`
	OrderBy        *shared.OrderByEnum        `queryParam:"style=form,explode=true,name=order_by"`
	OrderDirection *shared.OrderDirectionEnum `queryParam:"style=form,explode=true,name=order_direction"`
	Page           *int                       `queryParam:"style=form,explode=true,name=page"`
}

type ListRelatedIncidentsResponse struct {
	ContentType                  string
	ListRelatedIncidentsResponse *shared.ListRelatedIncidentsResponse
	StatusCode                   int
	RawResponse                  *http.Response
}
