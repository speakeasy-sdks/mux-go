package shared

type RealTimeBreakdownValue struct {
	ConcurrentViewers *int64   `json:"concurrent_viewers,omitempty"`
	DisplayValue      *string  `json:"display_value,omitempty"`
	MetricValue       *float64 `json:"metric_value,omitempty"`
	NegativeImpact    *int64   `json:"negative_impact,omitempty"`
	Value             *string  `json:"value,omitempty"`
}
