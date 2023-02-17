package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetIncidentPathParams struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=INCIDENT_ID"`
}

type GetIncidentRequest struct {
	PathParams GetIncidentPathParams
}

type GetIncidentResponse struct {
	ContentType      string
	IncidentResponse *shared.IncidentResponse
	StatusCode       int
}
