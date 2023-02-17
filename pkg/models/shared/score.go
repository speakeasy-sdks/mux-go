package shared

type Score struct {
	Items     []Metric `json:"items,omitempty"`
	Metric    *string  `json:"metric,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Value     *float64 `json:"value,omitempty"`
	ViewCount *int64   `json:"view_count,omitempty"`
	WatchTime *int64   `json:"watch_time,omitempty"`
}
