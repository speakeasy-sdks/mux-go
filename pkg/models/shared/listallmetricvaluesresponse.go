package shared

type ListAllMetricValuesResponse struct {
	Data          []Score `json:"data,omitempty"`
	Timeframe     []int64 `json:"timeframe,omitempty"`
	TotalRowCount *int64  `json:"total_row_count,omitempty"`
}
