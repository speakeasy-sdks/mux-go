package shared

type IncidentResponse struct {
	Data      *Incident `json:"data,omitempty"`
	Timeframe []int64   `json:"timeframe,omitempty"`
}
