package operations

type DeleteSpaceBroadcastPathParams struct {
	BroadcastID string `pathParam:"style=simple,explode=false,name=BROADCAST_ID"`
	SpaceID     string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type DeleteSpaceBroadcastRequest struct {
	PathParams DeleteSpaceBroadcastPathParams
}

type DeleteSpaceBroadcastResponse struct {
	ContentType string
	StatusCode  int
}
