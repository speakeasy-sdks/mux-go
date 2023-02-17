package shared

type RealTimeTimeseriesDatapoint struct {
	ConcurrentViewers *int64   `json:"concurrent_viewers,omitempty"`
	Date              *string  `json:"date,omitempty"`
	Value             *float64 `json:"value,omitempty"`
}
