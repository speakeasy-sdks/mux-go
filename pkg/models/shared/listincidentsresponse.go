package shared

type ListIncidentsResponse struct {
	Data          []Incident `json:"data,omitempty"`
	Timeframe     []int64    `json:"timeframe,omitempty"`
	TotalRowCount *int64     `json:"total_row_count,omitempty"`
}
