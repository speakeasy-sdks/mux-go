package shared

type RealTimeHistogramTimeseriesDatapoint struct {
	Average       *float64                                  `json:"average,omitempty"`
	BucketValues  []RealTimeHistogramTimeseriesBucketValues `json:"bucket_values,omitempty"`
	MaxPercentage *float64                                  `json:"max_percentage,omitempty"`
	Median        *float64                                  `json:"median,omitempty"`
	P95           *float64                                  `json:"p95,omitempty"`
	Sum           *int64                                    `json:"sum,omitempty"`
	Timestamp     *string                                   `json:"timestamp,omitempty"`
}
