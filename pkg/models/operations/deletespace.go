package operations

import (
	"net/http"
)

type DeleteSpaceRequest struct {
	SpaceID string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type DeleteSpaceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
