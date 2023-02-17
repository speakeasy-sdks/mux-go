package shared

type ListMonitoringMetricsResponseData struct {
	DisplayName *string `json:"display_name,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type ListMonitoringMetricsResponse struct {
	Data          []ListMonitoringMetricsResponseData `json:"data,omitempty"`
	Timeframe     []int64                             `json:"timeframe,omitempty"`
	TotalRowCount *int64                              `json:"total_row_count,omitempty"`
}
