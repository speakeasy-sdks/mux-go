package operations

import (
	"net/http"
)

type DeleteSpaceBroadcastRequest struct {
	BroadcastID string `pathParam:"style=simple,explode=false,name=BROADCAST_ID"`
	SpaceID     string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type DeleteSpaceBroadcastResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
