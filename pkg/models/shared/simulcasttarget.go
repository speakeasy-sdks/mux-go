package shared

type SimulcastTargetStatusEnum string

const (
	SimulcastTargetStatusEnumIdle         SimulcastTargetStatusEnum = "idle"
	SimulcastTargetStatusEnumStarting     SimulcastTargetStatusEnum = "starting"
	SimulcastTargetStatusEnumBroadcasting SimulcastTargetStatusEnum = "broadcasting"
	SimulcastTargetStatusEnumErrored      SimulcastTargetStatusEnum = "errored"
)

type SimulcastTarget struct {
	ID          *string                    `json:"id,omitempty"`
	Passthrough *string                    `json:"passthrough,omitempty"`
	Status      *SimulcastTargetStatusEnum `json:"status,omitempty"`
	StreamKey   *string                    `json:"stream_key,omitempty"`
	URL         *string                    `json:"url,omitempty"`
}
