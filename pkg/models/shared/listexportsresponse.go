package shared

type ListExportsResponse struct {
	Data          []string `json:"data,omitempty"`
	Timeframe     []int64  `json:"timeframe,omitempty"`
	TotalRowCount *int64   `json:"total_row_count,omitempty"`
}
