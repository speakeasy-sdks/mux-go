package operations

type DeleteLiveStreamPathParams struct {
	LiveStreamID string `pathParam:"style=simple,explode=false,name=LIVE_STREAM_ID"`
}

type DeleteLiveStreamRequest struct {
	PathParams DeleteLiveStreamPathParams
}

type DeleteLiveStreamResponse struct {
	ContentType string
	StatusCode  int
}
