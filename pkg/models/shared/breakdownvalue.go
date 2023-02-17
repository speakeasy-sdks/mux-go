package shared

type BreakdownValue struct {
	Field          *string  `json:"field,omitempty"`
	NegativeImpact *int     `json:"negative_impact,omitempty"`
	TotalWatchTime *int64   `json:"total_watch_time,omitempty"`
	Value          *float64 `json:"value,omitempty"`
	Views          *int64   `json:"views,omitempty"`
}
