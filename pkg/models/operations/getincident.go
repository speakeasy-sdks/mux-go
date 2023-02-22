package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type GetIncidentPathParams struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=INCIDENT_ID"`
}

type GetIncidentRequest struct {
	PathParams GetIncidentPathParams
	Retries    *utils.RetryConfig
}

type GetIncidentResponse struct {
	ContentType      string
	IncidentResponse *shared.IncidentResponse
	StatusCode       int
}
