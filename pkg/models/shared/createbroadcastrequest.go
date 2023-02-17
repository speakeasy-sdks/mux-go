package shared

type CreateBroadcastRequest struct {
	Background   *string                  `json:"background,omitempty"`
	Layout       *BroadcastLayoutEnum     `json:"layout,omitempty"`
	LiveStreamID string                   `json:"live_stream_id"`
	Passthrough  *string                  `json:"passthrough,omitempty"`
	Resolution   *BroadcastResolutionEnum `json:"resolution,omitempty"`
}
