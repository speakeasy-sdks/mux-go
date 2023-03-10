package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateAssetTrackPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type CreateAssetTrackRequest struct {
	PathParams CreateAssetTrackPathParams
	Request    shared.CreateTrackRequest `request:"mediaType=application/json"`
}

type CreateAssetTrackResponse struct {
	ContentType         string
	CreateTrackResponse *shared.CreateTrackResponse
	StatusCode          int
	RawResponse         *http.Response
}
