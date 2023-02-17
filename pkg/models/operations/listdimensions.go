package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type ListDimensionsRequest struct {
	Retries *utils.RetryConfig
}

type ListDimensionsResponse struct {
	ContentType            string
	ListDimensionsResponse *shared.ListDimensionsResponse
	StatusCode             int
}
