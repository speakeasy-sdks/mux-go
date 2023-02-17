package shared

type VideoViewResponse struct {
	Data      *VideoView `json:"data,omitempty"`
	Timeframe []int64    `json:"timeframe,omitempty"`
}
