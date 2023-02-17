package shared

type Insight struct {
	FilterColumn        *string  `json:"filter_column,omitempty"`
	FilterValue         *string  `json:"filter_value,omitempty"`
	Metric              *float64 `json:"metric,omitempty"`
	NegativeImpactScore *float32 `json:"negative_impact_score,omitempty"`
	TotalViews          *int64   `json:"total_views,omitempty"`
	TotalWatchTime      *int64   `json:"total_watch_time,omitempty"`
}
