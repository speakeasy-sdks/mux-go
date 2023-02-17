package shared

type ListVideoViewExportsResponse struct {
	Data          []ExportDate `json:"data,omitempty"`
	Timeframe     []int        `json:"timeframe,omitempty"`
	TotalRowCount *int         `json:"total_row_count,omitempty"`
}
