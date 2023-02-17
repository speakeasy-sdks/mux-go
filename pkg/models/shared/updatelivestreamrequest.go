package shared

type UpdateLiveStreamRequestLatencyModeEnum string

const (
	UpdateLiveStreamRequestLatencyModeEnumLow      UpdateLiveStreamRequestLatencyModeEnum = "low"
	UpdateLiveStreamRequestLatencyModeEnumReduced  UpdateLiveStreamRequestLatencyModeEnum = "reduced"
	UpdateLiveStreamRequestLatencyModeEnumStandard UpdateLiveStreamRequestLatencyModeEnum = "standard"
)

type UpdateLiveStreamRequest struct {
	LatencyMode                *UpdateLiveStreamRequestLatencyModeEnum `json:"latency_mode,omitempty"`
	MaxContinuousDuration      *int                                    `json:"max_continuous_duration,omitempty"`
	Passthrough                *string                                 `json:"passthrough,omitempty"`
	ReconnectSlateURL          *string                                 `json:"reconnect_slate_url,omitempty"`
	ReconnectWindow            *float32                                `json:"reconnect_window,omitempty"`
	UseSlateForStandardLatency *bool                                   `json:"use_slate_for_standard_latency,omitempty"`
}
