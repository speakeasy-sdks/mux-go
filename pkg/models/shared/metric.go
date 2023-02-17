package shared

type Metric struct {
	Measurement *string  `json:"measurement,omitempty"`
	Metric      *string  `json:"metric,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Type        *string  `json:"type,omitempty"`
	Value       *float64 `json:"value,omitempty"`
}
