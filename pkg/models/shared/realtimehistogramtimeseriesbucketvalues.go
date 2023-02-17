package shared

type RealTimeHistogramTimeseriesBucketValues struct {
	Count      *int64   `json:"count,omitempty"`
	Percentage *float64 `json:"percentage,omitempty"`
}
