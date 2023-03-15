package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetIncidentRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=INCIDENT_ID"`
}

type GetIncidentResponse struct {
	ContentType      string
	IncidentResponse *shared.IncidentResponse
	StatusCode       int
	RawResponse      *http.Response
}
