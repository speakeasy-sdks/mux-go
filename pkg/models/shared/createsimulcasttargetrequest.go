package shared

type CreateSimulcastTargetRequest struct {
	Passthrough *string `json:"passthrough,omitempty"`
	StreamKey   *string `json:"stream_key,omitempty"`
	URL         string  `json:"url"`
}
