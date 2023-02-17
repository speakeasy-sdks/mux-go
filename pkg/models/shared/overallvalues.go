package shared

type OverallValues struct {
	GlobalValue    *float64 `json:"global_value,omitempty"`
	TotalViews     *int64   `json:"total_views,omitempty"`
	TotalWatchTime *int64   `json:"total_watch_time,omitempty"`
	Value          *float64 `json:"value,omitempty"`
}
