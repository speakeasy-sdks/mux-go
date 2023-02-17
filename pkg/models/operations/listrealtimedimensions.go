package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type ListRealtimeDimensionsResponse struct {
	ContentType                    string
	ListRealTimeDimensionsResponse *shared.ListRealTimeDimensionsResponse
	StatusCode                     int
}
