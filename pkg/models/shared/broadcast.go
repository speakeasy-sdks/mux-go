package shared

type Broadcast struct {
	Background   *string                 `json:"background,omitempty"`
	ID           string                  `json:"id"`
	Layout       BroadcastLayoutEnum     `json:"layout"`
	LiveStreamID string                  `json:"live_stream_id"`
	Passthrough  *string                 `json:"passthrough,omitempty"`
	Resolution   BroadcastResolutionEnum `json:"resolution"`
	Status       BroadcastStatusEnum     `json:"status"`
}
