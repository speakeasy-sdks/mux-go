package shared

type GetRealTimeHistogramTimeseriesResponseMeta struct {
	Buckets []RealTimeHistogramTimeseriesBucket `json:"buckets,omitempty"`
}

type GetRealTimeHistogramTimeseriesResponse struct {
	Data          []RealTimeHistogramTimeseriesDatapoint      `json:"data,omitempty"`
	Meta          *GetRealTimeHistogramTimeseriesResponseMeta `json:"meta,omitempty"`
	Timeframe     []int64                                     `json:"timeframe,omitempty"`
	TotalRowCount *int64                                      `json:"total_row_count,omitempty"`
}
