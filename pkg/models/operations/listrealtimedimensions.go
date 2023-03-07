package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListRealtimeDimensionsResponse struct {
	ContentType                    string
	ListRealTimeDimensionsResponse *shared.ListRealTimeDimensionsResponse
	StatusCode                     int
	RawResponse                    *http.Response
}
