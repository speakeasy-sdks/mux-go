package shared

type GetMonitoringHistogramTimeseriesResponseMeta struct {
	Buckets []MonitoringHistogramTimeseriesBucket `json:"buckets,omitempty"`
}

type GetMonitoringHistogramTimeseriesResponse struct {
	Data          []MonitoringHistogramTimeseriesDatapoint      `json:"data,omitempty"`
	Meta          *GetMonitoringHistogramTimeseriesResponseMeta `json:"meta,omitempty"`
	Timeframe     []int64                                       `json:"timeframe,omitempty"`
	TotalRowCount *int64                                        `json:"total_row_count,omitempty"`
}
