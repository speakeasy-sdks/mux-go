package shared

type ListDimensionsResponseData struct {
	Advanced []string `json:"advanced,omitempty"`
	Basic    []string `json:"basic,omitempty"`
}

type ListDimensionsResponse struct {
	Data          *ListDimensionsResponseData `json:"data,omitempty"`
	Timeframe     []int64                     `json:"timeframe,omitempty"`
	TotalRowCount *int64                      `json:"total_row_count,omitempty"`
}
