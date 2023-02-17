package shared

type MonitoringHistogramTimeseriesBucket struct {
	End   *int64 `json:"end,omitempty"`
	Start *int64 `json:"start,omitempty"`
}
