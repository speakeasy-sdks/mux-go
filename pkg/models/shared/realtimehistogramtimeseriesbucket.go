package shared

type RealTimeHistogramTimeseriesBucket struct {
	End   *int64 `json:"end,omitempty"`
	Start *int64 `json:"start,omitempty"`
}
