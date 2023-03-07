package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListIncidentsQueryParams struct {
	Limit          *int                       `queryParam:"style=form,explode=true,name=limit"`
	OrderBy        *shared.OrderByEnum        `queryParam:"style=form,explode=true,name=order_by"`
	OrderDirection *shared.OrderDirectionEnum `queryParam:"style=form,explode=true,name=order_direction"`
	Page           *int                       `queryParam:"style=form,explode=true,name=page"`
	Severity       *shared.SeverityEnum       `queryParam:"style=form,explode=true,name=severity"`
	Status         *shared.IncidentStatusEnum `queryParam:"style=form,explode=true,name=status"`
}

type ListIncidentsRequest struct {
	QueryParams ListIncidentsQueryParams
}

type ListIncidentsResponse struct {
	ContentType           string
	ListIncidentsResponse *shared.ListIncidentsResponse
	StatusCode            int
	RawResponse           *http.Response
}
