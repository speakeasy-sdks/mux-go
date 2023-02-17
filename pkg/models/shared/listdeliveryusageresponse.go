package shared

type ListDeliveryUsageResponse struct {
	Data          []DeliveryReport `json:"data,omitempty"`
	Limit         *int64           `json:"limit,omitempty"`
	Timeframe     []int64          `json:"timeframe,omitempty"`
	TotalRowCount *int64           `json:"total_row_count,omitempty"`
}
