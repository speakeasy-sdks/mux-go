package shared

type ListRealTimeMetricsResponseData struct {
	DisplayName *string `json:"display_name,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type ListRealTimeMetricsResponse struct {
	Data          []ListRealTimeMetricsResponseData `json:"data,omitempty"`
	Timeframe     []int64                           `json:"timeframe,omitempty"`
	TotalRowCount *int64                            `json:"total_row_count,omitempty"`
}
